package handler

import (
	"fmt"
	"github.com/skydrive/config"
	"github.com/skydrive/db"
	"github.com/skydrive/meta"
	"github.com/skydrive/response"
	"github.com/skydrive/utils"
	"io"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"time"
)


//登录
func SignInHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
	}
	r.ParseForm()
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	if len(phone) == 0 || phone == "" || len(password) == 0 || password == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	if ok, _ := regexp.MatchString(config.Regex_MobilePhone, phone); !ok {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "手机号格式错误")
		return
	}
	if info, err := db.GetUserInfoByPhone(phone); err == nil && info.Id.Int32 > 0 {
		if BuildEncodePwd(password) == info.User_pwd.String {
			//todo 生成token
			if token, error := db.CreateUserTokenByUidPhone(info.Id.Int32, info.Phone.String); error == nil {
				response.ReturnResponse(w, config.Net_SuccessCode, "登录成功", token)
			} else {
				response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "登陆失败")
			}
		} else {
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户名或密码错误")
		}
	} else {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户未注册")
	}
}

//注册
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(32 << 20)
	}
	r.ParseForm()
	phone := r.FormValue("phone")
	password := r.FormValue("password")
	password2 := r.FormValue("password2")
	if password != password2 {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "两次密码不一致")
		return
	}
	if len(phone) == 0 || phone == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	// 验证手机号格式
	if ok, _ := regexp.MatchString(config.Regex_MobilePhone, phone); !ok {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "手机号格式错误")
		return
	}
	//是否重复注册
	if info, err := db.GetUserInfoByPhone(phone); err == nil && info.Id.Int32 > 0 {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户已经注册")
		return
	}
	if db.SaveUserInfo(phone, BuildEncodePwd(password), time.Now().Format("2006-01-02 15:04:05")) {
		if info, err := db.GetUserInfoByPhone(phone); err == nil {
			response.ReturnResponse(w, config.Net_SuccessCode, "注册成功", &User{
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
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "注册失败2")
		}
	} else {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "注册失败")
	}
}

//获取用户信息
func GetUserInfoByTokenHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	token := getToken(r)

	if len(token) == 0 || token == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	if info, err := db.GetUserInfoByPhone(utoken.Phone.String); err == nil && info.Id.Int32 > 0 {
		response.ReturnResponse(w, config.Net_SuccessCode, "获取成功", &User{
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
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode_Token_exprise, "获取用户信息失败")
	}
}

//登出
func SignOutHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	token := getToken(r)
	if len(token) == 0 || token == "" {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		return
	}
	if db.DeleteUserTokenByTid(utoken.Tid.Int64) {
		response.ReturnResponseCodeMessage(w, config.Net_SuccessCode, "登出成功")
	} else {
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "登出失败")
	}
}

func UpdateUserNameByUidHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
	r.ParseForm()
	if r.Method == "POST" {
		//uid, _ := strconv.ParseInt(r.FormValue("uid"), 10, 64)
		value := r.FormValue("name")
		//if uid == 0 || value == "" {
		//	ReturnResponseCodeMessage(w, config.Net_ErrorCode, "参数不合法")
		//	return
		//}
		if db.UpdateUserNameByUid(value, utoken.Uid.Int64) {
			response.ReturnResponseCodeMessage(w, config.Net_SuccessCode, "修改成功")
			return
		}
		response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "修改失败")
	}

}

func UpdataUploadUserPhotoHandler(w http.ResponseWriter, r *http.Request, utoken *db.TableUToken) {
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
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, fmt.Sprintf("获取文件出错 %s \n", error.Error()))
			return
		}
		error, path := utils.CreateDirbySha1( sha1, fileheader.Filename,utoken.Uid.Int64)
		if error!=nil{
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, fmt.Sprintf("创建文件夹出错 %s \n", error.Error()))
		}
		metaInfo := UserFile{
			FileName:     fileheader.Filename,
			Location:    	path,
			UpdateAtTime: time.Now().Format("2006-01-02 15:04:05"),
		}
		newfile, error := os.Create(metaInfo.Location)
		if error != nil {
			fmt.Printf("创建文件出错 %s \n", error.Error())
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "file create error")
			return
		}
		defer newfile.Close()
		metaInfo.FileSize, error = io.Copy(newfile, file)
		if error != nil {
			fmt.Printf("保存文件出错 %s \n", error.Error())
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "file copy error")
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
				response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "系统文件保存失败")
				return
			}
		}

		//文件表已经插入成功,再插入用户文件表
		if db.UpdateUserPhotoByUid(metaInfo.Location, metaInfo.Filesha1, utoken.Uid.Int64) {
			fmt.Println(" metaInfo: ", metaInfo)
			response.ReturnMetaInfo(w, config.Net_SuccessCode, "file save success", &metaInfo)
		} else {
			response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "用户文件保存失败")
		}
	}
}
