package main

import (
	"flag"
	"fmt"

	. "godori.net/tcpnet"
)

func main() {
	NetLibInitLog(LOG_LV_DEBUG, nil)

	netConfigClient := parseAppConfig()
	netConfigClient.WriteNetworkConfig(true)

	networkFunctors := SessionNetworkFunctors{}
	networkFunctors.OnConnect = OnConnect
	networkFunctors.OnClose = OnClose
	networkFunctors.OnReceive = OnReceive
	networkFunctors.OnReceiveBufferedData = OnReceiveBufferedData
	networkFunctors.PacketTotalSizeFunc = PacketHeaderSizeFunc
	networkFunctors.PacketHeaderSize = 5
	networkFunctors.IsConnectSession = true

	NetLibStartNetwork(&netConfigClient, networkFunctors)
}

func parseAppConfig() NetworkConfig {
	client := NetworkConfig{}

	flag.BoolVar(&client.IsTcp4Addr, "c_IsTcp4Addr", true, "bool flag")
	flag.StringVar(&client.BindAddress, "c_BindAddress", "127.0.0.1:11021", "string flag")
	flag.IntVar(&client.MaxSessionCount, "c_MaxSessionCount", 0, "int flag")
	flag.IntVar(&client.MaxPacketSize, "c_MaxPacketSize", 0, "int flag")

	flag.Parse()
	return client
}

func OnConnect(sessionIndex int32, seq uint64) {
	OutputLog(LOG_LV_INFO, "", 0, fmt.Sprintf("[OnConnect] sessionIndex: %d", sessionIndex))
}
func OnClose(id int32, pw uint64) {
	fmt.Println(id)
	fmt.Println(pw)
}
func OnReceive(id int32, pw uint64, packet []byte) bool {
	fmt.Println(id)
	fmt.Println(pw)
	fmt.Println(packet)
	return true
}
func OnReceiveBufferedData(id int32, pw uint64, packet []byte) bool {
	fmt.Println(id)
	fmt.Println(pw)
	fmt.Println(packet)
	return true
}
func PacketHeaderSizeFunc(packet []byte) int16 {
	fmt.Println(packet)
	return 5
}
