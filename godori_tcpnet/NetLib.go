//  NetLib
//  외부 접근 가능 함수 정의
//

package godori_tcpnet

func NetLibInitLog(loglevel int, logFunc func(int, string, uint64, string)) {
	_logLv = loglevel

	if logFunc != nil {
		OutputLog = logFunc
	}
}

func NetLibStartNetwork(clientConfig *NetworkConfig, networkFunctor SessionNetworkFunctors) {
	start_Network_Impl(clientConfig, networkFunctor)
}

func NetLibStopNetwork() {
	stopListen_impl()
}

// 특정 클라이언트에게 패킷 전송
var NetLibSendToClient func(int32, uint64, []byte) bool

// 모든 클라이언트에게 패킷 전송
var NetLibSendToAllClient func([]byte)

// [Async] 특정 클라이언트에게 패킷 전송
var NetLibPostSendToClient func(int32, uint64, []byte) bool

// [Async] 모든 클라이언트에게 패킷 전송
var NetLibPostSendToAllClient func([]byte)

// 특정 클라이언트 접속 강제 종료
func NetLibForceDisconnectClient(sessionIndex int32, sessionUnqiueID uint64) {
	_tcpSessionManager.forceDisconnectClient(sessionIndex, sessionUnqiueID)
}
