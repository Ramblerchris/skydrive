package handler

import (
	"fmt"
	"github.com/skydrive/beans"
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/handler/cache"
	"github.com/skydrive/logger"
	"github.com/skydrive/response"
	"github.com/skydrive/syncfile"
	"github.com/skydrive/utils"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// GetDiskDirFileListByPidHandler 获取用户文件夹内的所有文件
func GetDiskDirFileListByPidHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	lastId, _ := strconv.ParseInt(r.FormValue("lastId"), 10, 64)
	pid, _ := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	if pid == 0 {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	if pageSize == 0 {
		pageSize = 10
	}
	if byuid, err, total := db.GetDiskFileListByUidPid(utoken.Uid.Int64, pid, pageNo, pageSize, lastId); err == nil {
		metaFilelist := make([]beans.UserFile, 0)
		for _, value := range byuid {
			//logger.Error("GetUserDirFileListByPidHandler", value)
			metaFilelist = append(metaFilelist, *beans.GetUserFileObject(value))
		}
		nextPageId := lastId
		if len(metaFilelist) != 0 {
			nextPageId = metaFilelist[len(metaFilelist)-1].Id
		}
		//ReturnResponse(w, config.Net_SuccessCode, "get file success ", metaFilelist)
		response.ReturnResponsePage(w, config.Net_SuccessCode, config.Success, metaFilelist, pageNo, pageSize, nextPageId, total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.Error)
}

// DeleteDiskFileDirByUidHandler 批量删除指定文件夹
func DeleteDiskFileDirByUidHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	if r.Method == "POST" {
		//id, _ := strconv.ParseInt(r.FormValue("ids"), 10, 64)
		value := r.FormValue("ids")
		if value == "" {
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
			return
		}
		split := strings.Split(value, ";")
		if success,row:=db.UpdateUserDiskFileDirStatusByIds(split, -1);success {
			response.ReturnResponseCodeMessage(w, config.Net_SuccessCode, fmt.Sprintf("修改成功%d行",row))
		} else {
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.ErrorAlreadyupdate)
		}
	}
}

// AddDiskFileDirByUidPidHandler 创建文件夹
func AddDiskFileDirByUidPidHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	dirname := r.FormValue("filename")
	pid, _ := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	if len(dirname) == 0 || dirname == "" || pid == 0 {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	//todo 查询用户同级文件夹下的文件名是否存在
	if data, err := db.GetDiskFileInfoListByUidPidDirName(utoken.Uid.Int64, pid, dirname); err == nil && len(data) != 0 {
		//response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, dirname+"文件夹已经存在")
		//return
	}
	if isok, id := db.SaveDiskDirInfo(utoken.Uid.Int64, pid, utoken.Phone.String, dirname); isok {
		if value, err := db.GetDiskFileInfoById(id); err == nil {
			response.ReturnResponse(w, config.Net_SuccessCode, config.Success, *beans.GetUserFileObject(*value))
			return
		}
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.CreateError)
}

// HitPassDiskBySha1Handler 文件通过sha1 秒传
func HitPassDiskBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	sha1 := r.FormValue("sha1")
	pid, _ := strconv.ParseInt(r.FormValue("pid"), 10, 64)
	if len(sha1) == 0 || sha1 == "" || pid == 0 {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	if metaInfo, err := db.GetFileInfoBySha1(sha1); err == nil {
		//if metaInfo, err := db.GetUserFileMetaByPidUidSha1(sha1,utoken.Uid.Int64,pid); err == nil {
		//查看当前用户对应文件夹是否已经保存过
		if value, err := db.GetDiskFileMetaByPidUidSha1(sha1, utoken.Uid.Int64, pid); err == nil {
			//避免重复保存
			response.ReturnResponse(w, config.Net_SuccessAginCode, "already save success ", *beans.GetUserFileObject(*value))
			return
		}
		if db.SaveDiskFileInfo(utoken.Uid.Int64, pid, utoken.Phone.String, metaInfo.Filesha1.String, metaInfo.FileName.String, metaInfo.FileLocation.String, metaInfo.FileSize.Int64, metaInfo.Minitype.String, int(metaInfo.Ftype.Int32), metaInfo.Video_duration.String) {
			logger.Info(" metaInfo: ", metaInfo)
			logger.Infof("保存文件 成功，大小 %d ", metaInfo.FileSize)
			//更新当前文件夹的缩略图最新
			db.UpdateUserDiskFileDirPreSha1ById(metaInfo.Filesha1.String, pid)
			if value, err := db.GetDiskFileInfoByUidSha1(sha1, utoken.Uid.Int64); err == nil {
				response.ReturnResponse(w, config.Net_SuccessCode, config.Success, *beans.GetUserFileObject(*value))
				return
			}
		}
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户文件保存失败")
	} else {
		response.ReturnResponseCodeMessage(w, config.Net_EmptyCode, "未上传")
	}
}


//修改指定文件夹名称
func UpdateDiskDirNameById(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	newfilename := r.FormValue("newfilename")
	if newfilename == "" || len(newfilename)==0 {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	if db.UpdateDiskFileDirsNameById(newfilename, id) {
		response.ReturnResponseCodeMessage(w, config.Net_SuccessCode, config.Success)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.Error)
}

// UploadDiskFileHandler 上传文件
func UploadDiskFileHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	/*if r.Method == "GET" {
		//浏览器打开
		// data, error := ioutil.ReadFile("./static/view/index.html")
		// if error != nil {
		// 	ReturnResponseCodeMessage(w, Net_ErrorCode, "internel server error ")
		// 	return
		// }
		// // io.WriteString(w, string(data))
		// w.Write(data)
		//http.ServeFile(w, r, "./static/view/index.html")
		io.WriteString(w, string("<h1>请下载客户端<h1>"))
	} else */if r.Method == "POST" {
		sha1, _ := strconv.Unquote(r.FormValue("sha1"))
		minetype, _ := strconv.Unquote(r.FormValue("minetype"))
		file, fileheader, error := r.FormFile("file")
		pid, _ := strconv.ParseInt(r.FormValue("pid"), 10, 64)
		isVideo, _ := strconv.ParseBool(r.FormValue("isVideo"))
		videoduration, _ := strconv.ParseInt(r.FormValue("videoduration"), 10, 64)
		ftype := utils.GetFType(minetype, isVideo)
		if error != nil {
			logger.Error("获取文件出错 %s \n", error.Error())
			//ReturnResponseCodeMessage(w, config.Net_ErrorCode, "internel server error ")
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, fmt.Sprintf("获取文件出错 %s \n", error.Error()))
			return
		}
		error, path := utils.CreateDirbySha1(config.DiskFileRoot, sha1, fileheader.Filename, utoken.Uid.Int64)
		if error != nil {
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, fmt.Sprintf("创建文件夹出错 %s \n", error.Error()))
		}
		metaInfo := beans.UserFile{
			//Location:     path,
			UpdateAtTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		metaInfo.FileName = fileheader.Filename
		metaInfo.FileLocation = path
		newfile, error := os.Create(metaInfo.FileLocation)
		if error != nil {
			logger.Error("创建文件出错 %s \n", error.Error())
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "file create error")
			return
		}
		defer newfile.Close()
		metaInfo.FileSize, error = io.Copy(newfile, file)
		if error != nil {
			logger.Error("保存文件出错 %s \n", error.Error())
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "file copy error")
			return
		}
		syncfile.ChangeDirModTime(metaInfo.FileLocation,5)
		metaInfo.Filesha1 = utils.GetFileSha1(newfile)
		//logger.Error("file sha1", metaInfo.Filesha1)
		//todo 缓存添加
		//cache.AddOrUpdateFileMeta(metaInfo)
		//处理文件已经存在的情况
		_, ok := cache.GetFileMeta(metaInfo.Filesha1)
		if !ok {
			//如果不存在 先插入文件表
			if !db.SaveFileInfo(metaInfo.Filesha1, metaInfo.FileName, metaInfo.FileSize, metaInfo.FileLocation, minetype, ftype, videoduration) {
				//插入文件表不成功
				response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "系统文件保存失败")
				return
			}
		}
		//查看是否已经保存过
		if value, err := db.GetDiskFileInfoByUidSha1(metaInfo.Filesha1, utoken.Uid.Int64); err == nil {
			//避免重复保存
			response.ReturnResponse(w, config.Net_SuccessAginCode, "already save success ", *beans.GetUserFileObject(*value))
			return
		}
		//文件表已经插入成功,再插入用户文件表
		if db.SaveDiskFileInfo(utoken.Uid.Int64, pid, utoken.Phone.String, metaInfo.Filesha1, metaInfo.FileName, metaInfo.FileLocation, metaInfo.FileSize, minetype, ftype, utils.GetTimeStr(int(videoduration))) {
			//更新当前文件夹的缩略图最新
			db.UpdateUserDiskFileDirPreSha1ById(metaInfo.Filesha1, pid)
			logger.Info(" metaInfo: ", metaInfo)
			response.ReturnResponse(w, config.Net_SuccessCode, config.Success, &metaInfo)
		} else {
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.SaveFileError)
		}
	}
}
