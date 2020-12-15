package handler

import (
	"fmt"
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/meta"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

//通过文件sha1获取文件的详细信息
func GetFileInfoBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	if r.Method == "GET" {
		r.ParseForm()
		sha1 := r.FormValue("sha1")
		if metaInfo, ok := meta.GetFileMeta(sha1); ok {
			ReturnMetaInfo(w, config.Net_SuccessCode, "file save success", metaInfo)
			return
		}
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "empty")
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("internel server error"))
	}
}

//下载文件
func OpenFile1Handler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	if len(filesha1) == 0 || filesha1 == "" || len(filesha1) == 0 || filesha1 == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	data, ok := meta.GetFileMeta(filesha1)
	if !ok {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	http.ServeFile(w, r, data.Location)
}

//获取文件信息,暂时不支持
func UpdateFileInfoFileNameBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	newfilename := r.FormValue("filename")
	if len(filesha1) == 0 || filesha1 == "" || len(filesha1) == 0 || filesha1 == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	data, ok := meta.GetFileMeta(filesha1)
	if !ok {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	data.FileName = newfilename
	meta.AddOrUpdateFileMeta(*data)
	ReturnMetaInfo(w, config.Net_SuccessCode, "update file "+data.FileName+" success ", data)
}

//删除文件
func DeleteFileInfoBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	if len(filesha1) == 0 || filesha1 == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	data, ok := meta.GetFileMeta(filesha1)
	if !ok {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	/*
		真正删除本地文件
		if error := os.Remove(data.Location); error != nil {
			ReturnResponseCodeMessage(w, Net_ErrorCode, "delete file error "+filesha1)
			return
		}*/
	if meta.RemoveFileMeta(filesha1) {
		ReturnResponseCodeMessage(w, config.Net_SuccessCode, "delete file "+data.FileName+" success ")
	} else {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "delete file error:"+filesha1)
	}

}

//浏览器打开直接下载文件
func DownloadFileWebBySha1Handler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	filesha1 := r.FormValue("filesha1")
	if len(filesha1) == 0 || filesha1 == "" || len(filesha1) == 0 || filesha1 == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	data, ok := meta.GetFileMeta(filesha1)
	if !ok {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "not find filesha1:"+filesha1)
		return
	}
	file, error := os.Open(data.Location)
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "open  file error:"+error.Error())
		return
	}
	byteData, error := ioutil.ReadAll(file)
	if error != nil {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "read file error:"+error.Error())
		return
	}
	// w.Header().Set("Content-Type", "application/octect-stream")
	// w.Header().Set("Content-Description", "attachment;filename=\""+fm.FileName+"\"")
	ctype := mime.TypeByExtension(filepath.Ext(data.FileName))
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
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", data.FileName))

	w.Write(byteData)

}
