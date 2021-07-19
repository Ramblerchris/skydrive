package db

import (
	"database/sql"
	mysql "github.com/skydrive/db/mysqlconn"
	"github.com/skydrive/logger"
	"github.com/skydrive/utils"
)


const (
	 selectAllUserInfo 			= "select id ,user_name,user_pwd,photo_addr,photo_file_sha1,email,phone,email_validated,phone_validated,signup_at,last_active,profile,status from tbl_user  limit ?,?"
	 selectAllUserFile          = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at,filetype,minitype,ftype,video_duration from tbl_user_file  limit  ?,?"
	 selectAllDiskUserFile          = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at,filetype,minitype,ftype,video_duration from tbl_cloud_disk  limit  ?,?"
	 selectAllUTokenInfo		= "select id ,uid,phone,user_token,expiretime from tbl_user_token limit  ?,? "
	 selectAllFile 				= "select id,file_sha1,file_name,file_size,file_addr,minitype,ftype,video_duration,create_at from tbl_file limit ?,?"
	 selectCountByTable	 		= "select count(*) from  "
	tAG_admin             = "admin.go sql:"

)

func AdminGetAllFileList(pageNo,pageSize int64 ) (tableFile []TableFile, err error, total int64) {
	rowdata, err := getPageStmt(selectAllFile, pageNo, pageSize)
	defer rowdata.Close()
	if err != nil {
		return tableFile, err,0
	}
	tableFile = make([]TableFile, 0)
	for rowdata.Next() {
		tfile := TableFile{}
		//file_sha1,file_name,file_size,file_addr
		error :=rowdata.Scan(
			&tfile.Id,&tfile.Filesha1, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration, &tfile.Create_at)
		if error != nil {
			logger.Error(tAG_admin, "failed to Query error:", error)
			continue
		}
		tableFile = append(tableFile, tfile)
	}
	return tableFile, nil, GetCountByTableName("tbl_file")
}

func AdminGetAllUserList(pageNo,pageSize int64 ) (tableUserlist []TableUser, err error, total int64) {
	rowdata, err := getPageStmt(selectAllUserInfo, pageNo, pageSize)
	defer rowdata.Close()
	if err != nil {
		return tableUserlist, err,0
	}
	for rowdata.Next() {
		tUser := TableUser{}
		error := rowdata.Scan(
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
			logger.Error(tAG_admin, "failed to Query error:", error)
			continue
		}
		tableUserlist = append(tableUserlist, tUser)
	}
	return tableUserlist, nil, GetCountByTableName("tbl_user")
}

func AdminGetAllUserTokenList(pageNo,pageSize int64 ) (tableUsertokenFile []TableUToken,  err error, total int64) {
	rowdata, err := getPageStmt(selectAllUTokenInfo, pageNo, pageSize)
	defer rowdata.Close()
	if err != nil {
		return tableUsertokenFile, err,0
	}
	tableUsertokenFile = make([]TableUToken, 0)
	for rowdata.Next() {
		utoken := TableUToken{}
		error:= rowdata.Scan(
			&utoken.Tid,
			&utoken.Uid,
			&utoken.Phone,
			&utoken.User_token,
			&utoken.Expiretime)
		if error != nil {
			logger.Error(tAG_admin, "failed to Query error:", error)
			continue
		}
		tableUsertokenFile = append(tableUsertokenFile, utoken)
	}
	return tableUsertokenFile, nil, GetCountByTableName("tbl_user_token")
}

func AdminGetAllDiskUserFileInfoList(pageNo,pageSize int64 ) (tableUserFile []TableUserFile, err error, total int64) {
	rowdata, err := getPageStmt(selectAllDiskUserFile, pageNo, pageSize)
	defer rowdata.Close()
	if err != nil {
		return tableUserFile, err,0
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error := rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
		if error != nil {
			logger.Error(tAG_admin, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}
	return tableUserFile, nil, GetCountByTableName("tbl_user_file")
}

func AdminGetAllUserFileInfoList(pageNo,pageSize int64 ) (tableUserFile []TableUserFile, err error, total int64) {
	rowdata, err := getPageStmt(selectAllUserFile, pageNo, pageSize)
	defer rowdata.Close()
	if err != nil {
		return tableUserFile, err,0
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error := rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
		if error != nil {
			logger.Error(tAG_admin, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}
	return tableUserFile, nil, GetCountByTableName("tbl_user_file")
}

func getPageStmt(sql string,pageNo,pageSize int64)(*sql.Rows, error)   {
	stmt, error := mysql.DbConnect().Prepare(sql)
	if error != nil {
		logger.Error(tAG_admin, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	rowdata, error := stmt.Query(pageSize*(pageNo-1),pageSize)
	logger.Info(tAG_admin, utils.RunFuncName(), sql,pageSize*(pageNo-1),pageSize)
	if error != nil {
		logger.Error("failed to Exec error:", error)
		return nil, error
	}
	return rowdata,nil
}

//查询不同表下的数量
func GetCountByTableName(tablename string ) (count int64) {
	//目前忽略了文件类型
	stmt, error := mysql.DbConnect().Prepare(selectCountByTable+tablename)
	defer stmt.Close()
	if error != nil {
		logger.Error(tAG_admin, "failed to prepare statement error:", error)
		return 0
	}
	//result, err := stmt.Query(pid, uid, lastid)
	logger.Info(tAG_admin, utils.RunFuncName(), selectCountByTable+tablename)
	result, err := stmt.Query()
	defer result.Close()
	if err != nil {
		return 0
	}
	var countResult sql.NullInt64
	if result.Next(){
		result.Scan(&countResult)
	}
	return countResult.Int64
}
