package mdns 

import (
	peer "github.com/libp2p/go-libp2p-core/peer"
	ma "github.com/multiformats/go-multiaddr"
	"encoding/binary"
)


type Message struct {
	TransactionID uint16
	ServiceName string
	IsResponse bool
}

type Response struct {
	Ttl uint32
	PeerID peer.ID
	Multiaddresses []ma.Multiaddr 
}

func abstractName(message []byte) (string, error) {
	var result string
	pos := uint8(0)
	for message[pos] != 0 {
		pos += 1
		byteWord := message[pos : pos + message[pos - 1]]
		result += string(byteWord)  
		pos = pos + uint8(message[pos - 1])
		if (message[pos] != 0) {
			result += "."
		}
	}
	return result, nil
}

func unpackMessage(message []byte) Message {
	transactionID := binary.BigEndian.Uint16(message[:2])

	isResponse := true
	if (binary.BigEndian.Uint16(message[2:4]) == 0) {
		isResponse = false
	}
	serviceName, _ := abstractName(message[12:])
	return Message {
		TransactionID: transactionID,
		ServiceName: serviceName,
		IsResponse: isResponse,
	}
}

func unpackResponse(message []byte) Response {
	ttl := binary.BigEndian.Uint32(message[33:37])
	return Response {
		Ttl: ttl,
	}
}

