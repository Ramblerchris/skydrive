package syncfile

import (
	"os"
	"testing"
)

func Test_Scanfile(t *testing.T) {
	//StartScanFile("/Users/wisn/Desktop/video","/Users/wisn/Desktop/video","/Users/wisn/Desktop/video2", func(s string, info os.FileInfo, i int) {
	StartScanFile("/Users/wisn/GoWork/skydrive/upalbum","/Users/wisn/GoWork/skydrive/upalbum","/Users/wisn/Desktop/upalbum3", func(s string, info os.FileInfo, i int) {

	})
	// 获取路径的 上级目录路径
	//println(filepath.Dir(`/Users/wisn/Desktop/video/ad.txt`))
	//println(filepath.Dir(`/Users/wisn/Desktop/video`))
	//println(filepath.Dir(`/Users/wisn/Desktop`))
	//changeDirModTime("/Users/wisn/Desktop/scan/baidu/aip-ocr-android-sdk-1.4.6/libs/arm64-v8a/libocr-sdk.so",3)
}

