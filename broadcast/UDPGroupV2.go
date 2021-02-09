package broadcast

import (
	"fmt"
	"github.com/skydrive/config"
	"github.com/skydrive/logger"
	"net"
	"strings"
)

//3个并发处理管道
var conngrouplist = make(chan bool, 3)

//需要优化
func StartUDPGroupV2(UDPListenPort int,UDPServerSendPort int) {
	//如果第二参数为nil,它会使用系统指定多播接口，但是不推荐这样使用
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("225.0.0.1:%d",UDPListenPort))
	if err != nil {
		logger.Error(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Infof("Local: <%s> ", listener.LocalAddr().String())
	for {
		conngrouplist <- true
		go dealUdpGroup(UDPServerSendPort,listener)
	}

}
func dealUdpGroup(UDPServerSendPort int,listener *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := listener.ReadFromUDP(data)
	if err != nil {
		logger.Errorf("error during read: %s", err)
	}
	logger.Infof("<%s> %s\n", remoteAddr, data[:n])

	//ip := net.ParseIP("224.0.0.250")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: remoteAddr.IP, Port: UDPServerSendPort}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	defer conn.Close()
	if err != nil {
		logger.Error(err)
	}
	//defer conn.Close()
	conn.Write([]byte(fmt.Sprintf("{\"debug\":%t ,\"message\":\"pong %s\"}",config.Debug, strings.ToUpper(string(data[:n])))))
	logger.Infof("<%s>\n", conn.RemoteAddr())
	<-conngrouplist

}
