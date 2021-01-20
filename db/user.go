package db

import (
	mysql "github.com/skydrive/db/mysqlconn"
	"github.com/skydrive/logger"
	"github.com/skydrive/utils"
)

const (
	selectUserInfo = "select id ,user_name,user_pwd,photo_addr,photo_file_sha1,email,phone,email_validated,phone_validated,signup_at,last_active,profile,status from tbl_user where phone=? and status=1 limit 1"
	saveUserinfo = "insert into tbl_user(user_pwd,phone,signup_at,status,user_name) values(?,?,?,?,?)"
	//const saveUserinfo = "insert into tbl_user(user_pwd,phone,signup_at,status) values(?,?,?,?)"
	updateUserPhoto = "UPDATE tbl_user SET photo_addr=? ,photo_file_sha1=? where id=?"
	updateUserName  = "UPDATE tbl_user SET user_name=? where id=?"
	tAG_user             = "user.go sql:"

)

func (t *TableUser) String() {
	logger.Info("Id:%s User_name:%s User_pwd :%s Email:%s Phone:%s Email_validated:%s Phone_validated:%s Signup_at :%s Last_active:%s Profile:%s Status:%s",t.Id,t.User_name,t.User_pwd,t.Email,t.Phone,t.Email_validated,t.Phone_validated,t.Signup_at,t.Last_active,t.Profile,t.Status)
}

func SaveUserInfo(phone string, password string, time string) bool {
	stmt, error := mysql.DbConnect().Prepare(saveUserinfo)
	if error != nil {
		logger.Error("failed to prepare statement error:", error.Error())
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(password, phone, time, 1, phone)
	logger.Info(tAG_user,utils.RunFuncName(),saveUserinfo,password, phone, time, 1, phone)
	if error != nil {
		logger.Error("failed to Exec error:", error)
		return false
	}
	if _, error := exec.RowsAffected(); error == nil {
		return true
	}
	return false
}

func GetUserInfoByPhone(phone string) (*TableUser, error) {
	stmt, error := mysql.DbConnect().Prepare(selectUserInfo)
	if error != nil {
		logger.Error("failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	tUser := TableUser{}
	error = stmt.QueryRow(phone).Scan(
		&tUser.Id,
		&tUser.User_name,
		&tUser.User_pwd,
		&tUser.Photo_addr,
		&tUser.Photo_addr_sha1,
		&tUser.Email,
		&tUser.Phone,
		&tUser.Email_validated,
		&tUser.Phone_validated,
		&tUser.Signup_at,
		&tUser.Last_active,
		&tUser.Profile,
		&tUser.Status)
	logger.Info(tAG_user, utils.RunFuncName(), saveUserinfo,phone)

	if error != nil {
		logger.Error("failed to QueryRow error:", error)
		return nil, error
	}
	logger.Info("GetUserInfoByPhone:",tUser)
	return &tUser, nil
}

func UpdateUserPhotoByUid(photoAddr string,filesha1 string, uid  int64) bool {
	stmt, error := mysql.DbConnect().Prepare(updateUserPhoto)
	if error != nil {
		logger.Error("failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(photoAddr,filesha1, uid)
	logger.Info(tAG_user, utils.RunFuncName(), updateUserPhoto,photoAddr,filesha1, uid)

	if error != nil {
		logger.Error(tAG_user, "failed Exec error:", error)
		return false
	}
	if _, error := exec.RowsAffected(); error == nil {
		return true
	}
	return false
}

func UpdateUserNameByUid(userName string, uid  int64) bool {
	stmt, error := mysql.DbConnect().Prepare(updateUserName)
	if error != nil {
		logger.Error("failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(userName, uid)
	logger.Info(tAG_user, utils.RunFuncName(), updateUserName,userName, uid)

	if error != nil {
		logger.Error(tAG_user, "failed Exec error:", error)
		return false
	}
	if _, error := exec.RowsAffected(); error == nil {
		return true
	}
	return false
}
