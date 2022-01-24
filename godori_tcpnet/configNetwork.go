package godori_tcpnet

import (
	"fmt"
)

type NetworkConfig struct {
	IsTcp4Addr           bool
	BindAddress          string
	MaxSessionCount      int
	MaxPacketSize        int
	MaxReceiveBufferSize int
}

func (config NetworkConfig) WriteNetworkConfig(isClientSide bool) {
	logInfo("", 0, fmt.Sprintf("config - isClientSide : #{isClientSide}"))
	logInfo("", 0, fmt.Sprintf("config - IsTCP4Addr : #{config.IsTCP4Addr}"))
	logInfo("", 0, fmt.Sprintf("config - BindAddress : #{config.BindAddress}"))
	logInfo("", 0, fmt.Sprintf("config - MaxSessionCount : #{MaxSessionCount}"))
	logInfo("", 0, fmt.Sprintf("config - MaxPacketSize : #{MaxPacketSize}"))
	logInfo("", 0, fmt.Sprintf("config - MaxReceiveBufferSize : #{MaxReceiveBufferSize}"))
}
