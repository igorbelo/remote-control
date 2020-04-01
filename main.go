package main

import (
    "flag"
		"github.com/igorbelo/remote-control/server"
)

func main() {
	addr := flag.String("addr", ":8080", "http service address")
	flag.Parse()

	server.Serve(*addr)
}
