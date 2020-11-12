package utils

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"github.com/satori/go.uuid"
	"io"
	"os"
	"strings"
)

func GetFileSha1(file *os.File) string {
	//将文件位置移动到开始位置，否则得到的为空 的sha1 da39a3ee5e6b4b0d3255bfef95601890afd80709
	file.Seek(0, 0)
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum([]byte(nil)))
}

func GetStrSha1(str string) string {
	_sha1 := sha1.New()
	_sha1.Write([]byte(str))
	return hex.EncodeToString(_sha1.Sum([]byte(nil)))
}

func GetFileMD5(file *os.File) string {
	_md5 := md5.New()
	io.Copy(_md5, file)
	return hex.EncodeToString(_md5.Sum(nil))
}

func GetStrMD5(str string) string {
	_md5 := md5.New()
	_md5.Write([]byte(str))
	return hex.EncodeToString(_md5.Sum([]byte(nil)))
}

func BuildUUID() string {
	str:= uuid.NewV4().String()
	return strings.ReplaceAll(str,"-","")
}

