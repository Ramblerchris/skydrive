package syncfile

import (
	"fmt"
	"github.com/skydrive/logger"
	"github.com/skydrive/utils"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type HandleFile func(string, os.FileInfo, int)

// StartScanFile 开启备份扫描
//@param  srcDirPath 源文件，要备份的文件目录
//@param  srcRootPath 源文件，要备份的文件目录根目录，在备份文件夹中不会重复创建
//@param  backupRootPath
func StartScanFile(srcDirPath, srcRootPath, backupRootPath string, handle HandleFile) {
	err := os.MkdirAll(backupRootPath, os.ModePerm)
	logger.Info("CreatebackupRootPath path:", backupRootPath)
	if err != nil {
		return
	}
	countr, dircount, countsizer := scanDirBack(srcDirPath, srcRootPath, backupRootPath, handle, 0, 0, 0, 0)
	logger.Infof("文件数量：%d 文件夹%d 总大小：%d (%.5f GB)", countr, dircount, countsizer, float64(countsizer)/1e9)
}

//扫描文件夹
func scanDirBack(srcDirPath, srcRootPath, backupRootPath string, handle HandleFile, level int, filecount int64, dircount int64, countsize int64) (filecountr int64, dircountr int64, countsizer int64) {

	dir, err := ioutil.ReadDir(srcDirPath)
	if err != nil {
		logger.Error(err)
	}
	for index, info := range dir {
		tempfile := fmt.Sprintf("%s%s%s", srcDirPath, string(os.PathSeparator), info.Name())
		logger.Infof("%s L %d N %1d %s \n", info.ModTime().Format("2006-01-02 15:04:05"),  level, index, tempfile)
		newpath := strings.Replace(tempfile, srcRootPath, backupRootPath, 1)
		if info.IsDir() {
			dircount++
			existdir, _ ,newinfo:= utils.PathExistsInfo(newpath)
			if existdir {
				logger.Info("exist dir:",info.ModTime().Format("2006-01-02 15:04:05"), newinfo.ModTime().Format("2006-01-02 15:04:05"))
				//如果存在
				/*if  newinfo.ModTime().Before(info.ModTime()){
					//修改时间在原始文件时间之前，需要扫描
					fmt.Println("Before path:3333")
					writeModTime(newpath,info.ModTime())
					filecount, dircount, countsize = scanDirBack(tempfile, srcRootPath, backupRootPath, handle, level+1, filecount, dircount, countsize)
				}else {

				}*/
				//读取时间记录比较 同时同步修改时间
				modTime, err := readModTime(newpath)
				if err !=nil|| modTime.UnixNano()!=info.ModTime().UnixNano(){
					//修改时间不一致，继续扫描
					logger.Info("修改时间不一致，继续扫描")
					writeModTime(newpath,info.ModTime())
					filecount, dircount, countsize = scanDirBack(tempfile, srcRootPath, backupRootPath, handle, level+1, filecount, dircount, countsize)

				} else {
					logger.Info("修改时间一致，跳过扫描")
				}
			} else {
				//如果不存在，先创建文件夹，再继续扫描
				logger.Info("CreatebackupRootPath path:", newpath)
				err := os.MkdirAll(newpath, os.ModePerm)
				if err != nil {
					return
				}
				writeModTime(newpath,info.ModTime())
				filecount, dircount, countsize = scanDirBack(tempfile, srcRootPath, backupRootPath, handle, level+1, filecount, dircount, countsize)
			}
		} else {
			countsize = countsize + info.Size()
			handle(tempfile, info, index)
			exists, _ := utils.PathExists(newpath)
			if exists {
				//如果存在，比较文件名，文件大小，文件Alder32 /crc32, 如果不相同，直接复制，如果相同再次验证sha1
			} else {
				filecount++
				//不存在直接复制
				newfile, error := os.Create(newpath)
				if error != nil {
					logger.Error("创建文件出错 %s \n", error.Error())
					return
				}
				efile, err := os.Open(tempfile)
				if err != nil {
					logger.Error("could not open file for : %s", tempfile)
					return
				}
				_, error = io.Copy(newfile, efile)
				logger.Infof("复制文件  : %s",newpath)
			}
		}
	}
	return filecount, dircount, countsize
}


func writeModTime(DirPath string , modtime time.Time)  {
	var f *os.File
	var path = DirPath + string(os.PathSeparator) + "skytime.syncv1"
	exists, _ := utils.PathExists(path)
	if !exists {
		f,_=os.Create(path)
	}else{
		f, _= os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	}
	defer f.Close()
	str:=fmt.Sprintf("%d ",modtime.UnixNano())
	fmt.Println(str)
	n, error := f.WriteString(str)
	if error != nil {
		fmt.Printf("write %d %s \n", n, error)
	}
	f.Sync()
}

func readModTime(DirPath string ) (modtime time.Time ,err error) {
	f, _:= os.OpenFile(DirPath+string(os.PathSeparator)+"skytime.syncv1", os.O_RDONLY, 0600)
	contentByte,err:=ioutil.ReadAll(f)
	if err !=nil{
		return modtime, nil
	}
	content:=strings.TrimSpace(string(contentByte))
	timena, err := strconv.ParseInt(content, 10, 64)
	if err !=nil{
		fmt.Println("abc:",err.Error())
	}
	modtime=time.Unix(0,timena)
	logger.Infof("read  %s  %s %s  %d %s \n",contentByte,DirPath,content,timena,modtime.Format("2006-01-02 15:04:05"))
	return modtime,nil
}


func ChangeDirModTime(parentDir string ,dirlevel int )bool  {
	start:=time.Now().UnixNano()
	success:=true
	for i := 0; i < dirlevel; i++{
		parentDir=filepath.Dir(parentDir)
		now := time.Now()
		err := os.Chtimes(parentDir, now, now)
		if err!=nil{
			success=false
			break
		}
	}
	t:=time.Now().UnixNano()-start
	logger.Infof("ChangeDirModTime:%d \n",t)
	return success
}

//扫描文件夹
func scanDir(srcDirPath, backupRootPath string, handle HandleFile, level int, filecount int64, dircount int64, countsize int64) (filecountr int64, dircountr int64, countsizer int64) {
	tag := "|-"
	for i := 0; i < level; i++ {
		tag = "	" + tag
	}
	dir, err := ioutil.ReadDir(srcDirPath)
	if err != nil {
		log.Fatal(err)
	}
	for index, info := range dir {
		tempfile := srcDirPath + "/" + info.Name()
		logger.Infof("%s %s L %d N %1d %s \n", info.ModTime().Format("2006-01-02 15:04:05"), tag, level, index, tempfile)
		if info.IsDir() {
			dircount++
			filecount, dircount, countsize = scanDir(tempfile, backupRootPath, handle, level+1, filecount, dircount, countsize)
		} else {
			filecount++
			countsize = countsize + info.Size()
			handle(tempfile, info, index)
		}
	}
	return filecount, dircount, countsize
}
