package meta

import (
	"fmt"
	"github.com/skydrive/db"
	"time"
)

type FileMeta struct {
	Id               int64  `json:"id,omitempty"`
	PId              int64  `json:"pid,omitempty"`
	Filesha1         string `json:"sha1,omitempty"`
	FileHash_Pre     string `json:"sha1_pre,omitempty"`
	FileName         string `json:"filename,omitempty"`
	FileSize         int64  `json:"size,omitempty"`
	Location         string `json:"path,omitempty"`
	Filetype         int32  `json:"type,omitempty"`
	CreateAtTime     string `json:"createattimestr,omitempty"`
	UpdateAtTime     string `json:"updatattimestr,omitempty"`
	CreateAtTimeLong int64  `json:"createattimelong,omitempty"`
	UpdateAtTimeLong int64  `json:"updatattimelong,omitempty"`
	Minitype         string `json:"minitype,omitempty"`
	Ftype            int32  `json:"ftype,omitempty"`
	Video_duration   string `json:"video_duration,omitempty"`
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}
func AddOrUpdateFileMeta(filemeta FileMeta) {
	fileMetas[filemeta.Filesha1] = filemeta
}

func GetFileMeta(sha1 string) (*FileMeta, bool) {
	/*if filemeta, ok := fileMetas[sha1]; ok {
		return &filemeta, true
	}
	return nil, false*/
	if meta, err := db.GetFileInfoBySha1(sha1); err == nil {
		return &FileMeta{
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

func (filemeta *FileMeta) String() {
	fmt.Printf("filesha1:%s filename:%s  fileSize: %d  Location: %s  UpdateAtTime: %s ", filemeta.Filesha1, filemeta.FileName, filemeta.FileSize, filemeta.Location, filemeta.UpdateAtTime)
}

func GetNewFileMetaObject(value db.TableUserFile) *FileMeta {
	createlong, _ := time.Parse("2006-01-02 15:04:05", value.Create_at.String)
	updatelong, _ := time.Parse("2006-01-02 15:04:05", value.Update_at.String)
	return &FileMeta{
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
