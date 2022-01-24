package main

import (
	"flag"

	. "godori.net/tcpnet"
)

func main() {
	NetLibInitLog(LOG_LV_DEBUG, nil)

	netConfigClient := parseAppConfig()
	netConfigClient.WriteNetworkConfig(true)

	createServer(netConfigClient)
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
