package handler

import (
	"github.com/skydrive/beans"
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/response"
	"net/http"
	"strconv"
)

//获取所有用户
func GetAllUserListHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	if pageSize == 0 {
		pageSize = 10
	}
	if alluser, err, total := db.AdminGetAllUserList(pageNo, pageSize); err == nil {
		metaFilelist := make([]beans.User, 0)
		for _, value := range alluser {
			metaFilelist = append(metaFilelist, *beans.GetUserObject(value))
		}
		response.ReturnResponsePage(w, config.Net_SuccessCode, config.Success, metaFilelist, pageNo, pageSize, 0, total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.Error)
}

func GetAllUserTokenListHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	if pageSize == 0 {
		pageSize = 10
	}
	if alluser, err, total := db.AdminGetAllUserTokenList(pageNo, pageSize); err == nil {
		metaFilelist := make([]beans.UToken, 0)
		for _, value := range alluser {
			metaFilelist = append(metaFilelist, *beans.GetUserTokenObject(value))
		}
		response.ReturnResponsePage(w, config.Net_SuccessCode, config.Success, metaFilelist, pageNo, pageSize, 0, total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.Error)
}

func GetAllFileListHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	if pageSize == 0 {
		pageSize = 10
	}
	if alluser, err, total := db.AdminGetAllFileList(pageNo, pageSize); err == nil {
		metaFilelist := make([]beans.File, 0)
		for _, value := range alluser {
			metaFilelist = append(metaFilelist, *beans.GetFileObject(value))
		}
		response.ReturnResponsePage(w, config.Net_SuccessCode, config.Success, metaFilelist, pageNo, pageSize, 0, total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.Error)
}

func GetAllUserFileListHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	pageNo, _ := strconv.ParseInt(r.FormValue("pageNo"), 10, 64)
	pageSize, _ := strconv.ParseInt(r.FormValue("pageSize"), 10, 64)
	if pageSize == 0 {
		pageSize = 10
	}
	if alluser, err, total := db.AdminGetAllUserFileInfoList(pageNo, pageSize); err == nil {
		metaFilelist := make([]beans.UserFile, 0)
		for _, value := range alluser {
			metaFilelist = append(metaFilelist, *beans.GetUserFileObject(value))
		}
		response.ReturnResponsePage(w, config.Net_SuccessCode, config.Success, metaFilelist, pageNo, pageSize, 0, total)
		return
	}
	response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, config.Error)

}
