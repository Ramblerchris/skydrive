package broadcast

import (
	"fmt"
	"net"
)

//3个并发处理管道
var conngrouplist = make(chan bool, 3)

//需要优化
func StartUDPGroupV2(UDPServerSendPort int,UDPListenPort int) {
	//如果第二参数为nil,它会使用系统指定多播接口，但是不推荐这样使用
	addr, err := net.ResolveUDPAddr("udp", fmt.Sprintf("225.0.0.1:%d",UDPListenPort))
	if err != nil {
		fmt.Println(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Local: <%s> \n", listener.LocalAddr().String())
	for {
		conngrouplist <- true
		go dealUdpGroup(UDPServerSendPort,listener)
	}

}
func dealUdpGroup(UDPServerSendPort int,listener *net.UDPConn) {
	data := make([]byte, 1024)
	n, remoteAddr, err := listener.ReadFromUDP(data)
	if err != nil {
		fmt.Printf("error during read: %s", err)
	}
	fmt.Printf("<%s> %s\n", remoteAddr, data[:n])

	//ip := net.ParseIP("224.0.0.250")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: remoteAddr.IP, Port: UDPServerSendPort}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	defer conn.Close()
	if err != nil {
		fmt.Println(err)
	}
	//defer conn.Close()
	conn.Write([]byte("hello"))
	fmt.Printf("<%s>\n", conn.RemoteAddr())
	<-conngrouplist

}
