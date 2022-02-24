package main

import (
	"godori.net/core/config"
	"godori.net/core/packet"
)

func main() {
	config.Init()
	packet.NetworkInit()
}
