package broadcast

import (
	"fmt"
	"github.com/skydrive/logger"
	"net"
)

//需要优化
// Deprecated: broadcast.StartUDPGroupV2 instead.
func StartUDPGroup(UDPListenPort int) {
	//如果第二参数为nil,它会使用系统指定多播接口，但是不推荐这样使用
	addr, err := net.ResolveUDPAddr("udp", "225.0.0.1:8998")
	if err != nil {
		logger.Error(err)
	}
	listener, err := net.ListenMulticastUDP("udp", nil, addr)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(fmt.Sprintf("Local: <%s> \n", listener.LocalAddr().String()))
	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := listener.ReadFromUDP(data)
		if err != nil {
			logger.Error("error during read: %s", err)
		}
		logger.Info(fmt.Sprintf("<%s> %s\n", remoteAddr, data[:n]))

		//ip := net.ParseIP("224.0.0.250")
		srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
		dstAddr := &net.UDPAddr{IP: remoteAddr.IP, Port: 8999}
		conn, err := net.DialUDP("udp", srcAddr, dstAddr)
		if err != nil {
			logger.Error(err)
		}
		//defer conn.Close()
		conn.Write([]byte("hello"))
		logger.Info( conn.RemoteAddr())
	}

	/*//1. 得到一个interface
	en4, err := net.InterfaceByName("en4")
	if err != nil {
		logger.Info(err)
	}
	group := net.IPv4(224, 0, 0, 250)
	//2. bind一个本地地址
	c, err := net.ListenPacket("udp4", "0.0.0.0:1024")
	if err != nil {
		logger.Info(err)
	}
	defer c.Close()
	//3.
	p := ipv4.NewPacketConn(c)
	if err := p.JoinGroup(en4, &net.UDPAddr{IP: group}); err != nil {
		logger.Info(err)
	}
	//4.更多的控制
	if err := p.SetControlMessage(ipv4.FlagDst, true); err != nil {
		logger.Info(err)
	}
	//5.接收消息
	b := make([]byte, 1500)
	for {
		n, cm, src, err := p.ReadFrom(b)
		if err != nil {
			logger.Info(err)
		}
		if cm.Dst.IsMulticast() {
			if cm.Dst.Equal(group) {
				fmt.Printf("received: %s from <%s>\n", b[:n], src)
				n, err = p.WriteTo([]byte("world"), cm, src)
				if err != nil {
					logger.Info(err)
				}
			} else {
				logger.Info("Unknown group")
				continue
			}
		}
	}
*/
}

