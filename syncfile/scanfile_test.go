package syncfile

import (
	"os"
	"testing"
)

func Test_Scanfile(t *testing.T) {
	StartScanFile("/Users/mac/Desktop/Version", func(s string, info os.FileInfo, i int) {

	})
}

