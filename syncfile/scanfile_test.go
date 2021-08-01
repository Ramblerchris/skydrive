package syncfile

import (
	"os"
	"testing"
)

func Test_Scanfile(t *testing.T) {
	StartScanFile("/Users/wisn/Desktop/video", func(s string, info os.FileInfo, i int) {

	})
}

