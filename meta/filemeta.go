package meta

import (
	"time"
	"github.com/skydrive/db"
	"github.com/skydrive/handler"
)


var fileMetas map[string]handler.UserFile

func init() {
	fileMetas = make(map[string]handler.UserFile)
}
func AddOrUpdateFileMeta(filemeta handler.UserFile) {
	fileMetas[filemeta.Filesha1] = filemeta
}

func GetFileMeta(sha1 string) (*handler.UserFile, bool) {
	/*if filemeta, ok := fileMetas[sha1]; ok {
		return &filemeta, true
	}
	return nil, false*/
	if meta, err := db.GetFileInfoBySha1(sha1); err == nil {
		return &handler.UserFile{
			Id:             meta.Id.Int64,
			Filesha1:       meta.FileHash.String,
			FileName:       meta.FileName.String,
			FileSize:       meta.FileSize.Int64,
			Location:       meta.FileLocation.String,
			Minitype:       meta.Minitype.String,
			Ftype:          meta.Ftype.Int32,
			Video_duration: meta.Video_duration.String,
			UpdateAtTime:   "",
		}, true
	}
	return nil, false
}

func RemoveFileMeta(sha1 string) bool {
	delete(fileMetas, sha1)
	return db.UpdateFileInfoStatusBySha1(sha1, 0)
}


func GetNewFileMetaObject(value db.TableUserFile) *handler.UserFile {
	createlong, _ := time.Parse("2006-01-02 15:04:05", value.Create_at.String)
	updatelong, _ := time.Parse("2006-01-02 15:04:05", value.Update_at.String)
	return &handler.UserFile{
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

func GetNewUserMetaObject(info db.TableUser) *handler.User {
	return  &handler.User{
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
