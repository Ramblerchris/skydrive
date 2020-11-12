package redisconn

import (
	"fmt"
	"strconv"
	"testing"
	"time"

)

func Test_Client(t *testing.T)  {
	client := GetRedisClient()
	fmt.Println("client",client)
	for  i:=1;i<100;i++{
		go func(index int ) {
			itoa := strconv.Itoa(index)
			client.HMSet(CTX,"dd","key"+itoa,"value"+itoa, "keyT"+itoa,"valueT"+itoa)
		}(i)
	}
	time.Sleep(time.Second)
	/*
		client.HMSet(CTX,"aa","aaa31","aaa32")
		client.HMSet(CTX,"aa","aaa41","aaa42")
		client.HMSet(CTX,"aa","aaa61","aaa62")
		client.HMSet(CTX,"aa","aaa111","aaa112")
		client.HMSet(CTX,"bb","bbb881","bbb288")*/
	result, _ := client.HGetAll(CTX, "dd").Result()
	fmt.Println("result",result)
	for key ,value:=range result{
		fmt.Println(key,value)
	}
	/*for index:=0;index<len(result);index=index+2 {
		fmt.Println(result[index],result[index+1])
	}*/
}