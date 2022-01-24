package godori_tcpnet

import (
	"log"
	"net"
	"sync/atomic"
)

func start_Network_Impl(clientConfig *NetworkConfig, networkFunctor SessionNetworkFunctors) {
	// defer PrintPanicStack()

	// 아래 함수가 호출되면 무한 대기에 들어간다
	_tcpSessionManager = newClientSessionManager(clientConfig, networkFunctor)
	_start_TCPServer_block(clientConfig, networkFunctor)
}

func stopListen_impl() {
	_ = _mClientListener.Close()
}

func _start_TCPServer_block(config *NetworkConfig, networkFunctor SessionNetworkFunctors) {
	// defer PrintPanicStack
	logInfo("", 0, "TCPServerStart - Start")

	var err error
	tcpAddr, _ := net.ResolveTCPAddr("tcp", config.BindAddress)
	_mClientListener, err = net.ListenTCP("tcp", tcpAddr)

	if err != nil {
		log.Fatal("Error starting TCP Server")
	}
	defer _mClientListener.Close()

	log.Println("Server Listen...")

	for {
		conn, _ := _mClientListener.Accept()
		client := &TcpSession{
			SeqIndex:       SeqNumIncrement(),
			TcpConn:        conn,
			NetworkFunctor: networkFunctor,
		}

		_tcpSessionManager.addSession(client)
		go client.handleTcpRead(networkFunctor)
	}

	logInfo("", 0, "tcpServerStart - End")
}

// 보내기 함수(선언만 있는. 일종의 인터페이스)에 실제 동작함수를 연결한다
func _InitNetworkSendFunction() {
	NetLibSendToClient = sendToClient
	NetLibSendToAllClient = sendToAllClient
	NetLibPostSendToAllClient = postSendToAllClient
	NetLibPostSendToClient = postSendToClient

	logInfo("", 0, "call _InitNetworkSendFunction")
}

func sendToClient(sessionIndex int32, sessionUniqueID uint64, data []byte) bool {
	result := _tcpSessionManager.sendPacket(sessionIndex, sessionUniqueID, data)
	return result
}

func sendToAllClient(sendData []byte) {
	_tcpSessionManager.sendPacketAllClient(sendData)
}

func postSendToClient(sessionIndex int32, sessionUniqueID uint64, data []byte) bool {
	return sendToClient(sessionIndex, sessionUniqueID, data)
}

func postSendToAllClient(sendData []byte) {
	_tcpSessionManager.sendPacketAllClient(sendData)
}

func sendPacketToServer(sessionIndex int32, data []byte) bool {
	return false
}

func postSendPacketToServer(sessionIndex int32, data []byte) bool {
	return false
}

var _seqNumber uint64 // 절대 이것을 바로 사용하면 안 된다!!!

func SeqNumIncrement() uint64 {
	newValue := atomic.AddUint64(&_seqNumber, 1)
	return newValue
}

var _tcpSessionManager *tcpClientSessionManager
var _mClientListener *net.TCPListener
