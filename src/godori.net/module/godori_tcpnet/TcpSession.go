package godori_tcpnet

import (
	"net"
)

type TcpSession struct {
	Index          int32
	SeqIndex       uint64
	TcpConn        net.Conn
	NetworkFunctor SessionNetworkFunctors
}

func (session *TcpSession) handleTcpRead(networkFunctor SessionNetworkFunctors) {
	session.NetworkFunctor.OnConnect(session.Index, session.SeqIndex)

	var startRecvPos int16
	//var result int
	recvBuff := make([]byte, MAX_RECEIVE_BUFFER_SIZE)

	for {
		recvBytes, err := session.TcpConn.Read(recvBuff[startRecvPos:])
		if err != nil {
			session.closeProcess()
			return
		}

		if recvBytes < PACKET_HEADER_SIZE {
			session.closeProcess()
			return
		}

		readAbleByte := int16(startRecvPos) + int16(recvBytes)
		_, result := session.makePacket(readAbleByte, recvBuff)
		if result != NET_ERROR_NONE {
			session.closeProcess()
			return
		}

	}
}

func (session *TcpSession) makePacket(readAbleByte int16, recvBuff []byte) (int16, int) {
	sessionIndex := session.Index
	sessionUnique := session.SeqIndex

	PacketHeaderSize := session.NetworkFunctor.PacketHeaderSize
	PacketTotalSizeFunc := session.NetworkFunctor.PacketTotalSizeFunc
	var startRecvPos int16 = 0
	var readPos int16

	for {
		if readAbleByte < PacketHeaderSize {
			break
		}

		requireDataSize := PacketTotalSizeFunc(recvBuff[readPos:])

		if requireDataSize > readAbleByte {
			break
		}

		if requireDataSize > MAX_PACKET_SIZE {
			return startRecvPos, NET_ERROR_RECV_MAKE_PACKET_TOO_LARGE_PACKET_SIZE
		}

		ltvPacket := recvBuff[readPos:(readPos + requireDataSize)]
		readPos += requireDataSize
		readAbleByte -= requireDataSize

		session.NetworkFunctor.OnReceive(sessionIndex, sessionUnique, ltvPacket)
	}

	if readAbleByte > 0 {
		copy(recvBuff, recvBuff[readPos:(readPos+readAbleByte)])
	}

	startRecvPos = readAbleByte
	return startRecvPos, NET_ERROR_NONE
}

func (session *TcpSession) closeProcess() {
	session.TcpConn.Close()
	session.NetworkFunctor.OnClose(session.Index, session.SeqIndex)

	_tcpSessionManager.removeSession(session.Index, session.SeqIndex)
}

// Send bytes to client
func (session *TcpSession) sendPacket(b []byte) error {
	_, err := session.TcpConn.Write(b)
	return err
}

func (session *TcpSession) close() error {
	return session.TcpConn.Close()
}
