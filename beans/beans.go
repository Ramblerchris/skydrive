package beans

import (
	"github.com/skydrive/logger"
)

type MultiPartInfo struct {
	UploadId     string `json:"uploadid"`
	ChunkCount   int    `json:"chunkcount"`
	ChunkSize    int    `json:"chunksize"`
	Filesha1     string `json:"sha1,omitempty"`
	FileName     string `json:"filename,omitempty"`
	FileSize     int    `json:"size,omitempty"`
	Pid          int64  `json:"pid,omitempty"`
	SuccessIndex []int  `json:"successchunkindex,omitempty"`
}

type UToken struct {
	Tid        int64 `json:"id"`
	Uid        int64 `json:"uid"`
	Phone      string `json:"phone"`
	User_token string `json:"utoken"`
	Expiretime int64 `json:"expretime"`
}

type File struct {
	Id              int64  `json:"id,omitempty"`
	Filesha1        string `json:"sha1,omitempty"`
	FileName       	string `json:"filename,omitempty"`
	FileLocation   	string `json:"path,omitempty"`
	CreateAtTime     string `json:"createattimestr,omitempty"`
	CreateAtTimeLong int64  `json:"createattimelong,omitempty"`
	FileSize       	int64  `json:"size,omitempty"`
	Minitype       	string `json:"minitype,omitempty"`
	Ftype          	int32  `json:"ftype,omitempty"`
	Video_duration 	string `json:"video_duration,omitempty"`
}

type UserFile struct {
	File
	Filetype         int32  `json:"type,omitempty"`
	PId              int64  `json:"pid,omitempty"`
	CreateAtTime     string `json:"createattimestr,omitempty"`
	UpdateAtTime     string `json:"updatattimestr,omitempty"`
	CreateAtTimeLong int64  `json:"createattimelong,omitempty"`
	UpdateAtTimeLong int64  `json:"updatattimelong,omitempty"`
	FileHash_Pre     string `json:"sha1_pre,omitempty"`
}

type User struct {
	Id int32 `json:"id"`
	//`json:"-"` 字段不暴露给用户
	User_pwd        string `json:"-"`
	User_name       string `json:"user_name"`
	Email           string `json:"email"`
	Phone           string `json:"phone"`
	Photo_addr      string `json:"photo_addr"`
	Photo_addr_sha1 string `json:"photo_file_sha1"`
	Email_validated int32  `json:"email_validated"`
	Phone_validated int32  `json:"phone_validated"`
	Signup_at       string `json:"signup_at"`
	Last_active     string `json:"last_active"`
	//`json:"omitempty"`当字段为空时忽略此字段 不需要该字段返回时，让其赋值为空即可。
	Profile string `json:"profile,omitempty"`
	Status  int32  `json:"status"`
}

type SystenInfo struct {
	CpuPercent float64 `json:"CpuPercent"`
	//CpuTotal uint64 `json:"CpuTotal"`
	MemPercent float64 `json:"MemPercent"`
	MemTotal uint64 `json:"MemTotal"`
	DiskPercent float64 `json:"DiskPercent"`
	DiskTotal uint64 `json:"DiskTotal"`
	SwpPercent float64 `json:"SwpPercent"`
	SwpTotal uint64 `json:"SwpTotal"`
}

type HostInfo struct {
	Hostname string `json:"Hostname"`
	OS float32 `json:"OS"`
	IP float32 `json:"IP"`
}


func (filemeta *UserFile) String() {
	logger.Infof("filesha1:%s filename:%s  fileSize: %d  Location: %s  UpdateAtTime: %s ", filemeta.Filesha1, filemeta.FileName, filemeta.FileSize, filemeta.FileLocation, filemeta.UpdateAtTime)
}

