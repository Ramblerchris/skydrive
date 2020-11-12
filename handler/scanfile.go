package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)
type HandleFile func(string, os.FileInfo,int)


func StartScanFile(dirpath string,handle HandleFile) {
	countr, dircount, countsizer := scanDir(dirpath,handle, 0, 0, 0, 0)
	fmt.Printf("文件数量：%d 文件夹%d 总大小：%d (%.5f GB)\n", countr, dircount, countsizer, float64(countsizer)/1e9)
}
//扫描文件夹
func scanDir(dirpath string,handle HandleFile, level int, filecount int64,dircount int64, countsize int64) (filecountr int64,dircountr int64, countsizer int64) {
	tag := "|-"
	for i := 0; i < level; i++ {
		tag = "	" + tag
	}
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	for index, info := range dir {
		tempfile := dirpath + "/" + info.Name()
		fmt.Printf("%s L %d N %1d %s \n", tag, level, index, tempfile)
		if info.IsDir() {
			dircount++
			filecount,dircount,countsize= scanDir(tempfile,handle, level+1, filecount, dircount,countsize)
		} else {
			filecount++
			countsize =countsize+ info.Size()
			handle(tempfile,info,index)
		}
	}
	return filecount,dircount, countsize
}
