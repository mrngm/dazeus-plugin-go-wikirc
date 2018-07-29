package main

import (
	"bufio"
	"net"

	"github.com/dazeus/dazeus-go"
)

func ListenForMessages(listenaddr, dznetwork, dzchannel string, dz *dazeus.DaZeus) {
	if listenaddr == "" {
		return
	}

	udpaddr, _ := net.ResolveUDPAddr("udp", listenaddr)
	conn, err := net.ListenUDP("udp", udpaddr)
	if err != nil {
		panic(err)
	}

	for {
		reader := bufio.NewReader(conn)

		// These lines are provided by a MediaWiki instance through <https://www.mediawiki.org/wiki/Manual:%24wgRCFeeds>
		line, lerr := reader.ReadString('\n')
		if lerr != nil {
			panic(lerr)
		}
		dz.Message(dznetwork, dzchannel, line)
	}
}
