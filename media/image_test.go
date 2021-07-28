package media

import (
	"bytes"
	"fmt"
	"github.com/rwcarlsen/goexif/exif"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func Test_image(t *testing.T) {
	//ScaleImage("/Users/mac/Desktop/5_03bf5931ffcd43068197154706049452_IMG_20210218_181450.jpg")
	//var path = "/Users/mac/Desktop/5_03bf5931ffcd43068197154706049452_IMG_20210218_181450.jpg"
	//var output_path = "/Users/mac/Desktop/5_03bf5931ffcd43068197154706049452_IMG_20210218_181450_1.jpg"

/*	var path1 = "/Users/wisn/Desktop/video/1627138076033.jpeg"
	var path3 = "/Users/wisn/Desktop/video/1627138076033_1.jpeg"
	var path2 = "/Users/wisn/Desktop/video/1627138076033_2.jpeg"
	var testbmp = "/Users/wisn/Desktop/video/testbmp.bmp"
	getOrigin(path1)
	getOrigin(path3)
	getOrigin(path2)
	getOrigin(testbmp)*/

	var dirPath="/Users/wisn/Desktop/video/testtype"
	dir, err := ioutil.ReadDir(dirPath)
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range dir {
		tempfile := dirPath + "/" + info.Name()
		getOrigin(tempfile)
	}
	/*	var output_path = "/Users/wisn/Desktop/video/1627138076033_1.jpeg"
		//var path="/Users/mac/Desktop/10a3903d83294b3ab15fd05f1644bef0.jpg"
		//var output_path="/Users/mac/Desktop/10a3903d83294b3ab15fd05f1644bef0_1.jpg"
		ScaleImageByWidthAndQuity(path, 0, 0.5, 100, output_path)*/
}

func getOrigin(path string) {
	efile, err := os.Open(path)
	if err != nil {
		fmt.Printf("could not open file for exif decoder: %s", path)
		return
	}
	_, format, err := image.Decode(efile)
	if err != nil {
		fmt.Printf("文件信息 %s %s :%s \n",format,efile.Name(),err.Error())
		return
	}
	fmt.Printf("efile.Name() : %s ", efile.Name())
	fmt.Printf(" format : %s ", format)
	x, _ := exif.Decode(efile)
	if x != nil {
		orient, _ := x.Get(exif.Orientation)
		if orient != nil {
			fmt.Printf("%s had orientation %s \n", path, orient.String())
		}
	}else{
		fmt.Printf("  %s\n", path)
	}
}

func Test_Resize(t *testing.T) {
	//ScaleImage("/Users/mac/Desktop/5_03bf5931ffcd43068197154706049452_IMG_20210218_181450.jpg")
	imageCompress(
		func() (io.Reader, error) {
			return os.Open("/Users/mac/Desktop/5_03bf5931ffcd43068197154706049452_IMG_20210218_181450.jpg")
		},
		func() (*os.File, error) {
			return os.Open("/Users/mac/Desktop/5_03bf5931ffcd43068197154706049452_IMG_20210218_181450.jpg")
		},
		"/Users/mac/Desktop/5_03bf5931ffcd43068197154706049452_IMG_20210218_181450_1.jpg",
		50,
		1000,
		"jpg")
}

func ScaleImage(p string) {
	open, err := os.Open(p)
	defer open.Close()
	if err != nil {
		print(err.Error())
		return
	}
	decode, s, err := image.Decode(open)
	if err != nil {
		print(err.Error())
		return
	}
	if s != "jpg" && s != "jpeg" && s != "png" {
		return
	}

	buf := bytes.Buffer{}
	fmt.Println(s)
	//stat, _ := open.Stat()
	//fmt.Println(stat.Name())
	//fmt.Println(open.Name())
	//fmt.Println(filepath.Base(open.Name()))
	//fmt.Println(filepath.Ext(open.Name()))
	//fmt.Println(filepath.Clean(open.Name()))
	//fmt.Println(filepath.Abs(open.Name()))
	//fmt.Println(filepath.VolumeName(open.Name()))
	fmt.Println(strings.TrimSuffix(filepath.Base(open.Name()), filepath.Ext(open.Name())))
	suffix := strings.TrimSuffix(filepath.Base(open.Name()), filepath.Ext(open.Name()))

	if s == "jpg" || s == "jpeg" {
		err = jpeg.Encode(&buf, decode, &jpeg.Options{Quality: 5})
		if err != nil {
			return
		}
	} else if s == "png" {
		newImg := image.NewRGBA(decode.Bounds())
		draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
		draw.Draw(newImg, newImg.Bounds(), decode, decode.Bounds().Min, draw.Over)

		err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: 20})
		if err != nil {
			return
		}

	} else if s == "bmp" {
		//img := origin.(*image.RGBA)
		//subImg := decode.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
	}
	create, _ := os.Create("/Users/mac/Desktop/testimage222/" + suffix + "." + s)
	create.Write(buf.Bytes())
	create.Close()

	//fmt.Printf(s,"  bounds",decode.Bounds().Max.X, decode.Bounds().Max.Y)
}

func GetFileounds(path string) (width int, height int) {
	open, err := os.Open(path)
	//open, err := os.Open("/Users/mac/Desktop/1589191894238_8130.png")
	//open, err := os.Open("1589191894238_8130.png")
	if err != nil {
		print(err.Error())
		return
	}
	config, _, err := image.DecodeConfig(open)
	if err != nil {
		print(err.Error())
		return
	}

	return config.Width, config.Height

}
