package response

import (
	"fmt"
	"github.com/skydrive/db"
	"github.com/skydrive/handler/cache"
	"sync"
	"time"
)

const TAG = "filemeta.go"
var filecache=cache.NewFileCache()

func init() {}

type SafeMap struct {
	sync.RWMutex
	Map map[string]UserFile
}

func (sm *SafeMap) ReadMap(key string)(UserFile ,bool){
	sm.RLock()
	defer sm.RUnlock()
	a,b:= sm.Map[key]
	return a,b
}
func (sm *SafeMap) WriteMap(key string, value UserFile) {
	sm.RLock()
	defer sm.RUnlock()
	sm.Map[key] = value
}

func (sm *SafeMap) DeleteMap(key string) {
	sm.RLock()
	defer sm.RUnlock()
	delete(sm.Map,key)
}

func AddOrUpdateFileMeta(filemeta UserFile) {
	filecache.WriteFileCache(filemeta.Filesha1,filemeta)
}

func GetFileMeta(sha1 string) (*UserFile, bool) {
	/*if userinfo ,isExist:=fileMetas[sha1];isExist{
		fmt.Println(TAG,"缓存获取:", userinfo)
		return &userinfo,true
	}else{
		fmt.Println(TAG,"缓存获取失败", userinfo,len(fileMetas),fileMetas)
	}*/
	if userinfo ,isExist:=filecache.ReadFileCache(sha1);isExist{
		fmt.Println(TAG,"缓存获取:", userinfo)
		return &userinfo,true
	}else{
		fmt.Println(TAG,"缓存获取失败", userinfo)
	}
	if meta, err := db.GetFileInfoBySha1(sha1); err == nil {
		userinfo:= UserFile{
			Id:             meta.Id.Int64,
			Filesha1:       meta.Filesha1.String,
			FileName:       meta.FileName.String,
			FileSize:       meta.FileSize.Int64,
			Location:       meta.FileLocation.String,
			Minitype:       meta.Minitype.String,
			Ftype:          meta.Ftype.Int32,
			Video_duration: meta.Video_duration.String,
			UpdateAtTime:   "",
		}
		filecache.WriteFileCache(sha1,userinfo)
		return &userinfo,true
	}
	return nil, false
}

func RemoveFileMeta(sha1 string) bool {
	//delete(fileMetas, sha1)
	filecache.DeleteFileCache(sha1)
	return db.UpdateFileInfoStatusBySha1(sha1, 0)
}

func GetUserFileObject(value db.TableUserFile) *UserFile {
	createlong, _ := time.Parse("2006-01-02 15:04:05", value.Create_at.String)
	updatelong, _ := time.Parse("2006-01-02 15:04:05", value.Update_at.String)
	return &UserFile{
		Id:               value.Id.Int64,
		PId:              value.PId.Int64,
		Filesha1:         value.FileHash.String,
		FileHash_Pre:     value.FileHash_Pre.String,
		FileName:         value.FileName.String,
		FileSize:         value.FileSize.Int64,
		Location:         value.FileLocation.String,
		Filetype:         value.Filetype.Int32,
		CreateAtTime:     value.Create_at.String,
		UpdateAtTime:     value.Update_at.String,
		CreateAtTimeLong: createlong.UnixNano(),
		UpdateAtTimeLong: updatelong.UnixNano(),
		Minitype:         value.Minitype.String,
		Ftype:            value.Ftype.Int32,
		Video_duration:   value.Video_duration.String,
	}
}

func GetUserObject(info db.TableUser) *User {
	return &User{
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
	}
}

func GetUserTokenObject(info db.TableUToken) *UToken {
	return &UToken{
		Tid:        info.Tid.Int64,
		Uid:        info.Uid.Int64,
		Phone:      info.Phone.String,
		User_token: info.User_token.String,
		Expiretime: info.Expiretime.Int64,
	}
}

func GetFileObject(info db.TableFile) *File {
	return &File{
		Id:             info.Id.Int64,
		Filesha1:       info.Filesha1.String,
		FileName:       info.FileName.String,
		FileLocation:   info.FileLocation.String,
		FileSize:       info.FileSize.Int64,
		Minitype:       info.Minitype.String,
		Ftype:          info.Ftype.Int32,
		Video_duration: info.Video_duration.String,
	}
}

