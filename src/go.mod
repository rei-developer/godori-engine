module godori.net

go 1.17

require (
	godori.net/core v0.0.0
	godori.net/module/godori_tcpnet v0.0.0
)

require github.com/davecgh/go-spew v1.1.1 // indirect

replace (
	godori.net/core v0.0.0 => ./godori.net/core
	godori.net/module/godori_tcpnet v0.0.0 => ./godori.net/module/godori_tcpnet
)
