package db

import (
	"database/sql"
	"fmt"
	mysql "github.com/skydrive/db/mysqlconn"
)

const selectUserInfo = "select id ,user_name,user_pwd,photo_addr,photo_file_sha1,email,phone,email_validated,phone_validated,signup_at,last_active,profile,status from tbl_user where phone=? and status=1 limit 1"
const saveUserinfo = "insert into tbl_user(user_pwd,phone,signup_at,status,user_name) values(?,?,?,?,?)"
//const saveUserinfo = "insert into tbl_user(user_pwd,phone,signup_at,status) values(?,?,?,?)"
const updateUserPhoto = "UPDATE tbl_user SET photo_addr=? ,photo_file_sha1=? where id=?"
const updateUserName = "UPDATE tbl_user SET user_name=? where id=?"

type TabUser struct {
	Id              sql.NullInt32  `db:"id"`
	User_name       sql.NullString `db:"user_name"`
	User_pwd        sql.NullString `db:"user_pwd"`
	Email           sql.NullString `db:"email"`
	Phone           sql.NullString `db:"phone"`
	Photo_addr      sql.NullString `db:"photo_addr"`
	Photo_addr_sha1 sql.NullString `db:"photo_file_sha1"`
	Email_validated sql.NullInt32  `db:"email_validated"`
	Phone_validated sql.NullInt32  `db:"phone_validated"`
	Signup_at       sql.NullString `db:"signup_at"`
	Last_active     sql.NullString `db:"last_active"`
	Profile         sql.NullString `db:"profile"`
	Status          sql.NullInt32  `db:"status"`
}

func (t *TabUser) String() {
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

func GetUserInfoByPhone(phone string) (*TabUser, error) {
	stmt, error := mysql.DbConnect().Prepare(selectUserInfo)
	if error != nil {
		fmt.Println("failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	tUser := TabUser{}
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
