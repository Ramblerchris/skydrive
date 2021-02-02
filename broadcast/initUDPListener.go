package broadcast

//监听UDP内网 网络广播
func InitUDP(listenPort,sendPort,groupListenPort,groupSendPort int )  {
	go StartUDPServerV2(listenPort,sendPort)
	go StartUDPGroupV2(groupListenPort,groupSendPort )
}