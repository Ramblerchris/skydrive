package db

import (
	"fmt"
	mysql "github.com/skydrive/db/mysqlconn"
	"github.com/skydrive/utils"
)

const selectFile = "select id,file_sha1,file_name,file_size,file_addr,minitype,ftype,video_duration from tbl_file where file_sha1=? and status=1 limit 1"
const saveFileinfo = "insert into tbl_file(file_sha1,file_name,file_size,file_addr,status,minitype,ftype,video_duration) values(?,?,?,?,?,?,?,?)"
const updateFileInfo = "update tbl_file set status=? where file_sha1=?"
const tAG_fileInfo ="file.go"

func SaveFileInfo(filehash string, filename string, filesize int64, location string,minitype string ,ftype int ,video_duration int64) bool {
	//stmt ,error:=mysql.DbConnect().Prepare(saveFile)
	stmt, error := mysql.DbConnect().Prepare(saveFileinfo)
	if error != nil {
		fmt.Println(tAG_fileInfo,"failed to prepare statement error:", error.Error())
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(filehash, filename, filesize, location, 1,minitype,ftype,utils.GetTimeStr(int(video_duration)))
	if error != nil {
		fmt.Println(tAG_fileInfo,"failed to Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，但未添加，
			fmt.Println(tAG_fileInfo,"failed with hash has been upload", error)
			return false
		}
	}
	return true
}

func UpdateFileInfoStatusBySha1(filehash string, filestatus int8) bool {
	//stmt ,error:=mysql.DbConnect().Prepare(saveFile)
	stmt, error := mysql.DbConnect().Prepare(updateFileInfo)
	if error != nil {
		fmt.Println(tAG_fileInfo,"failed to prepare statement error:", error.Error())
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(filestatus, filehash)
	if error != nil {
		fmt.Println(tAG_fileInfo,"failed Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，修改未成功，
			fmt.Println(tAG_fileInfo,"failed with hash has been upload", error)
			return false
		}
	}
	return true
}


func GetFileInfoBySha1(filehash string) (*TableFile, error) {
	stmt, error := mysql.DbConnect().Prepare(selectFile)
	if error != nil {
		fmt.Println(tAG_fileInfo,"failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	tfile := TableFile{}
	//file_sha1,file_name,file_size,file_addr
	error =stmt.QueryRow(filehash).Scan(
		&tfile.Id,&tfile.Filesha1, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
	if error != nil {
		fmt.Println(tAG_fileInfo,"failed to QueryRow error:", error)
		return nil, error
	}
	return &tfile, nil
}