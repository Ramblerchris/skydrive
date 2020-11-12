package main

import (
	"fmt"
	"github.com/skydrive/broadcast"
	"github.com/skydrive/cache/redisconn"
	"github.com/skydrive/config"
	"github.com/skydrive/handler"
	"github.com/skydrive/media"
	"github.com/skydrive/utils"
	"net/http"
	_ "net/http/pprof"
	"os"
	"time"
)

func handleFile(path string, fileinfo os.FileInfo, index int) {
	//fmt.Print(" 处理", fileinfo.Mode())
	//ext := filepath.Ext(fileinfo.Name())

	media.ScaleImage(path)

}

func main() {
	//media.ScaleImage("/Users/mac/Desktop/1589265866339_8927.JPG")
	//media.ScaleImage("/Users/mac/Desktop/1589191894238_8130.png")

	//handler.StartScanFile("/Users/mac/Desktop/image/", handleFile)

	go broadcast.StartUDPServerV2(config.UDP_SERVER_ListenPORT)

	//go broadcast.StartUDPGroup(config.UDP_SERVER_ListenPORT)
	go broadcast.StartUDPGroupV2(config.UDP_GroupSERVER_SendPORT,config.UDP_GroupSERVER_ListenPORT)

	fmt.Print(" AAAA")
	var bar utils.Bar
	bar.NewOption(0, 100)
	//bar.NewOptionWithGraph(0, 100,"*")
	for i := 0; i <= 100; i++ {
		time.Sleep(2 * time.Millisecond)
		bar.Play(int64(i))
		fmt.Print(" 处理", i)
	}
	bar.Finish()
	redisconn.GetRedisClient()
	httpServer()

}

func httpServer() {
	http.HandleFunc("/user/checknet", handler.CheckNet)
	http.HandleFunc("/user/register", handler.Register)
	http.HandleFunc("/user/signin", handler.Signin)
	http.HandleFunc("/user/getuserinfo", handler.TokenCheckInterceptor(handler.GetUserInfo))
	http.HandleFunc("/user/signout", handler.TokenCheckInterceptor(handler.SignOut))
	http.HandleFunc("/user/updatePhoto", handler.TokenCheckInterceptor(handler.UploadUserPhotoHandler))
	http.HandleFunc("/user/updateUserName", handler.TokenCheckInterceptor(handler.UploadUserNameHandler))
	//浏览器打开直接下载文件
	http.HandleFunc("/file/get_download", handler.TokenCheckInterceptor(handler.DownloadFile))
	//浏览器直接打开查看
	http.HandleFunc("/file/get_open", handler.TokenCheckInterceptor(handler.OpenFile1))
	http.HandleFunc("/userfile/getSha1IsExistList", handler.TokenCheckInterceptor(handler.GetSha1sIsExistByUser))
	http.HandleFunc("/userfile/getAllSha1sByUser", handler.TokenCheckInterceptor(handler.GetAllSha1sByUser))
	http.HandleFunc("/userfile/upload", handler.TokenCheckInterceptor(handler.UploadHandler))

	http.HandleFunc("/userfile/deletefiles", handler.TokenCheckInterceptor(handler.DeleteFilesBySha1sUser))
	http.HandleFunc("/userfile/deleteDir", handler.TokenCheckInterceptor(handler.DeleteFilesDirByUser))

	http.HandleFunc("/userfile/getlist", handler.TokenCheckInterceptor(handler.GetUserFileList))
	http.HandleFunc("/userfile/hitpass", handler.TokenCheckInterceptor(handler.HitPass))
	http.HandleFunc("/userfile/adddir", handler.TokenCheckInterceptor(handler.AddDir))
	http.HandleFunc("/userfile/dirlist", handler.TokenCheckInterceptor(handler.GetUserDirFileList))
	http.HandleFunc("/userfile/initmultipartinfo", handler.TokenCheckInterceptor(handler.InitMultipartUploadHandler))
	http.HandleFunc("/userfile/finishmultipartinfo", handler.TokenCheckInterceptor(handler.FinishMultipartUploadHandler))
	http.HandleFunc("/userfile/uploadmultipartinfo", handler.TokenCheckInterceptor(handler.UploadMultipartHandler))

	http.HandleFunc("/file/getinfo", handler.TokenCheckInterceptor(handler.GetFileSha1))
	http.HandleFunc("/file/update", handler.TokenCheckInterceptor(handler.UpdateFileInfo))
	http.HandleFunc("/file/delete", handler.TokenCheckInterceptor(handler.DeleteFile))

	fmt.Printf("开始启动本地服务 地址为 %s \r\n", config.ServeLocation)
	if error := http.ListenAndServe(config.ServeLocation, nil); error != nil {
		fmt.Printf("启动错误 error:%s \r\n", error.Error())
	}
}
