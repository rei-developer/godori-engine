package packet

import (
	"strconv"
	"strings"
)

func PacketDecode(packet string) map[string]interface{} {
	packetMap := make(map[string]interface{})
	dataLine := strings.Split(packet, "|")
	for _, str := range dataLine {
		kvLine := strings.Split(str, ":")
		if len(kvLine) < 2 {
			break
		}
		kvLine[1] = strings.Replace(kvLine[1], "\\cm", ":", -1)
		kvLine[1] = strings.Replace(kvLine[1], "\\v", "|", -1)
		if strings.Contains(kvLine[1], "\\i") {
			kvLine[1] = strings.Replace(kvLine[1], "\\i", "", -1)
			val, _ := strconv.ParseInt(kvLine[1], 0, 64)
			packetMap[kvLine[0]] = val
		} else if strings.Contains(kvLine[1], "\\f") {
			kvLine[1] = strings.Replace(kvLine[1], "\\f", "", -1)
			val, _ := strconv.ParseFloat(kvLine[1], 64)
			packetMap[kvLine[0]] = val
		} else {
			packetMap[kvLine[0]] = kvLine[1]
		}
	}
	return packetMap
}
