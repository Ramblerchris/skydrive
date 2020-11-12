package main

import (
	"bufio"
	"fmt"
	"github.com/skydrive/config"
	"net"
	"os"
	"strconv"
	"strings"
	"testing"
	"time"
)

func Test_join(t *testing.T)  {
	var array =[]string{"aaa","bbbb","ccc"}
	println(strings.Join(array,"','"))
//	aaa','bbbb','ccc

}
func Test_t(t *testing.T) {
	//testTime()
	//testUdp()
	TestRange(t)
}
func TestRange(t *testing.T) {
	var ma = make(map[string]string)
	for i := 0; i < 20; i++ {
		ma[strconv.Itoa(i)] = strconv.Itoa(i) + "value"
	}
	for k, v := range ma {
		fmt.Println("&k:", &k, " &v:", &v, "k:", k, " v:", v)
		go doWork(&k)
	}
	time.Sleep(time.Second * 2)
}
func doWork(value *string) {
	time.Sleep(time.Nanosecond)
	fmt.Println("doWork", *value)
}

//每次发送的最大长度
var SERVER_RECV_LEN = 100

func testUdp() {
	conn, _ := net.Dial("udp", "localhost:"+strconv.Itoa(config.UDP_SERVER_ListenPORT))

	defer conn.Close()

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()

		lineLen := len(line)

		n := 0
		for written := 0; written < lineLen; written += n {
			var toWrite string
			if lineLen-written > SERVER_RECV_LEN {
				toWrite = line[written : written+SERVER_RECV_LEN]
			} else {
				toWrite = line[written:]
			}

			n, _ = conn.Write([]byte(toWrite))

			fmt.Println("Write:", toWrite)

			msg := make([]byte, SERVER_RECV_LEN)
			n, _ = conn.Read(msg)

			fmt.Println("Response:", string(msg))
		}
	}
}

func testTime() {
	t := time.Now()
	fmt.Printf("%d-%d-%d %d:%d:%d 星期%d,一年第%d天 时间戳%d 当前秒时间戳%d 当前毫秒时间戳%d 当前纳秒时间戳%d,%d\n",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Weekday(), t.YearDay(), t.Unix(), t.UnixNano()/1e9, t.UnixNano()/1e6, t.UnixNano(), t.Nanosecond())
	fmt.Println(t.Format("时间：15:04:05"))
	fmt.Println(t.Format("t 日期：2006-01-02 时间：15:04:05"))
	t2 := t.Add(time.Second * 4)
	//t2 := t.AddDate(10, 3, 3)
	fmt.Println(t2.Format("taaaa日期：2006-01-02 时间：15:04:05"))
	//t3 := t2.AddDate(-2, -2, -3)
	//fmt.Println(t3.Format("t3 日期：2006-01-02 时间：15:04:05"))
	//fmt.Println(t3.After(t2))
}
