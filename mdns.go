package mdns

import (
	reuse "github.com/libp2p/go-reuseport"
	"net"
)

const (
	ipv4mdns              = "224.0.0.251"
	ipv6mdns              = "ff02::fb"
	mdnsPort              = 5353
)

var (
	ipv4Addr = &net.UDPAddr{
		IP:   net.ParseIP(ipv4mdns),
		Port: mdnsPort,
	}
	ipv6Addr = &net.UDPAddr{
		IP:   net.ParseIP(ipv6mdns),
		Port: mdnsPort,
	}
)

type Connection struct {
	packetConn4 net.PacketConn
	shutdown bool
}

func Init() *Connection{
	packetConn, err := reuse.ListenPacket("udp4", "0.0.0.0:5353")
	if err != nil {
		logf("[ERR] mdns: Failed to start udp4: %v", err)
	}
	conn := &Connection {
		packetConn4: packetConn,
		shutdown: false,
	}
	conn.Poll()
	return conn
}

