package syncfile

import (
	"os"
	"testing"
)

func Test_Scanfile(t *testing.T) {
	StartScanFile("/Users/wisn/Desktop/video","/Users/wisn/Desktop/video","/Users/wisn/Desktop/video2", func(s string, info os.FileInfo, i int) {

	})
}

