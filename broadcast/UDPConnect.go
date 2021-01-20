package broadcast

import (
	"github.com/skydrive/logger"
	"net"
	"os"
	"strconv"
	"strings"
)
//需要优化
// Deprecated: broadcast.broStartUDPServerV2 instead.
func StartUDPServer(UDPListenPort int )  {
	address :=  ":" + strconv.Itoa(UDPListenPort)
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
		data := make([]byte, 65507)
		_, rAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			logger.Error(err)
			continue
		}
		strData := string(data)
		logger.Info("Received:", strData, rAddr)
		//指定客户端端口
		//rAddr.Port=SEND_PORT
		upper := strings.ToUpper(strData)
		//10s 后给客户端再回复消息
		//time.Sleep(time.Second*10)
		logger.Info("aaa:", len(upper))
		_, err = conn.WriteToUDP([]byte("pong"), rAddr)
		if err != nil {
			logger.Error(err)
			continue
		}
		logger.Info("Send:", upper)
	}

}
