package main

import (
	"flag"
	"fmt"

	"github.com/dazeus/dazeus-go"
)

func main() {
	var toDaZeus = flag.String("dzconn", "unix:/tmp/dazeus.sock", "Set the connection parameters for DaZeus")
	var listenAddr = flag.String("listen", "127.0.0.1:13337", "Listen on this UDP address for MediaWiki messages")
	var toChannel = flag.String("channel", "#example", "Send the messages to this channel on the first network")
	flag.Parse()
	fmt.Printf("Connecting to DaZeus on %s\n", *toDaZeus)

	dz, err := dazeus.ConnectWithLoggingToStdErr(*toDaZeus)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Going to listen on udp:%s, sending incoming messages to channel %s\n", *listenAddr, *toChannel)

	networks, _ := dz.Networks()
	// Change this to your needs. It's partly hardcoded, yeah. Bite me.
	ListenForMessages(*listenAddr, networks[0], *toChannel, dz)
}
