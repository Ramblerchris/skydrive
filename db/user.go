package db

import (
	"fmt"
	mysql "github.com/skydrive/db/mysqlconn"
)

const selectUserInfo = "select id ,user_name,user_pwd,photo_addr,photo_file_sha1,email,phone,email_validated,phone_validated,signup_at,last_active,profile,status from tbl_user where phone=? and status=1 limit 1"
const saveUserinfo = "insert into tbl_user(user_pwd,phone,signup_at,status,user_name) values(?,?,?,?,?)"
//const saveUserinfo = "insert into tbl_user(user_pwd,phone,signup_at,status) values(?,?,?,?)"
const updateUserPhoto = "UPDATE tbl_user SET photo_addr=? ,photo_file_sha1=? where id=?"
const updateUserName = "UPDATE tbl_user SET user_name=? where id=?"

func (t *TableUser) String() {
	fmt.Printf("Id:%s User_name:%s User_pwd :%s Email:%s Phone:%s Email_validated:%s Phone_validated:%s Signup_at :%s Last_active:%s Profile:%s Status:%s",t.Id,t.User_name,t.User_pwd,t.Email,t.Phone,t.Email_validated,t.Phone_validated,t.Signup_at,t.Last_active,t.Profile,t.Status)
}

func SaveUserInfo(phone string, password string, time string) bool {
	stmt, error := mysql.DbConnect().Prepare(saveUserinfo)
	if error != nil {
		fmt.Println("failed to prepare statement error:", error.Error())
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(password, phone, time, 1, phone)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
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
		fmt.Println("failed to prepare statement error:", error)
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
	if error != nil {
		fmt.Println("failed to QueryRow error:", error)
		return nil, error
	}
	fmt.Println("GetUserInfoByPhone:",tUser)
	return &tUser, nil
}

func UpdateUserPhotoByUid(photoAddr string,filesha1 string, uid  int64) bool {
	stmt, error := mysql.DbConnect().Prepare(updateUserPhoto)
	if error != nil {
		fmt.Println("failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(photoAddr,filesha1, uid)
	if error != nil {
		fmt.Println(tAG_userfile, "failed Exec error:", error)
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
		fmt.Println("failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(userName, uid)
	if error != nil {
		fmt.Println(tAG_userfile, "failed Exec error:", error)
		return false
	}
	if _, error := exec.RowsAffected(); error == nil {
		return true
	}
	return false
}
