package packet

import (
	"flag"
	"fmt"
	"strconv"

	. "godori.net/module/godori_tcpnet"
)

func NetworkInit() {
	NetLibInitLog(LOG_LV_DEBUG, nil)
	netConfigClient := parseAppConfig()
	netConfigClient.WriteNetworkConfig(true)

	networkFunctors := SessionNetworkFunctors{}
	networkFunctors.OnConnect = OnConnect
	networkFunctors.OnClose = OnClose
	networkFunctors.OnReceive = OnReceive
	networkFunctors.OnReceiveBufferedData = nil
	networkFunctors.PacketTotalSizeFunc = PacketTotalSizeFunc
	networkFunctors.PacketHeaderSize = PACKET_HEADER_SIZE
	networkFunctors.IsConnectSession = true

	NetLibStartNetwork(&netConfigClient, networkFunctors)
}

//--------------------------------------------------------

func parseAppConfig() NetworkConfig {
	client := NetworkConfig{}

	flag.BoolVar(&client.IsTcp4Addr, "c_IsTcp4Addr", true, "bool flag")
	flag.StringVar(&client.BindAddress, "c_BindAddress", "127.0.0.1:11021", "string flag")
	flag.IntVar(&client.MaxSessionCount, "c_MaxSessionCount", 10, "int flag")
	flag.IntVar(&client.MaxPacketSize, "c_MaxPacketSize", 100, "int flag")

	flag.Parse()
	return client
}

func OnConnect(sessionIndex int32, SeqIndex uint64) {
	OutputLog(LOG_LV_INFO, "", 0, fmt.Sprintf("[OnConnect] sessionIndex: %d", sessionIndex))
}

func OnClose(id int32, pw uint64) {
	fmt.Println(id)
	fmt.Println(pw)
}

// End Packet Byte Code : 124
func PacketTotalSizeFunc(data []byte) int16 {
	parseHeader, _ := strconv.Atoi(string(data[:PACKET_HEADER_SIZE]))
	header := int16(parseHeader)
	// fmt.Println("TOTAL SIZE : ", data)
	// fmt.Println("HEADER SIZE : ", header)
	// fmt.Println("BYTE DATA : ", PACKET_HEADER_SIZE+header-1, data[PACKET_HEADER_SIZE+header-1])
	if data[PACKET_HEADER_SIZE+header-1] == 124 {
		return header + PACKET_HEADER_SIZE
	} else {
		return 0
	}
}
