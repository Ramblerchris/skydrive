package db

import (
	"database/sql"
	"fmt"
	mysql "github.com/skydrive/db/mysqlconn"
	"github.com/skydrive/logger"
	"github.com/skydrive/utils"
	"strings"
)

const (
	selectUDiskFileBySha1      = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at,filetype,minitype,ftype,video_duration from tbl_cloud_disk where file_sha1=? and uid=? and status=1 limit 1"
	selectUDiskFileByid        = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at,filetype,minitype,ftype,video_duration from tbl_cloud_disk where id=?  and status=1 limit 1"
	selectUDiskFileByUid       = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_cloud_disk where uid=? and status=1 "
	selectUDiskFileByUidPage   = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_cloud_disk where uid=? and status=1 and id>? limit ? "
	selectUDiskFileByUidAndPid = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_cloud_disk where uid=? and pid=? and status=1 "
	//selectUDiskFileByUidAndPidPage        = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_cloud_disk where uid=? and pid=? and status=1 and id>? limit ? "
	selectUDiskFileByUidAndPidPage        = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_cloud_disk where uid=? and pid=? and status=1 and id<? order by id desc  limit ? "
	selectUDiskFileByUidAndPidAndSha1     = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_cloud_disk where file_sha1=? and uid=? and pid=? and status=1 limit 1 "
	selectUDiskFileByUidAndPidAndFileName = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at ,filetype,minitype,ftype,video_duration from tbl_cloud_disk where uid=? and pid=? and file_name=? and status=1 "
	saveUDiskFileinfo                     = "insert into tbl_cloud_disk(uid,pid,phone,file_sha1,file_name,file_size,file_addr,status,minitype,ftype,video_duration) values(?,?,?,?,?,?,?,?,?,?,?)"
	saveUDiskFileDirinfo                  = "insert into tbl_cloud_disk(uid,pid,phone,file_name,status,filetype) values(?,?,?,?,?,?)"
	updateUDiskFileInfo                   = "update tbl_cloud_disk set status=? where file_sha1=?"
	updateUDiskFileStatus                 = "update tbl_cloud_disk set status=? where uid=? and  pid=? and file_sha1 in('%s')"
	updateUDiskFileDirStatus              = "update tbl_cloud_disk set status=? where id=? "
	updateUDiskFileDirfile_sha1_pre       = "update tbl_cloud_disk set file_sha1_pre=? where id=? "
	updateUDiskFileDirsStatus             = "update tbl_cloud_disk set status=? where id in('%s')"
	selectUDiskFileBysha1s                = "select id,pid, uid,phone,file_sha1,file_sha1_pre,file_name,file_size,file_addr,create_at,update_at,filetype,minitype,ftype,video_duration from tbl_cloud_disk where uid=? and file_sha1 in('%s') "
	selectUDiskFileAllsha1s               = "select file_sha1 from tbl_cloud_disk where uid=?"
	//selectUDiskFileCountByUid 			  = "select count(*) from tbl_cloud_disk where  uid=? and status=1 and id>? "
	selectUDiskFileCountByUid = "select count(*) from tbl_cloud_disk where  uid=? and status=1  "
	//selectUDiskFileCountByUidPid 		  = "select count(*) from tbl_cloud_disk where  pid=? and uid=? and status=1 and id>? "
	selectUDiskFileCountByUidPid = "select count(*) from tbl_cloud_disk where  pid=? and uid=? and status=1  "
	selectMaxIdFromUserDiskFile  = "select max(id) from tbl_cloud_disk where  pid=? and uid=? and status=1  order by id desc"
	tAG_userDiskfile             = "diskfile.go：sql"
)

//创建一个文件夹
func SaveDiskDirInfo(uid, pid int64, phone, dirName string) (bool, int64) {
	//stmt ,error:=mysql.DbConnect().Prepare(saveFile)
	stmt, error := mysql.DbConnect().Prepare(saveUDiskFileDirinfo)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error.Error())
		return false, -1
	}
	defer stmt.Close()
	if pid == 0 {
		pid = -1
	}
	//filetype 文件夹1 文件-1
	//logger.Info(tAG_userDiskfile, utils.RunFuncName(),"start")
	exec, error := stmt.Exec(uid, pid, phone, dirName, 1, 1)
	logger.Info(tAG_userDiskfile, utils.RunFuncName(),saveUDiskFileDirinfo,uid, pid, phone, dirName, 1, 1)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to Exec error:", error)
		return false, -1
	}
	lastinsertId, _ := exec.LastInsertId()

	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，但未添加，
			logger.Error(tAG_userDiskfile, "failed with hash has been upload", error)
			return false, -1
		}
	}
	return true, lastinsertId
}

//保存文件信息到用户对应文件夹
func SaveDiskFileInfo(uid, pid int64, phone, filehash, filename, location string, filesize int64, minitype string, ftype int, video_duration string) bool {
	stmt, error := mysql.DbConnect().Prepare(saveUDiskFileinfo)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error.Error())
		return false
	}
	defer stmt.Close()
	if pid == 0 {
		pid = -1
	}
	logger.Info(tAG_userDiskfile, utils.RunFuncName(),"start")

	exec, error := stmt.Exec(uid, pid, phone, filehash, filename, filesize, location, 1, minitype, ftype, video_duration)
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), saveUDiskFileinfo,uid, pid, phone, filehash, filename, filesize, location, 1, minitype, ftype, video_duration)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，但未添加，
			logger.Errorf(tAG_userDiskfile+"failed with hash:%s has been upload", error)
			return false
		}
	}
	return true
}

//修改所有用户的这个文件状态
func UpdateDiskUserInfoStatusBySha1(filehash string, filestatus int8) bool {
	stmt, error := mysql.DbConnect().Prepare(updateUDiskFileInfo)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error.Error())
		return false
	}
	logger.Info(tAG_userDiskfile, utils.RunFuncName(),"start")
	defer stmt.Close()
	exec, error := stmt.Exec(filestatus, filehash)
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), updateUDiskFileInfo)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，修改未成功，
			logger.Error(tAG_userDiskfile+"failed with hash:%s has been upload", error)
			return false
		}
	}
	return true
}

//查看单个文件具体信息
func GetDiskFileInfoById(id int64) (*TableUserFile, error) {
	stmt, error := mysql.DbConnect().Prepare(selectUDiskFileByid)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), selectUDiskFileByid)
	tfile := TableUserFile{}
	row := stmt.QueryRow(id)
	error = row.Scan(
		&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to QueryRow error:", error)
		return nil, error
	}
	return &tfile, nil
}


//查看当前用户是否已经存储filehash对应的文件
func GetDiskFileMetaByPidUidSha1(filehash string, uid, pid int64) (*TableUserFile, error) {
	stmt, error := mysql.DbConnect().Prepare(selectUDiskFileByUidAndPidAndSha1)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	tfile := TableUserFile{}
	row := stmt.QueryRow(filehash, uid, pid)
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), selectUDiskFileByUidAndPidAndSha1,filehash, uid, pid)
	error = row.Scan(
		&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to QueryRow error:", error)
		return nil, error
	}
	return &tfile, nil
}

//查看当前用户是否已经存储filehash对应的文件
func GetDiskFileInfoByUidSha1(filehash string, uid int64) (*TableUserFile, error) {
	stmt, error := mysql.DbConnect().Prepare(selectUDiskFileBySha1)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	logger.Info(tAG_userDiskfile, utils.RunFuncName(),selectUDiskFileBySha1)
	tfile := TableUserFile{}
	row := stmt.QueryRow(filehash, uid)
	error = row.Scan(
		&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to QueryRow error:", error)
		return nil, error
	}
	return &tfile, nil
}

//查询用户同级文件夹下相同文件名的文件夹列表
func GetDiskFileInfoListByUidPidDirName(uid, pid int64, filename string) (tableUserFile []TableUserFile, err error) {
	stmt, error := mysql.DbConnect().Prepare(selectUDiskFileByUidAndPidAndFileName)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return nil, error
	}
	defer stmt.Close()
	rowdata, error := stmt.Query(uid, pid, filename)
	defer  rowdata.Close()
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), selectUDiskFileByUidAndPidAndFileName, uid, pid, filename)
	if error != nil {
		logger.Error("failed to Exec error:", error)
		return tableUserFile, error
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error = rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
		if error != nil {
			logger.Error(tAG_userDiskfile, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}
	return tableUserFile, nil
}

//查看当前用户 pid 对应子目录所有文件列表，包括文件列表
func GetDiskFileListByUidPid(uid, pid int64, pageNo, pageSize, lastid int64) (tableUserFile []TableUserFile, err error, total int64) {
	//目前忽略了文件类型
	stmt, error := mysql.DbConnect().Prepare(selectUDiskFileByUidAndPidPage)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return tableUserFile, error, 0
	}
	defer stmt.Close()
	if lastid == -1 {
		lastid = GetUserDiskListMaxCountByUid(uid, pid)+1
	}
	rowdata, error := stmt.Query(uid, pid, lastid, pageSize)
	defer  rowdata.Close()
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), selectUDiskFileByUidAndPidPage, uid, pid, lastid, pageSize)
	if error != nil {
		logger.Error("failed to Exec error:", error)
		return tableUserFile, error, 0
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error = rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
		if error != nil {
			logger.Error(tAG_userDiskfile, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}

	return tableUserFile, nil, GetUserDiskFileListCountByUidPid(uid, pid, lastid)
}

//查询当前用户指定文件下的文件数
func GetUserDiskFileListCountByUidPid(uid, pid int64, lastid int64) (count int64) {
	//目前忽略了文件类型
	stmt, error := mysql.DbConnect().Prepare(selectUDiskFileCountByUidPid)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return 0
	}
	defer stmt.Close()
	//result, err := stmt.Query(pid, uid, lastid)
	result, err := stmt.Query(pid, uid)
	defer result.Close()
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), selectUDiskFileCountByUidPid, pid, uid)
	if err != nil {
		return 0
	}
	var countResult sql.NullInt64
	if result.Next() {
		result.Scan(&countResult)
	}
	return countResult.Int64
}

//查询当前用户指定文件下最大的id
func GetUserDiskListMaxCountByUid(uid, pid int64) (maxid int64) {
	//目前忽略了文件类型
	stmt, error := mysql.DbConnect().Prepare(selectMaxIdFromUserDiskFile)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return 0
	}
	defer stmt.Close()
	//result, err := stmt.Query(uid, lastid)
	result, err := stmt.Query(pid, uid)
	defer  result.Close()
	logger.Info(tAG_userDiskfile, utils.RunFuncName(),selectMaxIdFromUserDiskFile, pid, uid)
	if err != nil {
		return 0
	}
	var countResult sql.NullInt64
	if result.Next() {
		result.Scan(&countResult)
	}
	return countResult.Int64
}


//查询当前用户指定文件下的文件数
func GetUserDiskDirFileListCountByUid(uid, lastid int64) (count int64) {
	//目前忽略了文件类型
	stmt, error := mysql.DbConnect().Prepare(selectUDiskFileCountByUid)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return 0
	}
	defer stmt.Close()
	//result, err := stmt.Query(uid, lastid)
	result, err := stmt.Query(uid)
	defer result.Close()
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), selectUDiskFileCountByUid, uid)
	if err != nil {
		return 0
	}
	var countResult sql.NullInt64
	if result.Next() {
		result.Scan(&countResult)
	}
	return countResult.Int64
}

//查询用户所有的文件
func GetUserDiskFileListMetaByUid(uid int64, pageNo, pageSize, lastid int64) (tableUserFile []TableUserFile, err error, total int64) {
	stmt, error := mysql.DbConnect().Prepare(selectUDiskFileByUidPage)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return nil, error, 0
	}
	defer stmt.Close()
	rowdata, error := stmt.Query(uid, lastid, pageSize)
	defer rowdata.Close()
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), selectUDiskFileByUidPage, uid, lastid, pageSize)
	if error != nil {
		logger.Error("failed to Exec error:", error)
		return tableUserFile, error, 0
	}
	tableUserFile = make([]TableUserFile, 0)
	for rowdata.Next() {
		tfile := TableUserFile{}
		error = rowdata.Scan(
			&tfile.Id, &tfile.PId, &tfile.Uid, &tfile.Phone, &tfile.FileHash, &tfile.FileHash_Pre, &tfile.FileName, &tfile.FileSize, &tfile.FileLocation, &tfile.Create_at, &tfile.Update_at, &tfile.Filetype, &tfile.Minitype, &tfile.Ftype, &tfile.Video_duration)
		if error != nil {
			logger.Error(tAG_userDiskfile, "failed to Query error:", error)
			continue
		}
		tableUserFile = append(tableUserFile, tfile)
	}
	return tableUserFile, nil, GetUserDiskDirFileListCountByUid(uid, lastid)
}

//批量修改当前用户文件状态
func UpdateUserDiskFileStatusBySha1sUidPid(uid, pid int64, filestatus int8, sha1s []string) bool {
	sha1sJoin := strings.Join(sha1s, "','")
	sprintf := fmt.Sprintf(updateUDiskFileStatus, sha1sJoin)
	stmt, error := mysql.DbConnect().Prepare(sprintf)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(filestatus, uid, pid)
	logger.Info(tAG_userDiskfile, utils.RunFuncName(),sprintf, filestatus, uid, pid)

	if error != nil {
		logger.Error(tAG_userDiskfile, "failed Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows <= 0 {
			//执行成功，修改未成功，
			logger.Error(tAG_userDiskfile+ "failed with hash:%s has been upload", error)
			return false
		}
	}
	return true
}


//当前用户文件状态批量修改
func UpdateUserDiskFileDirStatusByIds(ids []string, filestatus int8) (issuccess bool,rows int64 ) {
	sha1sJoin := strings.Join(ids, "','")
	sprintf := fmt.Sprintf(updateUDiskFileDirsStatus, sha1sJoin)
	stmt, error := mysql.DbConnect().Prepare(sprintf)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		issuccess = false
		rows = 0
	}
	defer stmt.Close()
	exec, error := stmt.Exec(filestatus)
	logger.Info(tAG_userDiskfile, utils.RunFuncName(), updateUDiskFileDirsStatus, filestatus)

	if error != nil {
		logger.Error(tAG_userDiskfile, "failed Exec error:", error)
		issuccess = false
		rows = 0
	}
	if rows, error = exec.RowsAffected(); error != nil {
		if rows <= 0 {
			//执行成功，修改未成功，
			logger.Error(tAG_userDiskfile+"failed with hash:%s has been upload\n", error)
			issuccess = false
			rows = 0
		}
	} else {
		issuccess = true
	}
	return issuccess, rows
}

//当前用户文件预览图
func UpdateUserDiskFileDirPreSha1ById(sha1 string, id int64) bool {
	stmt, error := mysql.DbConnect().Prepare(updateUDiskFileDirfile_sha1_pre)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed to prepare statement error:", error)
		return false
	}
	defer stmt.Close()
	exec, error := stmt.Exec(sha1, id)
	logger.Info(tAG_userDiskfile,  utils.RunFuncName(), updateUDiskFileDirfile_sha1_pre, sha1, id)
	if error != nil {
		logger.Error(tAG_userDiskfile, "failed Exec error:", error)
		return false
	}
	if rows, error := exec.RowsAffected(); error == nil {
		if rows == 1 {
			return true
		}
	}
	return false
}
