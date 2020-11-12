package handler

import (
	"fmt"
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/meta"
	"github.com/skydrive/utils"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)

type User struct {
	Id int32 `json:"id"`
	//`json:"-"` 字段不暴露给用户
	User_pwd        string `json:"-"`
	User_name       string `json:"user_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Photo_addr      string `json:"photo_addr"`
	Photo_addr_sha1      string `json:"photo_file_sha1"`
	Email_validated int32  `json:"email_validated"`
	Phone_validated int32  `json:"phone_validated"`
	Signup_at       string `json:"signup_at"`
	Last_active     string `json:"last_active"`
	//`json:"omitempty"`当字段为空时忽略此字段 不需要该字段返回时，让其赋值为空即可。
	Profile string `json:"profile,omitempty"`
	Status  int32  `json:"status"`
}

//登录
func Signin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
	}
	r.ParseForm()
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	if len(phone) == 0 || phone == "" || len(password) == 0 || password == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	if ok, _ := regexp.MatchString(config.Regex_MobilePhone, phone); !ok {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "手机号格式错误")
		return
	}
	if info, err := db.GetUserInfo(phone); err == nil && info.Id.Int32 > 0 {
		if BuildEncodePwd(password) == info.User_pwd.String {
			//todo 生成token
			if token, error := db.SaveUToken(info.Id.Int32, info.Phone.String); error == nil {
				ReturnResponse(w, config.Net_SuccessCode, "登录成功", token)
			} else {
				ReturnResponseCodeMessage(w, config.Net_ErrorCode, "登陆失败")
			}
		} else {
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户名或密码错误")
		}
	} else {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户未注册")
	}
}

//注册
func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
	}
	r.ParseForm()
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")
	if password != password2 {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "两次密码不一致")
		return
	}
	if len(phone) == 0 || phone == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	// 验证手机号格式
	if ok, _ := regexp.MatchString(config.Regex_MobilePhone, phone); !ok {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "手机号格式错误")
		return
	}
	//是否重复注册
	if info, err := db.GetUserInfo(phone); err == nil && info.Id.Int32 > 0 {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户已经注册")
		return
	}
	if db.SaveUserInfo(phone, BuildEncodePwd(password), time.Now().Format("2006-01-02 15:04:05")) {
		if info, err := db.GetUserInfo(phone); err == nil {
			ReturnResponse(w, config.Net_SuccessCode, "注册成功", &User{
				Id:              info.Id.Int32,
				User_name:       info.User_name.String,
				Email:           info.Email.String,
				Phone:           info.Phone.String,
				Email_validated: info.Email_validated.Int32,
				Phone_validated: info.Phone_validated.Int32,
				Signup_at:       info.Signup_at.String,
				Last_active:     info.Last_active.String,
				Profile:         info.Profile.String,
				Status:          info.Status.Int32,
			})
		} else {
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, "注册失败2")
		}
	} else {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "注册失败")
	}
}

//获取用户信息
func GetUserInfo(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	token := getToken(r)

	if len(token) == 0 || token == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	if info, err := db.GetUserInfo(utoken.Phone.String); err == nil && info.Id.Int32 > 0 {
		ReturnResponse(w, config.Net_SuccessCode, "获取成功", &User{
			Id:              info.Id.Int32,
			User_name:       info.User_name.String,
			Email:           info.Email.String,
			Phone:           info.Phone.String,
			Photo_addr:      info.Photo_addr.String,
			Photo_addr_sha1: info.Photo_addr_sha1.String,
			Email_validated: info.Email_validated.Int32,
			Phone_validated: info.Phone_validated.Int32,
			Signup_at:       info.Signup_at.String,
			Last_active:     info.Last_active.String,
			Profile:         info.Profile.String,
			Status:          info.Status.Int32,
		})
	} else {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode_Token_exprise, "获取用户信息失败")
	}
}

//登出
func SignOut(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	token := getToken(r)
	if len(token) == 0 || token == "" {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	if db.DeleteUTokenById(utoken.Tid.Int64) {
		ReturnResponseCodeMessage(w, config.Net_SuccessCode, "登出成功")
	} else {
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "登出失败")
	}
}

func UploadUserNameHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	if r.Method == "POST" {
		//uid, _ := strconv.ParseInt(r.FormValue("uid"), 10, 64)
		value := r.FormValue("name")
		//if uid == 0 || value == "" {
		//	ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		//	return
		//}
		if db.UpdateUserName(value, utoken.Uid.Int64) {
			ReturnResponseCodeMessage(w, config.Net_SuccessCode, "修改成功")
			return
		}
		ReturnResponseCodeMessage(w, config.Net_ErrorCode, "修改失败")
	}

}

func UploadUserPhotoHandler(w http.ResponseWriter, r *http.Request, utoken *db.UToken) {
	r.ParseForm()
	if r.Method == "POST" {
		file, fileheader, error := r.FormFile("file")
		sha1 := r.FormValue("sha1")
		minetype := r.FormValue("minetype")
		isVideo, _ := strconv.ParseBool(r.FormValue("isVideo"))
		videoduration, _ := strconv.ParseInt(r.FormValue("videoduration"), 10, 64)
		ftype := utils.GetFType(minetype, isVideo)
		if error != nil {
			fmt.Printf("获取文件出错 %s \n", error.Error())
			//ReturnResponseCodeMessage(w, config.Net_ErrorCode, "internel server error ")
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, fmt.Sprintf("获取文件出错 %s \n", error.Error()))
			return
		}
		error, path := utils.CreateDirbySha1( sha1, fileheader.Filename,utoken.Uid.Int64)
		if error!=nil{
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, fmt.Sprintf("创建文件夹出错 %s \n", error.Error()))
		}
		metaInfo := meta.FileMeta{
			FileName:     fileheader.Filename,
			Location:    	path,
			UpdateAtTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		newfile, error := os.Create(metaInfo.Location)
		if error != nil {
			fmt.Printf("创建文件出错 %s \n", error.Error())
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, "file create error")
			return
		}
		defer newfile.Close()
		metaInfo.FileSize, error = io.Copy(newfile, file)
		if error != nil {
			fmt.Printf("保存文件出错 %s \n", error.Error())
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, "file copy error")
			return
		}
		metaInfo.Filesha1 = utils.GetFileSha1(newfile)
		fmt.Println("file sha1", metaInfo.Filesha1)
		meta.AddOrUpdateFileMeta(metaInfo)
		//处理文件已经存在的情况
		_, ok := meta.GetFileMeta(metaInfo.Filesha1)
		if !ok {
			//如果不存在 先插入文件表
			if !db.SaveFileInfo(metaInfo.Filesha1, metaInfo.FileName, metaInfo.FileSize, metaInfo.Location, minetype, ftype, videoduration) {
				//插入文件表不成功
				ReturnResponseCodeMessage(w, config.Net_ErrorCode, "系统文件保存失败")
				return
			}
		}

		//文件表已经插入成功,再插入用户文件表
		if db.UpdateUserPhoto(metaInfo.Location, metaInfo.Filesha1, utoken.Uid.Int64) {
			fmt.Println(" metaInfo: ", metaInfo)
			ReturnMetaInfo(w, config.Net_SuccessCode, "file save success", &metaInfo)
		} else {
			ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户文件保存失败")
		}
	}
}
