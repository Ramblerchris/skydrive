package utils

import (
	"fmt"
	"github.com/skydrive/logger"
	"hash/crc32"
	"io/ioutil"
	"math"
	"os"
	"strconv"
)

//合并
func FileMerge(chunkpath string, newfile string) bool {
	dirfiles, _ := ioutil.ReadDir(chunkpath)
	err := os.Remove(newfile)
	logger.Info(err)
	fill, error := os.OpenFile(newfile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if error != nil {
		println(error)
		return false
	}
	for index, filetemp := range dirfiles {
		//ioutil.ReadFile(filetemp.)
		fmt.Printf("index %d filename %s ", index, filetemp.Name())
		fmt.Printf("index %d filepath%s \n", index, filetemp.Name())
		file, error := os.OpenFile(chunkpath+strconv.Itoa(index)+".chunk", os.O_RDONLY, os.ModePerm)
		if error != nil {
			println(error)
		}
		all, _ := ioutil.ReadAll(file)
		fill.Write(all)
		file.Close()
	}
	return true
}

//创建文件存储路径
func CreateDirbySha1(rootpath, sha, filename string,uid int64) (error, string) {
	if sha == "" {
		sha = filename
	}
	path := getDirPath(rootpath, sha)
	err := os.MkdirAll(path, os.ModePerm)
	logger.Info("CreateDirbySha1 path:", path," err:",err)
	if err!=nil{
		return err,path
	}
	filename=fmt.Sprintf("%d_%s_%s", uid, BuildUUID(), filename)
	return nil, fmt.Sprintf("%s/%s", path,  filename)
}

// PathExists 文件是否存在
func PathExists(path string) (isExist bool, err error) {
	isExist, err, _ = PathExistsInfo(path)
	return
}

// PathExistsInfo 文件是否存在
func PathExistsInfo(path string) (bool, error, os.FileInfo) {
	info, err := os.Stat(path)
	if err == nil {
		return true, nil, info
	}
	if os.IsNotExist(err) {
		return false, nil, info
	}
	return false, err, info
}

func getDirPath(rootpath string, data string) string {
	code := HashCode(data)
	logger.Info("file sha1 ", data," code ",code)
	return fmt.Sprintf("%s/%d/%d/%d/%d/%d", rootpath, code&0xf, (code>>4)&0xf, (code>>8)&0xf, (code>>12)&0xf, (code>>16)&0xf)
}

//创建缩略图
func CreateThumbDir(rootpath, sha, q, filename string) (error, string) {
	if sha == "" {
		sha = filename
	}
	path := getDirPath(rootpath, sha)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err, ""
	}
	return nil, fmt.Sprintf("%s/%s_%s_%s", path, q, sha, filename)
}


func HashCode(str string) int {
	v := int(crc32.ChecksumIEEE([]byte(str)))
	return int(math.Abs(float64(v)))
}

func FileNameShort() string {
	return BuildUUID()
}

func GetFType(minetype string,isVideo bool) int {
	//0图片/1视频/2音乐/3文档/4压缩包)
	if isVideo{
		return 1
	}else {
		return 0
	}

}



