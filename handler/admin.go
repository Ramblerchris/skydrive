package handler

import (
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/meta"
	"github.com/skydrive/response"
	"net/http"
	"strconv"
)

//获取所有用户
func GetAllUserListHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	if pageSize==0 {
		pageSize=10
	}
	if alluser, err ,total := db.AdminGetAllUserList(pageNo,pageSize); err == nil {
		metaFilelist := make([]User, 0)
		for _, value := range alluser {
			metaFilelist = append(metaFilelist,*meta.GetNewUserMetaObject(value))
		}
		response.ReturnResponsePage(w, config.Net_SuccessCode, "get file success ", metaFilelist,pageNo,pageSize,0,total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "get file success ")
}

func GetAllUserTokenListHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	/*r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	if pageSize==0 {
		pageSize=10
	}
	if alluser, err ,total := db.AdminGetAllUserTokenList(pageNo,pageSize); err == nil {
		metaFilelist := make([]User, 0)
		for _, value := range alluser {
			metaFilelist = append(metaFilelist,*meta.GetNewUserMetaObject(value))
		}
		response.ReturnResponsePage(w, config.Net_SuccessCode, "get file success ", metaFilelist,pageNo,pageSize,0,total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "get file success ")*/
}

func GetAllFileListHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	/*r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	if pageSize==0 {
		pageSize=10
	}
	if alluser, err ,total := db.AdminGetAllFileList(pageNo,pageSize); err == nil {
		metaFilelist := make([]User, 0)
		for _, value := range alluser {
			metaFilelist = append(metaFilelist,*meta.get(value))
		}
		response.ReturnResponsePage(w, config.Net_SuccessCode, "get file success ", metaFilelist,pageNo,pageSize,0,total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "get file success ")*/
}

func GetAllUserFileListHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	if pageSize==0 {
		pageSize=10
	}
	if alluser, err ,total := db.AdminGetAllUserFileInfoList(pageNo,pageSize); err == nil {
		metaFilelist := make([]UserFile, 0)
		for _, value := range alluser {
			metaFilelist = append(metaFilelist,*meta.GetNewFileMetaObject(value))
		}
		response.ReturnResponsePage(w, config.Net_SuccessCode, "get file success ", metaFilelist,pageNo,pageSize,0,total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "get file success ")

}

