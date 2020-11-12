package utils

import "fmt"

const hour = 3600 * 1000
const minute = 60 * 1000

func GetTimeStr(time int)string  {
	var hours int
	var minutet int
	var second int
	if time > hour {
		hours = time / hour
		time = time - hours*hour
	}
	if time > minute {
		minutet = time / minute
		time = time - minutet*minute
	}
	second = time / 1000
	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%02d", hours, minutet, second)
	}
	if minutet > 0 {
		return fmt.Sprintf("00:%02d:%02d", minutet, second)
	}
	return fmt.Sprintf("00:%02d:%02d", minutet, second)
}
