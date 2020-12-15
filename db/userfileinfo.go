package db

import (
	"database/sql"
	"fmt"
	mysql "github.com/skydrive/db/mysqlconn"
	"strings"
)

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
	Filetype       sql.NullInt32
	Create_at      sql.NullString
	Update_at      sql.NullString
	Minitype       sql.NullString
	Ftype          sql.NullInt32
	Video_duration sql.NullString
}

const (
	selectUFileBySha1                 = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at,filetype,minitype,ftype,video_duration from tbl_user_file where file_sha1=? and uid=? and status=1 limit 1"
	selectUFileByid                   = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at,filetype,minitype,ftype,video_duration from tbl_user_file where id=?  and status=1 limit 1"
	selectUFileByUid                  = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_user_file where uid=? and status=1 "
	selectUFileByUidAndPid            = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_user_file where uid=? and pid=? and status=1 "
	selectUFileByUidAndPidAndSha1     = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_user_file where file_sha1=? and uid=? and pid=? and status=1 limit 1 "
	selectUFileByUidAndPidAndFileName = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_user_file where uid=? and pid=? and file_name=? and status=1 "
	saveUFileinfo                     = "insert into tbl_user_file(uid,pid,phone,file_sha1,file_name,file_size,file_addr,status,minitype,ftype,video_duration) values(?,?,?,?,?,?,?,?,?,?,?)"
	saveUFileDirinfo                  = "insert into tbl_user_file(uid,pid,phone,file_name,status,filetype) values(?,?,?,?,?,?)"
	updateUFileInfo                   = "update tbl_user_file set status=? where file_sha1=?"
	updateUFileStatus                 = "update tbl_user_file set status=? where uid=? and  pid=? and file_sha1 in('%s')"
	updateUFileDirStatus              = "update tbl_user_file set status=? where id=? "
	updateUFileDirfile_sha1_pre       = "update tbl_user_file set file_sha1_pre=? where id=? "
	updateUFileDirsStatus             = "update tbl_user_file set status=? where id in('%s')"
	selectUFileBysha1s                = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at,filetype,minitype,ftype,video_duration from tbl_user_file where uid=? and file_sha1 in('%s') "
	selectUFileAllsha1s               = "select file_sha1 from tbl_user_file where uid=?"
	tAG_userfile                      = "userfileinfo.go"
)

//创建一个文件夹
func SaveUserDirInfo(uid, pid int64, phone, dirName string) (bool, int64) {
	//stmt ,error:=mysql.DbConnect().Prepare(saveFile)
	stmt, error := mysql.DbConnect().Prepare(saveUFileDirinfo)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error.Error())
		return false, -1
	}
	defer stmt.Close()
	if pid == 0 {
		pid = -1
	}
	//filetype 文件夹1 文件-1
	exec, error := stmt.Exec(uid, pid, phone, dirName, 1, 1)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to Exec error:", error)
		return false, -1
	}
	lastinsertId, _ := exec.LastInsertId()

	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，但未添加，
			fmt.Printf(tAG_userfile, "failed with hash has been upload", error)
			return false, -1
		}
	}
	return true, lastinsertId
}

//保存文件信息到用户对应文件夹
func SaveUserFileInfo(uid, pid int64, phone, filehash, filename, location string, filesize int64, minitype string, ftype int, video_duration string) bool {
	stmt, error := mysql.DbConnect().Prepare(saveUFileinfo)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error.Error())
		return false
	}
	defer stmt.Close()
	if pid == 0 {
		pid = -1
	}
	exec, error := stmt.Exec(uid, pid, phone, filehash, filename, filesize, location, 1, minitype, ftype, video_duration)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，但未添加，
			fmt.Printf(tAG_userfile, "failed with hash:%s has been upload", error)
			return false
		}
	}
	return true
}

//修改所有用户的这个文件状态
func UpdateUserFileInfoStatusBySha1(filehash string, filestatus int8) bool {
	stmt, error := mysql.DbConnect().Prepare(updateUFileInfo)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error.Error())
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(filestatus, filehash)
	if error != nil {
		fmt.Println(tAG_userfile, "failed Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，修改未成功，
			fmt.Printf(tAG_userfile, "failed with hash:%s has been upload", error)
			return false
		}
	}
	return true
}

//查看单个文件具体信息
func GetUserDirInfoById(id int64) (*TableUserFile, error) {
	stmt, error := mysql.DbConnect().Prepare(selectUFileByid)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	tfile := TableUserFile{}
	error = stmt.QueryRow(id).Scan(
		&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype,&tfile.Minitype,&tfile.Ftype,&tfile.Video_duration)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to QueryRow error:", error)
		return nil, error
	}
	return &tfile, nil
}

//查看当前用户是否已经存储filehash对应的文件
func GetUserFileMetaByPidUidSha1(filehash string, uid, pid int64) (*TableUserFile, error) {
	stmt, error := mysql.DbConnect().Prepare(selectUFileByUidAndPidAndSha1)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	tfile := TableUserFile{}
	row := stmt.QueryRow(filehash, uid, pid)
	error =row.Scan(
		&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype,&tfile.Minitype,&tfile.Ftype,&tfile.Video_duration)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to QueryRow error:", error)
		return nil, error
	}
	return &tfile, nil
}

//查看当前用户是否已经存储filehash对应的文件
func GetUserFileInfoByUidSha1(filehash string, uid int64) (*TableUserFile, error) {
	stmt, error := mysql.DbConnect().Prepare(selectUFileBySha1)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	tfile := TableUserFile{}
	error = stmt.QueryRow(filehash, uid).Scan(
		&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype,&tfile.Minitype,&tfile.Ftype,&tfile.Video_duration)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to QueryRow error:", error)
		return nil, error
	}
	return &tfile, nil
}

//查询用户同级文件夹下相同文件名的文件夹列表
func GetUserDirListByUidPidDirName(uid, pid int64, filename string) (tableUserFile []TableUserFile, err error) {
	stmt, error := mysql.DbConnect().Prepare(selectUFileByUidAndPidAndFileName)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	rowdata, error := stmt.Query(uid, pid, filename)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
		return tableUserFile, error
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error = rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype,&tfile.Minitype,&tfile.Ftype,&tfile.Video_duration)
		if error != nil {
			fmt.Println(tAG_userfile, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}
	return tableUserFile, nil
}

//查看当前用户 pid 对应子目录所有文件列表，包括文件列表
func GetUserDirListByUidPid(uid, pid int64) (tableUserFile []TableUserFile, err error) {
	stmt, error := mysql.DbConnect().Prepare(selectUFileByUidAndPid)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return tableUserFile, error
	}
	defer stmt.Close()
	rowdata, error := stmt.Query(uid, pid)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
		return tableUserFile, error
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error = rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype,&tfile.Minitype,&tfile.Ftype,&tfile.Video_duration)
		if error != nil {
			fmt.Println(tAG_userfile, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}
	return tableUserFile, nil
}

//查询用户所有的文件
func GetUserFileListMetaByUid(uid int64) (tableUserFile []TableUserFile, err error) {
	stmt, error := mysql.DbConnect().Prepare(selectUFileByUid)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	rowdata, error := stmt.Query(uid)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
		return tableUserFile, error
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error = rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype,&tfile.Minitype,&tfile.Ftype,&tfile.Video_duration)
		if error != nil {
			fmt.Println(tAG_userfile, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}
	return tableUserFile, nil
}

//查询用户所有的文件
func GetUserFileAllSha1ListByUid(uid int64) (sha1s []string, err error) {
	stmt, error := mysql.DbConnect().Prepare(selectUFileAllsha1s)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	rowdata, error := stmt.Query(uid)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
		return sha1s, error
	}
	sha1s = make([]string, 0)
	for rowdata.Next() {
		var sha1 string
		error = rowdata.Scan(&sha1)
		if error != nil {
			fmt.Println(tAG_userfile, "failed to Query error:", error)
			continue
		}
		if sha1 != "" {
			sha1s = append(sha1s, sha1)
		}
	}
	return sha1s, nil
}

//查询用户所有的文件
func GetUserFileListBySha1s(uid int64, sha1s []string) (tableUserFile []TableUserFile, err error) {
	sha1sJoin := strings.Join(sha1s, "','")
	sprintf := fmt.Sprintf(selectUFileBysha1s, sha1sJoin)
	stmt, error := mysql.DbConnect().Prepare(sprintf)
	fmt.Println(tAG_userfile, "GetUserFileListBySha1s sprintf:", sprintf)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	rowdata, error := stmt.Query(uid)
	if error != nil {
		fmt.Println("failed to Exec error:", error)
		return tableUserFile, error
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error = rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype,&tfile.Minitype,&tfile.Ftype,&tfile.Video_duration)
		if error != nil {
			fmt.Println(tAG_userfile, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}
	return tableUserFile, nil
}

//批量修改当前用户文件状态
func UpdateUserFileStatusBySha1sUidPid(uid, pid int64, filestatus int8, sha1s []string) bool {
	sha1sJoin := strings.Join(sha1s, "','")
	sprintf := fmt.Sprintf(updateUFileStatus, sha1sJoin)
	stmt, error := mysql.DbConnect().Prepare(sprintf)
	fmt.Println(tAG_userfile, "GetUserFileListBySha1s sprintf:", sprintf)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(filestatus,  uid,pid)
	if error != nil {
		fmt.Println(tAG_userfile, "failed Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，修改未成功，
			fmt.Printf(tAG_userfile, "failed with hash:%s has been upload", error)
			return false
		}
	}
	return true
}

//当前用户文件状态批量修改
func UpdateUserFileDirStatusByIds(ids []string, filestatus int8) bool {
	sha1sJoin := strings.Join(ids, "','")
	sprintf := fmt.Sprintf(updateUFileDirsStatus, sha1sJoin)
	stmt, error := mysql.DbConnect().Prepare(sprintf)
	fmt.Println(tAG_userfile, "GetUserFileListBySha1s sprintf:", updateUFileDirsStatus)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(filestatus)
	if error != nil {
		fmt.Println(tAG_userfile, "failed Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，修改未成功，
			fmt.Printf(tAG_userfile, "failed with hash:%s has been upload\n", error)
			return false
		}
	}
	return true
}

//当前用户文件预览图
func UpdateUserFileDirPreSha1ById(sha1 string, id int64) bool {
	stmt, error := mysql.DbConnect().Prepare(updateUFileDirfile_sha1_pre)
	fmt.Println(tAG_userfile, "UpdateUserFileDirPreSha1ById sprintf:", updateUFileDirfile_sha1_pre)
	if error != nil {
		fmt.Println(tAG_userfile, "failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(sha1, id)
	if error != nil {
		fmt.Println(tAG_userfile, "failed Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows == 1 {
			return true
		}
	}
	return false
}
