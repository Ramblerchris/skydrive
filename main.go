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

func main() {
	broadcast.InitUDP()
	redisconn.GetRedisClient()
	httpServer()
}

func httpServer() {
	http.HandleFunc("/user/checknet", handler.CheckNetIsOkHandler)
	http.HandleFunc("/user/register", handler.RegisterHandler)
	http.HandleFunc("/user/signin", handler.SignInHandler)
	http.HandleFunc("/user/getuserinfo", handler.TokenCheckInterceptor(handler.GetUserInfoByTokenHandler))
	http.HandleFunc("/user/signout", handler.TokenCheckInterceptor(handler.SignOutHandler))
	http.HandleFunc("/user/updatePhoto", handler.TokenCheckInterceptor(handler.UpdataUploadUserPhotoHandler))
	http.HandleFunc("/user/updateUserName", handler.TokenCheckInterceptor(handler.UpdateUserNameByUidHandler))

	http.HandleFunc("/file/getinfo", handler.TokenCheckInterceptor(handler.GetFileInfoBySha1Handler))
	http.HandleFunc("/file/update", handler.TokenCheckInterceptor(handler.UpdateFileInfoFileNameBySha1Handler))
	http.HandleFunc("/file/delete", handler.TokenCheckInterceptor(handler.DeleteFileInfoBySha1Handler))
	//浏览器打开直接下载文件
	http.HandleFunc("/file/download", handler.TokenCheckInterceptor(handler.DownloadFileWebBySha1Handler))
	//浏览器直接打开查看
	http.HandleFunc("/file/open", handler.TokenCheckInterceptor(handler.OpenFile1Handler))

	http.HandleFunc("/userfile/getSha1IsExistList", handler.TokenCheckInterceptor(handler.GetSha1ListIsExistByUidHandler))
	http.HandleFunc("/userfile/getAllSha1sByUser", handler.TokenCheckInterceptor(handler.GetAllSha1ListByUidHandler))
	http.HandleFunc("/userfile/upload", handler.TokenCheckInterceptor(handler.UploadUserFileHandler))
	http.HandleFunc("/userfile/deletefiles", handler.TokenCheckInterceptor(handler.DeleteFileListBySha1sUidHandler))
	http.HandleFunc("/userfile/deleteDir", handler.TokenCheckInterceptor(handler.DeleteFileDirByUidHandler))
	http.HandleFunc("/userfile/getlist", handler.TokenCheckInterceptor(handler.GetUserFileListByUidHandler))
	http.HandleFunc("/userfile/hitpass", handler.TokenCheckInterceptor(handler.HitPassBySha1Handler))
	http.HandleFunc("/userfile/adddir", handler.TokenCheckInterceptor(handler.AddFileDirByUidPidHandler))
	http.HandleFunc("/userfile/dirlist", handler.TokenCheckInterceptor(handler.GetUserDirFileListByPidHandler))
	http.HandleFunc("/userfile/initmultipartinfo", handler.TokenCheckInterceptor(handler.InitMultipartUploadHandler))
	http.HandleFunc("/userfile/finishmultipartinfo", handler.TokenCheckInterceptor(handler.FinishMultipartUploadHandler))
	http.HandleFunc("/userfile/uploadmultipartinfo", handler.TokenCheckInterceptor(handler.UploadMultipartHandler))

	http.HandleFunc("/admin/allUserList", handler.TokenCheckInterceptor(handler.GetAllUserListHandler))
	http.HandleFunc("/admin/allUserTokenList", handler.TokenCheckInterceptor(handler.GetAllUserTokenListHandler))
	http.HandleFunc("/admin/allFileList", handler.TokenCheckInterceptor(handler.GetAllFileListHandler))
	http.HandleFunc("/admin/allUserFileList", handler.TokenCheckInterceptor(handler.GetAllUserFileListHandler))


	fmt.Printf("开始启动本地服务 地址为 %s \r\n", config.ServeLocation)
	if error := http.ListenAndServe(config.ServeLocation, nil); error != nil {
		fmt.Printf("启动错误 error:%s \r\n", error.Error())
	}
}




func handleFile(path string, fileinfo os.FileInfo, index int) {
	//fmt.Print(" 处理", fileinfo.Mode())
	//ext := filepath.Ext(fileinfo.Name())
	media.ScaleImage(path)
	//media.ScaleImage("/Users/mac/Desktop/1589265866339_8927.JPG")
	//media.ScaleImage("/Users/mac/Desktop/1589191894238_8130.png")
	//handler.StartScanFile("/Users/mac/Desktop/image/", handleFile)

}

func testProgressBar() {
	var bar utils.Bar
	bar.NewOption(0, 100)
	//bar.NewOptionWithGraph(0, 100,"*")
	for i := 0; i <= 100; i++ {
		time.Sleep(2 * time.Millisecond)
		bar.Play(int64(i))
		fmt.Print(" 处理", i)
	}
	bar.Finish()
}