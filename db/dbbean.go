package db

import "database/sql"

//这里的sql.NullString 是对象，不能直接用于json序列化
type TableFile struct {
	Id             sql.NullInt64
	Filesha1       sql.NullString
	FileName       sql.NullString
	FileLocation   sql.NullString
	Create_at      sql.NullString
	FileSize       sql.NullInt64
	Minitype       sql.NullString
	Ftype          sql.NullInt32
	Video_duration sql.NullString
}


type TableUser struct {
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


type TableUserFile struct {
	Id             sql.NullInt64
	PId            sql.NullInt64
	Uid            sql.NullInt64
	Phone          sql.NullString
	FileHash       sql.NullString
	FileHash_Pre   sql.NullString
	FileName       sql.NullString
	FileLocation   sql.NullString
	FileSize       sql.NullInt64
	Status         sql.NullInt32
	Filetype       sql.NullInt32
	Create_at      sql.NullString
	Update_at      sql.NullString
	Minitype       sql.NullString
	Ftype          sql.NullInt32
	Video_duration sql.NullString
}


type TableUToken struct {
	Tid        sql.NullInt64
	Uid        sql.NullInt64
	Phone      sql.NullString
	User_token sql.NullString
	Expiretime sql.NullInt64
}
