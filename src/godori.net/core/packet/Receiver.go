package packet

import (
	"fmt"
)

func OnReceive(sessionIndex int32, SeqIndex uint64, data []byte) bool {
	packet := string(data[3:])
	recv := PacketDecode(packet)
	fmt.Println("recv:", recv)
	fmt.Println("part:", recv["part"])
	fmt.Printf("id: %s %T\n", recv["id"], recv["id"])
	fmt.Printf("pw: %s %T\n", recv["pw"], recv["pw"])
	return true
}
