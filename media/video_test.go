package media

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"
	"strings"
	"testing"
)

func TestVideoThumbnail(t *testing.T) {

	var dirpath="/Users/wisn/Desktop/video/"
	var videoThumbNail="/Users/wisn/Desktop/video/result/"
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	for index, info := range dir {
		if info.IsDir() {
			continue
		}
		fmt.Printf("index %d name:%s\n",index,info.Name())

		var filepath=dirpath+info.Name()
		ext := strings.ToLower(path.Ext(filepath))
		if ext!=".mp4"&&ext!=".mkv"&&ext!=".rmvb"&&ext!=".mov"{
			continue
		}
		VideoThumbnail(filepath,fmt.Sprintf("%s%d.jpg",videoThumbNail,index))
	}

}

func TestVideoThumbnailGif(t *testing.T) {

	var dirpath="/Users/wisn/Desktop/video/"
	var videoThumbNail="/Users/wisn/Desktop/video/result/"
	dir, err := ioutil.ReadDir(dirpath)
	if err != nil {
		log.Fatal(err)
	}
	for index, info := range dir {
		if info.IsDir() {
			continue
		}
		fmt.Printf("index %d name:%s\n",index,info.Name())

		var filepath=dirpath+info.Name()
		ext := strings.ToLower(path.Ext(filepath))
		if ext!=".mp4"&&ext!=".mkv"&&ext!=".rmvb"&&ext!=".mov"{
			continue
		}
		VideoThumbnailGif(filepath,fmt.Sprintf("%s%d.gif",videoThumbNail,index),20)
	}

}
