package handler

import (
	"fmt"
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/handler/cache"
	"github.com/skydrive/media"
	"github.com/skydrive/response"
	"github.com/skydrive/utils"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

//通过文件sha1获取文件的详细信息
func GetFileInfoBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	if r.Method == "GET" {
		r.ParseForm()
		sha1 := r.FormValue("sha1")
		if metaInfo, ok := cache.GetFileMeta(sha1); ok {
			response.ReturnMetaInfo(w, config.Net_SuccessCode, "file save success", metaInfo)
			return
		}
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "empty")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internel server error"))
	}
}

//下载文件
func OpenFile1Handler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	q, _ := strconv.ParseInt(r.FormValue("q"), 10, 64)
	if len(filesha1) == 0 || filesha1 == "" || len(filesha1) == 0 || filesha1 == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	data, ok := cache.GetFileMeta(filesha1)
	if !ok {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	//只针对图片压缩
	setHeaderFileName(w, data.FileName,nil)
	if q != 0 && data.Ftype == 0 {
		err, target := utils.CreateThumbDir(config.ThumbnailRoot, filesha1, strconv.FormatInt(q, 10), data.FileName)
		if err == nil {
			//_, error := os.Open(data.FileLocation)
			exists, _ := utils.PathExists(target)
			if !exists && media.ScaleImageQualityV1(data.FileLocation, target, config.Thumbnail_Quality) {
				http.ServeFile(w, r, target)
			} else {
				http.ServeFile(w, r, target)
			}
			return
		}
	}
	http.ServeFile(w, r, data.FileLocation)
}

func OpenFile1HandlerV2(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	q, _ := strconv.ParseInt(r.FormValue("q"), 10, 64)
	widthf, _ := strconv.ParseFloat(r.FormValue("widthf"), 10)
	width, _ := strconv.ParseInt(r.FormValue("width"), 10, 64)
	if len(filesha1) == 0 || filesha1 == "" || len(filesha1) == 0 || filesha1 == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	data, ok := cache.GetFileMeta(filesha1)
	if !ok {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	//只针对图片压缩
	setHeaderFileName(w, data.FileName, nil)
	if q != 0 && data.Ftype == 0 {
		err, target := utils.CreateThumbDir(config.ThumbnailRoot, filesha1, strconv.FormatInt(q, 10), data.FileName)
		if err == nil {
			//_, error := os.Open(data.FileLocation)
			exists, _ := utils.PathExists(target)
			if !exists {
				media.ScaleImageByWidthAndQuity(data.FileLocation, int(width), widthf, config.Thumbnail_Quality, target)
			}
			exists, _ = utils.PathExists(target)
			if exists {
				http.ServeFile(w, r, target)
				return
			}
			//ScaleImageByWidthAndQuity(path, 0, 0.5, 100, output_path)
			//if !exists && media.ScaleImageByWidthAndQuity(data.FileLocation, int(width),widthf,config.Thumbnail_Quality,target) {
			//	http.ServeFile(w, r, target)
			//} else {
			//	http.ServeFile(w, r, target)
			//}
		}
	}
	http.ServeFile(w, r, data.FileLocation)
}

func setHeaderFileName(w http.ResponseWriter, fileName string,file *os.File) {
	ctype := mime.TypeByExtension(filepath.Ext(fileName))
	if ctype == "" && file!=nil {
		// read a chunk to decide between utf-8 text and binary
		var buf [512]byte
		n, _ := io.ReadFull(file, buf[:])
		ctype = http.DetectContentType(buf[:n])
		file.Seek(0, io.SeekStart) // rewind to output whole file
	}
	w.Header().Set("Content-Type", ctype)
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", fileName))
	/*ctype := mime.TypeByExtension(filepath.Ext(data.FileName))
	if ctype == "" {
		// read a chunk to decide between utf-8 text and binary
		var buf [512]byte
		n, _ := io.ReadFull(file, buf[:])
		ctype = http.DetectContentType(buf[:n])
		file.Seek(0, io.SeekStart) // rewind to output whole file

	}
	w.Header().Set("Content-Type", ctype)
	// w.Header().Set("Content-Type", "application/octect-stream")
	// w.Header().Set("Content-Description", "attachment; filename=\""+url.QueryEscape(data.FileName)+"\";charset=UTF-8")
	// w.Header().Set("Content-Description", fmt.pr("attachment;filename=%s", data.FileName))
	// w.Header().Set("Content-Description", "attachment;filename=\""+data.FileName+"\"; charset=UTF-8")
	// w.Header().Set("Content-Description", "attachment; filename* = UTF-8''"+url.QueryEscape(data.FileName))
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", data.FileName))*/
}

//获取文件信息,暂时不支持
func UpdateFileInfoFileNameBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	newfilename := r.FormValue("filename")
	if len(filesha1) == 0 || filesha1 == "" || len(filesha1) == 0 || filesha1 == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	data, ok :=  cache.GetFileMeta(filesha1)
	if !ok {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	data.FileName = newfilename
	cache.AddOrUpdateFileMeta(*data)
	response.ReturnMetaInfo(w, config.Net_SuccessCode, "update file "+data.FileName+" success ", data)
}

//删除文件
func DeleteFileInfoBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	if len(filesha1) == 0 || filesha1 == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	data, ok := cache.GetFileMeta(filesha1)
	if !ok {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	/*
		真正删除本地文件
		if error := os.Remove(data.Location); error != nil {
			ReturnResponseCodeMessage(w, Net_ErrorCode, "delete file error "+filesha1)
			return
		}*/

	if db.UpdateFileInfoStatusBySha1(filesha1, 0) {
		response.ReturnResponseCodeMessage(w, config.Net_SuccessCode, "delete file "+data.FileName+" success ")
	} else {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "delete file error:"+filesha1)
	}
}

//浏览器打开直接下载文件
func DownloadFileWebBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	if len(filesha1) == 0 || filesha1 == "" || len(filesha1) == 0 || filesha1 == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.FormValueError)
		return
	}
	data, ok := cache.GetFileMeta(filesha1)
	if !ok {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	file, error := os.Open(data.FileLocation)
	if error != nil {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "open  file error:"+error.Error())
		return
	}
	byteData, error := ioutil.ReadAll(file)
	if error != nil {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "read file error:"+error.Error())
		return
	}
	// w.Header().Set("Content-Type", "application/octect-stream")
	// w.Header().Set("Content-Description", "attachment;filename=\""+fm.FileName+"\"")
	setHeaderFileName(w,data.FileName,file)
	w.Write(byteData)

}
