package utils

import (
	"fmt"
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
	fmt.Println(err)
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
	code := HashCode(sha)
	dir := fmt.Sprintf("%s/%d/%d/%d/%d/%d", rootpath, code&0xf, (code>>4)&0xf, (code>>8)&0xf, (code>>12)&0xf, (code>>16)&0xf)
	err := os.MkdirAll(dir, os.ModePerm)
	if err!=nil{
		return err,""
	}
	filename=fmt.Sprintf("%d_%s_%s", uid, BuildUUID(), filename)
	return nil, fmt.Sprintf("%s/%s", dir,  filename)
}

//文件是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

//创建缩略图
func CreateThumbDir(rootpath, sha, q,filename string) (error, string) {
	if sha == "" {
		sha = filename
	}
	code := HashCode(sha)
	//dir := fmt.Sprintf("%s/%d/%d/%d/%d/%d", rootpath, code&0xf, (code>>4)&0xf, (code>>8)&0xf, (code>>12)&0xf, (code>>16)&0xf)
	dir := fmt.Sprintf("%s/%d/%d", rootpath, code&0xf, (code>>4)&0xf)
	err := os.MkdirAll(dir, os.ModePerm)
	if err!=nil{
		return err,""
	}
	return nil, fmt.Sprintf("%s/%s_%s_%s", dir, q, sha,filename)
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



