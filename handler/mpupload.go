package handler

import (
	"fmt"
	 "github.com/skydrive/cache/redisconn"
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/meta"
	"github.com/skydrive/utils"
	"math"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type MultiPartInfo struct {
	UploadId     string `json:"uploadid"`
	ChunkCount   int    `json:"chunkcount"`
	ChunkSize    int    `json:"chunksize"`
	Filesha1     string `json:"sha1,omitempty"`
	FileName     string `json:"filename,omitempty"`
	FileSize     int    `json:"size,omitempty"`
	Pid     int64    `json:"pid,omitempty"`
	SuccessIndex []int  `json:"successchunkindex,omitempty"`
}

//初始化分块上传
func InitMultipartUploadHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	filesha1 := r.FormValue("sha1")
	filename := r.FormValue("filename")
	filesize, error := strconv.Atoi(r.FormValue("filesize"))
	pid, _ := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	println("aaa", error, filesha1)
	if len(filesha1) == 0 || filesha1 == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	_, ok := meta.GetFileMeta(filesha1)
	if ok {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "文件已经上传:"+filesha1)
		return
	}
	uploadId := filesha1 + fmt.Sprintf("%x", time.Now().UnixNano())
	chunkcount := int(math.Ceil(float64(filesize) / config.CHUNK_Size))
	client := redisconn.GetRedisClient()
	client.HMSet(redisconn.CTX, "MP_"+uploadId, "chunkcount", chunkcount)
	client.HMSet(redisconn.CTX, "MP_"+uploadId, "filesha1", filesha1)
	client.HMSet(redisconn.CTX, "MP_"+uploadId, "filesize", filesize)
	client.HMSet(redisconn.CTX, "MP_"+uploadId, "filename", filename)
	client.HMSet(redisconn.CTX, "MP_"+uploadId, "pid", pid)

	ReturnResponse(w, config.Net_SuccessCode, "成功", &MultiPartInfo{
		UploadId:   uploadId,
		ChunkCount: chunkcount,
		ChunkSize:  config.CHUNK_Size,
		Filesha1:   filesha1,
		Pid:   pid,
		FileName:   filename,
		FileSize:   filesize,
	})
}

//分块上传部分
func UploadMultipartHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	fmt.Println(r.Method)
	if r.Method != "POST" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "bad request")
		return
	}
	uploadId := r.FormValue("uploadId")
	chunkindex := r.FormValue("chunkindex")
	if len(uploadId) == 0 || uploadId == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	//todo 写文件
	fpath := "temp/mpdata/" + uploadId + "/" + chunkindex
	os.MkdirAll(path.Dir(fpath), 0744)
	create, error := os.Create(fpath)
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "文件创建失败")
		return
	}
	defer create.Close()
	buf := make([]byte, 1024*1024)
	for {
		index, error := r.Body.Read(buf)
		create.Write(buf[:index])
		if error != nil {
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, "文件写入失败")
			return
		}
	}
	client := redisconn.GetRedisClient()
	client.HMSet(redisconn.CTX, "MP_"+uploadId, "chunk_index"+chunkindex, 1)
	ReturnResponseCodeMessage(w, config.Net_SuccessCode, "成功")
}

//通知分块上传完成
func FinishMultipartUploadHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	filesha1 := r.FormValue("sha1")
	uploadId := r.FormValue("uploadId")
	pid, _ := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	if len(filesha1) == 0 || filesha1 == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	_, ok := meta.GetFileMeta(filesha1)
	if ok {
		ReturnResponseCodeMessage(w, config.Net_SuccessAginCode, "文件已经上传:"+filesha1)
		return
	}
	client := redisconn.GetRedisClient()

	result, _ := client.HGetAll(redisconn.CTX, "MP_"+uploadId).Result()

	if result == nil || len(result) == 0 {
		//未上传
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "文件上传任务不存在:"+filesha1)
		return
	} else {
		filesha1Cache := result["filesha1"]
		//对比sha1
		if filesha1 != filesha1Cache {
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, "文件 sha1 错误:"+filesha1)
			return
		}
		var SuccessIndex []int
		var filesize int
		var pids int64
		successCount := 0
		countSum := 0
		filename := ""
		fmt.Println("result", result)
		for key, value := range result {
			if strings.HasPrefix(key, "chunk_index") && value == "1" {
				successCount++
				index, _ := strconv.Atoi(key[len("chunk_index")-1:])
				SuccessIndex = append(SuccessIndex, index)
			} else if key == "chunkcount" {
				countSum, _ = strconv.Atoi(value)
			} else if key == "filename" {
				filename = value
			} else if key == "filesize" {
				filesize, _ = strconv.Atoi(value)
			}else if key == "pid" {
				pids, _ = strconv.ParseInt(value, 10, 64)
			}
		}
		fmt.Println(pids)
		resultCode := config.Net_SuccessCode
		errorMessage := "成功"
		if countSum != successCount {
			resultCode = config.Net_ErrorCode
			errorMessage = "未上传完成"
		}
		//todo 合并文件，更新db
		fpath := "temp/mpdata/" + uploadId + "/"
		if merge := utils.FileMerge(fpath, filename); merge {
			//fileLocation := fpath + filename
			//todo 如果不存在 先插入文件表
			/*if !db.SaveFileInfo(filesha1, filename, int64(filesize), fpath+filename) {
				//插入文件表不成功
				ReturnResponseCodeMessage(w, config.Net_ErrorCode, "系统文件db保存失败")
				return
			}*/
			//查看是否已经保存过
			if _, err := db.GetUserFileInfoByUidSha1(filesha1, utoken.Uid.Int64); err != nil {
				//避免重复保存
				//文件表已经插入成功,再插入用户文件表
				/*if !db.SaveUserFileInfo(utoken.Uid.Int64, pid, utoken.Phone.String, filesha1, filename, fileLocation, int64(filesize)) {
					ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户文件保存失败")
					return
				}*/
			}
			ReturnResponse(w, int32(resultCode), errorMessage, &MultiPartInfo{
				UploadId:     uploadId,
				ChunkCount:   countSum,
				ChunkSize:    config.CHUNK_Size,
				Pid:     pid,
				Filesha1:     filesha1,
				FileName:     filename,
				FileSize:     filesize,
				SuccessIndex: SuccessIndex,
			})
		} else {
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, "文件分块合并失败")
		}

	}

}

