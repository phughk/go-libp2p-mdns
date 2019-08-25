package mdns

import (
	ma "github.com/multiformats/go-multiaddr"
	peer "github.com/libp2p/go-libp2p-core/peer"
	"encoding/base32"
	"encoding/binary"
	// "encoding/hex"
	"strings"
	"time"
	"math"
	"math/rand"
	"fmt"
)

const (
	SERVICE_NAME = "_p2p._udp.local"
)

func nameAppend(result []byte, name string) ([]byte, error) {
	for _, element := range strings.Split(name, ".") {
		if len(element) > 256 || len(element) < 1 {
			return result, ErrInvalidServiceName
		}
		result = append(result, uint8(len(element)))
		for _, c := range element {
			result = append(result, byte(c))
		}
	} 
	result = append(result, 0)
	return result, nil
}

func u16Append(result []byte, val uint16) ([]byte) {
	r := make([]byte, 2)
	binary.BigEndian.PutUint16(r, val)
	result = append(result, r[0])
	result = append(result, r[1])
	return result
}

func u32Append(result []byte, val uint32) ([]byte) {
	r := make([]byte, 4)
	binary.BigEndian.PutUint32(r, val)
	result = append(result, r[0])
	result = append(result, r[1])
	result = append(result, r[2])
	result = append(result, r[3])
	return result
}


func BuildQuerry() ([]byte, error) {

	// Initialise the query with 12 byte header
	var query []byte

	// Append a 16 bit transiction id
	rand.Seed(time.Now().Unix())
	transictionId := uint16(rand.Int())
	query = u16Append(query, transictionId)

	// Append 0x0 for a regular Query
	query = u16Append(query, 0x0)

	// Number of Questions
	query = u16Append(query, 0x01)

	// Number of answers, authorities, and additionals.
	query = u16Append(query, 0x0)
	query = u16Append(query, 0x0)
	query = u16Append(query, 0x0)

	// Our single question: The Service Name
    query, err := nameAppend(query, SERVICE_NAME)
    if err != nil {
    	return nil, err
    }

    // Flags
	query = u16Append(query, 0x0c)
	query = u16Append(query, 0x01)


	if len(query) > 33 {
		return nil, ErrExcessBodySize
	}
	// fmt.Print(hex.EncodeToString(query))
	return query, nil
}

func peerNameFromID(peerID peer.ID) string{
	base32ID := base32.StdEncoding.EncodeToString([]byte(peerID.String()))
	return fmt.Sprintf("%s.%s", base32ID, SERVICE_NAME)
}


func appendCharacterString(out []byte, data string) []byte {
	out = append(out, '"');
	for _, chr := range data {
		if chr == '\\' {
            out = append(out, '\\');
            out = append(out, '\\');
        } else if chr == '"' {
            out = append(out, '\\');
            out = append(out, '"');
        } else {
            out = append(out, byte(chr));
        }
	}
	out = append(out, '"');
	return out
}

func BuildQueryResponse(
	transictionId uint16,
	peerID peer.ID,
	addresses []ma.Multiaddr,
	duration time.Duration,
) ([]byte, error) {

	var response []byte
	// Initialise the response with the id
	response = u16Append(response, transictionId)

	// For Answer Response
	response = u16Append(response, 0x8400)

    // Number of questions, answers, authorities, additionals.
	response = u16Append(response, 0x0)
	response = u16Append(response, 0x1)
	response = u16Append(response, 0x0)
	response = u16Append(response, uint16(len(addresses)))

    // Our single answer: The name.
    response, err := nameAppend(response, SERVICE_NAME)
    if err != nil {
    	return nil, err
    }

    // Flags.
	response = u16Append(response, 0x000c)
	response = u16Append(response, 0x0001)

    // TTL for the answer
    response = u32Append(response, uint32(duration.Seconds()))


    //Encode Id to Base58
    peerIDBase58 := peerID.String()

    //peerName as base 32 encoding of peer id
    peerName := peerNameFromID(peerID)
    var peerIdByte []byte
    peerIdByte, err = nameAppend(peerIdByte, peerName)
    if(len(peerIdByte) > 0xffff) {
    	return nil, ErrExcessBodySize
    }
    response = u16Append(response, uint16(len(peerIdByte)))
    response = arrayJoin(response, peerIdByte)

     // The TXT records for answers.
     var entries [][]byte
    for _, addr := range addresses {
	 	txtToSend := fmt.Sprintf("dnsaddr=%s/p2p/%s", addr.String(), peerIDBase58)
	 	var entry []byte
	 	entry = appendCharacterString(entry, txtToSend)
	 	fmt.Print(string(entry), "\n")
	 	entries = append(entries, entry)
    }
    response, _ = appendTxtRecord(response, peerName, uint32(duration.Seconds()), entries)
    // The DNS specs specify that the maximum allowed size is 9000 bytes.
    if len(response) > 9000 {
        return nil, ErrExcessBodySize
    }

    return response, nil
}

func arrayJoin(out []byte, entries []byte) []byte {
	for _, c:= range entries {
    	out = append(out, c)
    }
    return out
}

func appendTxtRecord(
	out []byte, 
	name []byte,
	ttlSecs uint32,
	entries [][]byte,
) ([]byte, error) {

	// Name
	out = arrayJoin(out, name)

	// Flags
	out = append(out, 0x00)
	out = append(out, 0x10)
	out = append(out, 0x80)
	out = append(out, 0x01)

	// TTL
	out = u32Append(out, ttlSecs)

	var buf []byte
	for entry:= range(entries) {
		if len(entry) > 256 {
			return nil, ErrExcessBodySize
		}

		buf = append(buf, uint8(len(entry)));
		buf = arrayJoin(buf, entry)		
	}

	if len(buf) == 0 {
		buf = append(buf, 0)
	}

	if len(buf) > math.MaxInt16 {
		return nil, ErrExcessBodySize
	}

	out = u16Append(out, len(buf))
    out = arrayJoin(out, buf);

	return out

}

