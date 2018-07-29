package main

import (
	"os"

	"github.com/dazeus/dazeus-go"
)

func main() {
	connStr := "unix:/tmp/dazeus.sock"
	if len(os.Args) > 1 {
		connStr = os.Args[1]
	}

	dz, err := dazeus.ConnectWithLoggingToStdErr(connStr)
	if err != nil {
		panic(err)
	}

	networks, _ := dz.Networks()
	// Change this to your needs. It's hardcoded, yeah. Bite me.
	ListenForMessages("127.0.0.1:13337", networks[0], "#example", dz)
}
