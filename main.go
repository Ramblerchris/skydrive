package main

import (
	"fmt"
	"github.com/skydrive/broadcast"
	"github.com/skydrive/cache/redisconn"
	"github.com/skydrive/config"
	"github.com/skydrive/handler"
	"github.com/skydrive/logger"
	"github.com/skydrive/media"
	"github.com/skydrive/utils"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"time"
)
var (
	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
	Debug    ="true" // 是否为开发环境
)

func main() {
	config.Debug,_=strconv.ParseBool(Debug)
	logger.Setup()
	versionInfo()
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

	http.HandleFunc("/disk/upload", handler.TokenCheckInterceptor(handler.UploadDiskFileHandler))
	http.HandleFunc("/disk/delete", handler.TokenCheckInterceptor(handler.DeleteDiskFileDirByUidHandler))
	http.HandleFunc("/disk/hitpass", handler.TokenCheckInterceptor(handler.HitPassDiskBySha1Handler))
	http.HandleFunc("/disk/adddir", handler.TokenCheckInterceptor(handler.AddDiskFileDirByUidPidHandler))
	http.HandleFunc("/disk/dirlist", handler.TokenCheckInterceptor(handler.GetDiskDirFileListByPidHandler))

	http.HandleFunc("/admin/allUserList", handler.TokenCheckInterceptor(handler.GetAllUserListHandler))
	http.HandleFunc("/admin/allUserTokenList", handler.TokenCheckInterceptor(handler.GetAllUserTokenListHandler))
	http.HandleFunc("/admin/allFileList", handler.TokenCheckInterceptor(handler.GetAllFileListHandler))
	http.HandleFunc("/admin/allUserFileList", handler.TokenCheckInterceptor(handler.GetAllUserFileListHandler))
	http.HandleFunc("/admin/allDiskUserFileList", handler.TokenCheckInterceptor(handler.GetAllDiskUserFileListHandler))
	http.HandleFunc("/admin/systemInfo", handler.TokenCheckInterceptor(handler.GetSystemInfoHandler))

	logger.Infof("开始启动本地服务 地址为 %s ", config.ServeLocation)
	if error := http.ListenAndServe(config.ServeLocation, nil); error != nil {
		logger.Errorf("启动错误 error:%s ", error.Error())
	}
}

// Version 版本信息
func versionInfo() {
	if len(AppName) != 0 && AppName != "" {
		logger.Infof("App Name:\t %s", AppName)
		logger.Infof("App Version:\t %s", AppVersion)
		logger.Infof("Build version:\t %s", BuildVersion)
		logger.Infof("Build time:\t %s", BuildTime)
		logger.Infof("Git revision:\t %s", GitRevision)
		logger.Infof("Git branch:\t %s", GitBranch)
		logger.Infof("Golang Version: \t %s", GoVersion)
		logger.Infof("Debug :\t %s", Debug)
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
