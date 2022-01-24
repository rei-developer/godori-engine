module main

go 1.17

require (
    godori.net/tcpnet v0.0.0
)

replace (
    godori.net/tcpnet v0.0.0 => ./godori_tcpnet
)