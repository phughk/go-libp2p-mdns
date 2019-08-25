package mdns

import (
	reuse "github.com/libp2p/go-reuseport"
	"net"
	"fmt"
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
	packetConn, _ := reuse.ListenPacket("udp4", "224.0.0.251:5353")
	conn := &Connection {
		packetConn4: packetConn,
		shutdown: false,
	}
	conn.Poll()
	return conn
}

func (c *Connection)sendQuerry() error {
	query, err := BuildQuerry()
	if err != nil {
		return err
	}
	c.packetConn4.WriteTo(query, ipv4Addr)
	return nil
}

func (c *Connection)Poll() error {
	c.sendQuerry()
	c.readPackets()
	return nil
}

// recv is a long running routine to receive packets from an interface
func (c *Connection) readPackets() {
	if c == nil {
		return
	}
	buf := make([]byte, 65536)
	for !c.shutdown {
		n, from, err := c.packetConn4.ReadFrom(buf)
		if err != nil {
			print(err)
		}
		if err := c.parsePacket(buf[:n], from); err != nil {
			logf("[ERR] mdns: Failed to handle query: %v", err)
		}
	}
}

func (c * Connection) parsePacket(buf []byte, from net.Addr) error {
	fmt.Print("Apple")
	return nil
}