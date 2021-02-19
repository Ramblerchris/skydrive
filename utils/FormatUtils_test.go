package utils

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"
)

func Test_Format(t *testing.T) {
	numF := 0.2253
	// 保留两位小数, 通用
	value, _ := strconv.ParseFloat(fmt.Sprintf("%.3f", numF), 64)
	fmt.Println(reflect.TypeOf(value), value)

	//保留两位小数，舍弃尾数，无进位运算
	num, _ := FormatFloat(numF, 3)
	fmt.Println(reflect.TypeOf(num), num)


	// 舍弃的尾数不为0，强制进位
	num, _ = FormatFloatCeil(0.2295, 3)
	fmt.Println(reflect.TypeOf(num), num)


	// 强制舍弃尾数
	num, _ = FormatFloatFloor(0.2295, 3)
	fmt.Println(reflect.TypeOf(num), num)
}