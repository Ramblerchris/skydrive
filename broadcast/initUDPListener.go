package broadcast

import "github.com/skydrive/config"

//监听UDP内网 网络广播
func InitUDP()  {
	go StartUDPServerV2(config.UDP_SERVER_ListenPORT)
	//go broadcast.StartUDPGroup(config.UDP_SERVER_ListenPORT)
	go StartUDPGroupV2(config.UDP_GroupSERVER_SendPORT,config.UDP_GroupSERVER_ListenPORT)
}