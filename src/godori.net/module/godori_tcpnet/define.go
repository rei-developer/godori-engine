package godori_tcpnet

const (
	MAX_RECEIVE_BUFFER_SIZE = 8012
	PACKET_HEADER_SIZE      = 3
	MAX_PACKET_SIZE         = 1024
)

const (
	NET_ERROR_NONE                                   = 0
	NET_ERROR_RECV_MAKE_PACKET_TOO_LARGE_PACKET_SIZE = 1
)

const (
	NET_CLOSE_REMOTE                     = 1
	NET_CLOSE_REMOTE_TOO_SMALL_RECV_DATA = 2
)

type SessionNetworkFunctors struct {
	OnConnect func(int32, uint64)

	OnClose func(int32, uint64)

	// 데이터 도착 이벤트
	OnReceive func(int32, uint64, []byte) bool

	// 데이터 도착 이벤트, []byte가 링버퍼에 저장되어 있음
	OnReceiveBufferedData func(int32, uint64, []byte) bool

	// 데이터를 분석하여 패킷의 크기를 반환
	PacketTotalSizeFunc func([]byte) int16

	// 패킷 헤더의 크기
	PacketHeaderSize int16

	// true이면 client 와 연결된 세션
	IsConnectSession bool
}
