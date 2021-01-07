package media

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"os"
	"path/filepath"
	"strings"
)

func ScaleImageQuality(p string, target string,quality int ) (isSuccess bool) {
	open, err := os.Open(p)
	defer open.Close()
	if err != nil {
		print(err.Error())
		return false
	}
	decode, s, err := image.Decode(open)
	if err != nil {
		print(err.Error())
		return false
	}
	if s != "jpg" && s != "jpeg" && s != "png" {
		return false
	}

	buf := bytes.Buffer{}
	fmt.Println(s)
	fmt.Println(strings.TrimSuffix(filepath.Base(open.Name()), filepath.Ext(open.Name())))
	//suffix := strings.TrimSuffix(filepath.Base(open.Name()), filepath.Ext(open.Name()))

	if s == "jpg" || s == "jpeg" {
		err = jpeg.Encode(&buf, decode, &jpeg.Options{Quality: quality})
		if err != nil {
			return false
		}
	} else if s == "png" {
		newImg := image.NewRGBA(decode.Bounds())
		draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
		draw.Draw(newImg, newImg.Bounds(), decode, decode.Bounds().Min, draw.Over)

		err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: quality})
		if err != nil {
			return false
		}

	} else if s == "bmp" {
		//img := origin.(*image.RGBA)
		//subImg := decode.SubImage(image.Rect(x0, y0, x1, y1)).(*image.RGBA)
	}
	create, _ := os.Create(target)
	create.Write(buf.Bytes())
	create.Close()
	return true
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
