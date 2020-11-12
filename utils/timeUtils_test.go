package utils

import (
	"fmt"
	"testing"
)

func Test_timeUtils( t *testing.T)  {
	fmt.Println(GetTimeStr(13000))
	fmt.Println(GetTimeStr(133*3600*1000+34*60*1000+4*1000+333))
	fmt.Println(GetTimeStr(13*3600*1000+34*60*1000+4*1000+333))
	fmt.Println(GetTimeStr(13*3600*1000+34*60*1000+24*1000+333))
	fmt.Println(GetTimeStr(13*3600*1000+4*1000+333))
	fmt.Println(GetTimeStr(13*3600*1000+34*60*1000+333))
	fmt.Println(GetTimeStr(13*3600*1000+333))
}