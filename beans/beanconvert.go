package beans

import (
	"github.com/skydrive/db"
	"time"
)

func GetUserFileObject(value db.TableUserFile) *UserFile {
	createlong, _ := time.Parse("2006-01-02 15:04:05", value.Create_at.String)
	updatelong, _ := time.Parse("2006-01-02 15:04:05", value.Update_at.String)
	return &UserFile{
		File : File{
			Id:             value.Id.Int64,
			Filesha1:       value.FileHash.String,
			FileName:       value.FileName.String,
			FileSize:       value.FileSize.Int64,
			FileLocation:   value.FileLocation.String,
			Minitype:       value.Minitype.String,
			Ftype:          value.Ftype.Int32,
			Video_duration: value.Video_duration.String,
		},
		PId:          value.PId.Int64,
		FileHash_Pre: value.FileHash_Pre.String,

		Filetype:         value.Filetype.Int32,
		CreateAtTime:     value.Create_at.String,
		UpdateAtTime:     value.Update_at.String,
		CreateAtTimeLong: createlong.UnixNano(),
		UpdateAtTimeLong: updatelong.UnixNano(),
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
