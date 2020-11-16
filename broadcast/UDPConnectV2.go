package broadcast

import (
	"fmt"
	"github.com/skydrive/config"
	"net"
	"os"
	"strconv"
	"strings"
)

//3个并发处理管道
var connlist = make(chan bool, 3)

//需要优化
func StartUDPServerV2(UDPListenPort int) {
	address := ":" + strconv.Itoa(UDPListenPort)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		connlist <- true
		go dealRead(conn)
	}

}
func dealRead(conn *net.UDPConn) {
	//defer  conn.Close()
	data := make([]byte, 65507)
	n, rAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println(err)
	}

	strData := string(data[:n])
	fmt.Println("Received:", strData, rAddr)
	//指定客户端端口
	upper := strings.ToUpper(strData)
	//10s 后给客户端再回复消息
	//time.Sleep(time.Second*1)
	rAddr.Port=config.UDP_SERVER_SendPORT
	_, err = conn.WriteToUDP([]byte("pong "+upper), rAddr)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Send:", upper)
	<-connlist
}
