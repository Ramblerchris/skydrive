package media

import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
)

// VideoThumbnail source：是原始文件的名字，可以是mov，mpeg，avi，wmv各类格式，ffmpeg基本都支持。
//-s wxh: 指定视频的宽和高
//-b : 设定视频的比特率
//-aspect: 保持视频的比率。如4:3或者16:9
//-y : 如果目标文件存在时，直接覆盖原有的目标文件。
//-f : 指定转换的文件格式，这里是flv格式。（其实如果不指定文件格式，ffmpeg也会按文件的后缀名来进行转换）。
//dest: 转换的目标文件名字，并不一定需要是flv，可以是mov，mpeg以及其他的常用格式。
//参数说明：
//-L license
//-h 帮助
//-fromats 显示可用的格式，编解码的，协议的
//-f fmt 强迫采用格式fmt
//-I filename 输入文件
//-y 覆盖输出文件
//-t duration 设置纪录时间 hh:mm:ss[.xxx]格式的记录时间也支持
//-ss position 搜索到指定的时间 [-]hh:mm:ss[.xxx]的格式也支持
//s wxh: 指定视频的宽和高
//获取视频静态缩略图
func VideoThumbnail(videoPath, videoThumbNail string) bool {
	//ffmpeg -i /Users/mac/Desktop/media/video/video_ddd.mkv -y -f image2 -ss 8 -t 0.001  /Users/mac/Desktop/media/video/aaa.jpg
	//cmdArguments := []string{"-i", videoPath, "-y", "-f",
	//	"mjpeg", "-ss", "0.9", "-t", "0.001", videoThumbNail}
	cmdArguments := []string{"-i", videoPath, "-y", "-f",
		"mjpeg", "-ss", "3", "-t", "0.001", videoThumbNail}
	//cmdArguments := []string{"-i", videoPath, "-y", "-f", "-ss", "1", videoThumbNail}
	return execCommend("ffmpeg", cmdArguments)
}

// VideoThumbnailGif 获取视频gif 缩略图
func VideoThumbnailGif(videoPath, videoThumbNail string, countfram int) bool {
	cmdArguments := []string{"-i", videoPath, "-vframes", strconv.Itoa(countfram),
		"-y", "-f", "gif", videoThumbNail}
	return execCommend("ffmpeg", cmdArguments)
}

func ImageThumbnailJPG(quality, resizef int, videoPath, jpgehumbNail string) bool {
	//convert -quality 30 -resize 500 1627138076033.jpeg  1627138076033_2.jpeg
	cmdArguments := []string{"-quality", strconv.Itoa(quality), "-resize", fmt.Sprintf("%d%%", resizef),videoPath, jpgehumbNail}
	return execCommend("convert", cmdArguments)
}

func ImageThumbnailGif(fuzz int, videoPath, gifThumbNail string) bool {
	//convert test.gif -fuzz 5% -layers Optimize result.gif
	cmdArguments := []string{ videoPath, "-fuzz", fmt.Sprintf("%d%%", fuzz), "-layers", "Optimize", gifThumbNail}
	return execCommend("convert", cmdArguments)
}

func execCommend(name string, arguments []string) bool {
	cmd := exec.Command(name, arguments...)
	var out bytes.Buffer
	cmd.Stdout = &out
	fmt.Printf("command start %s %s \n ",name,arguments)
	err := cmd.Run()
	if err != nil {
		fmt.Printf("command err %s\n", err.Error())
		return false
	}
	fmt.Printf("command success : %q\n", out.String())
	return true
}
