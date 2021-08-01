package syncfile

import (
	"fmt"
	"github.com/skydrive/logger"
	"io/ioutil"
	"log"
	"os"
	"strings"
)


type HandleFile func(string, os.FileInfo,int)


func StartScanFile(srcDirPath,srcRootPath,backupRootPath string,handle HandleFile) {
	err := os.MkdirAll(backupRootPath, os.ModePerm)
	logger.Info("CreatebackupRootPath path:", backupRootPath, " err:", err)
	if err != nil {
		return
	}
	countr, dircount, countsizer := scanDirBack(srcDirPath,srcRootPath,backupRootPath,handle, 0, 0, 0, 0)
	logger.Infof("文件数量：%d 文件夹%d 总大小：%d (%.5f GB)", countr, dircount, countsizer, float64(countsizer)/1e9)
}


//扫描文件夹
func scanDirBack(srcDirPath,srcRootPath,backupRootPath string,handle HandleFile, level int, filecount int64,dircount int64, countsize int64) (filecountr int64,dircountr int64, countsizer int64) {
	tag := "|-"
	for i := 0; i < level; i++ {
		tag = "	" + tag
	}
	dir, err := ioutil.ReadDir(srcDirPath)
	if err != nil {
		log.Fatal(err)
	}
	for index, info := range dir {
		tempfile := fmt.Sprintf("%s%s%s",srcDirPath ,string(os.PathSeparator) ,info.Name())
		fmt.Printf("%s %s L %d N %1d %s \n",info.ModTime().Format("2006-01-02 15:04:05"), tag, level, index, tempfile)
		if info.IsDir() {
			dircount++
			newpath:=strings.Replace(tempfile,srcRootPath,backupRootPath,1)
			err := os.MkdirAll(newpath, os.ModePerm)
			logger.Info("CreatebackupRootPath path:", newpath, " err:", err)
			if err != nil {
				return
			}
			filecount,dircount,countsize= scanDirBack(tempfile,srcRootPath,backupRootPath,handle, level+1, filecount, dircount,countsize)
		} else {
			filecount++
			countsize =countsize+ info.Size()
			handle(tempfile,info,index)
			/*newfile, error := os.Create(metaInfo.FileLocation)
			if error != nil {
				fmt.Printf("创建文件出错 %s \n", error.Error())
				response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "file create error")
				return
			}
			defer newfile.Close()
			metaInfo.FileSize, error = io.Copy(newfile, file)
			if error != nil {
				logger.Infof("保存文件出错 %s ", error.Error())
				response.ReturnResponseCodeMessage(w, config.Net_ErrorCode, "file copy error")
				return
			}*/
		}
	}
	return filecount,dircount, countsize
}



//扫描文件夹
func scanDir(srcDirPath,backupRootPath string,handle HandleFile, level int, filecount int64,dircount int64, countsize int64) (filecountr int64,dircountr int64, countsizer int64) {
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
		fmt.Printf("%s %s L %d N %1d %s \n",info.ModTime().Format("2006-01-02 15:04:05"), tag, level, index, tempfile)
		if info.IsDir() {
			dircount++
			filecount,dircount,countsize= scanDir(tempfile,backupRootPath,handle, level+1, filecount, dircount,countsize)
		} else {
			filecount++
			countsize =countsize+ info.Size()
			handle(tempfile,info,index)
		}
	}
	return filecount,dircount, countsize
}
