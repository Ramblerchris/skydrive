package broadcast

import (
	"fmt"
	"github.com/skydrive/config"
	"net"
	"testing"
)

func Test_udp(t *testing.T)  {
	StartUDPServerV2(config.UDP_SERVER_ListenPORT)

	//broadcast.StartUDPGroup(config.UDP_SERVER_ListenPORT)
	StartUDPGroupV2(config.UDP_GroupSERVER_ListenPORT)

}
func Test_client( t *testing.T)  {
	ip := net.ParseIP("224.0.0.250")
	srcAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 0}
	dstAddr := &net.UDPAddr{IP: ip, Port: 9981}
	conn, err := net.DialUDP("udp", srcAddr, dstAddr)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()
	conn.Write([]byte("hello111"))
	fmt.Printf("<%s>\n", conn.RemoteAddr())

}
