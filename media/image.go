package media

import (
	"bytes"
	"github.com/disintegration/imaging"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/sirupsen/logrus"
	"github.com/skydrive/logger"
	"image"
	"image/color"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io"
	"os"
)

func ScaleImageQualityV1(p string, target string,quality int ) (isSuccess bool) {
	open, err := os.Open(p)
	defer open.Close()
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	decode, s, err := image.Decode(open)
	if err != nil {
		logger.Error("文件信息%s :%d",open.Name(),err.Error())
		return false
	}
	if s != "jpg" && s != "jpeg" && s != "png" {
		return false
	}

	buf := bytes.Buffer{}
	//logger.Info(s)
	//logger.Info(strings.TrimSuffix(filepath.Base(open.Name()), filepath.Ext(open.Name())))
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

	}
	create, _ := os.Create(target)
	create.Write(buf.Bytes())
	create.Close()
	return true
}
// 支持尺寸，质量压缩
func ScaleImageByWidthAndQuity(originPath string, targetWidth int, targetWidthFloat float64, targetQuality int, outputPath string) (isSuccess bool) {
	efile, err := os.Open(originPath)
	if err != nil {
		logrus.Warnf("could not open file for exif decoder: %s", originPath)
		return false
	}
	defer efile.Close()
	open, format, err := image.Decode(efile)
	if err != nil {
		logger.Error(err.Error())
		return false
	}
	if format != "jpg" && format != "jpeg" && format != "png" {
		logger.Error("图片格式不支持",format)
		return false
	}
	var config image.Config
	efile.Seek(0, io.SeekStart)
	if format == "jpg" || format == "jpeg" {
		config, err = jpeg.DecodeConfig(efile)
		if err != nil {
			logger.Error(err.Error())
			return false
		}
	} else if format == "png" {
		config, err = png.DecodeConfig(efile)
		if err != nil {
			logger.Error(err.Error())
			return false
		}
	}

	if targetWidth == 0 {
		targetWidth = int(targetWidthFloat * float64(config.Width))
	}
	if targetWidth > config.Width || targetWidth == 0 {
		targetWidth = config.Width
	}

	//fmt.Printf("%d ====== %d", targetWidth, config.Width)
	/** 做等比缩放 */
	width := targetWidth
	height := targetWidth * config.Height / config.Width
	//open, err := imaging.Open(originPath)
	thumb := imaging.Thumbnail(open, width, height, imaging.CatmullRom)
	dst := imaging.New(width, height, color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, thumb, image.Pt(0, 0))
	efile.Seek(0, io.SeekStart)
	x, _ := exif.Decode(efile)
	if x != nil {
		orient, _ := x.Get(exif.Orientation)
		if orient != nil {
			logrus.Infof("%s had orientation %s", originPath, orient.String())
			dst = reverseOrientation(dst, orient.String())
		} else {
			logrus.Warnf("%s had no orientation - implying 1", originPath)
			dst = reverseOrientation(dst, "1")
		}
	}
	buf := bytes.Buffer{}
	if format == "jpg" || format == "jpeg" {
		err = jpeg.Encode(&buf, dst, &jpeg.Options{Quality: targetQuality})
		if err != nil {
			return false
		}
	} else if format == "png" {
		newImg := image.NewRGBA(dst.Bounds())
		draw.Draw(newImg, newImg.Bounds(), &image.Uniform{C: color.White}, image.Point{}, draw.Src)
		draw.Draw(newImg, newImg.Bounds(), dst, dst.Bounds().Min, draw.Over)
		err = jpeg.Encode(&buf, newImg, &jpeg.Options{Quality: targetQuality})
		if err != nil {
			return false
		}
	}
	create, _ := os.Create(outputPath)
	create.Write(buf.Bytes())
	create.Close()
	return true
}

func reverseOrientation(img image.Image, o string) *image.NRGBA {
	switch o {
	case "1":
		return imaging.Clone(img)
	case "2":
		return imaging.FlipV(img)
	case "3":
		return imaging.Rotate180(img)
	case "4":
		return imaging.Rotate180(imaging.FlipV(img))
	case "5":
		return imaging.Rotate270(imaging.FlipV(img))
	case "6":
		return imaging.Rotate270(img)
	case "7":
		return imaging.Rotate90(imaging.FlipV(img))
	case "8":
		return imaging.Rotate90(img)
	}
	logrus.Errorf("unknown orientation %s, expect 1-8", o)
	return imaging.Clone(img)
}


