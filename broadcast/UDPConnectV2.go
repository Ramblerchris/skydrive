package broadcast

import (
	"github.com/skydrive/logger"
	"net"
	"os"
	"strconv"
	"strings"
)

//3个并发处理管道
var connlist = make(chan bool, 3)

//需要优化
func StartUDPServerV2(UDPListenPort ,sendPort int) {
	address := ":" + strconv.Itoa(UDPListenPort)
	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}
	defer conn.Close()
	for {
		connlist <- true
		go dealRead(conn,sendPort)
	}

}
func dealRead(conn *net.UDPConn,sendport int ) {
	//defer  conn.Close()
	data := make([]byte, 65507)
	n, rAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		logger.Error(err)
	}

	strData := string(data[:n])
	logger.Info("Received:", strData, rAddr)
	//指定客户端端口
	upper := strings.ToUpper(strData)
	//10s 后给客户端再回复消息
	//time.Sleep(time.Second*1)
	rAddr.Port=sendport
	_, err = conn.WriteToUDP([]byte("pong "+upper), rAddr)
	if err != nil {
		logger.Error(err)
	}
	logger.Info("Send:", upper)
	<-connlist
}
