package parser

import (
	"encoding/hex"
	"fmt"
)

func ParseCodec12(data []byte) (*ProtocolParser, error) {
	offset := 0
	// Check minimum length for header, type, response size, CRC
	if err := assertCanRead(data, offset, 4+4+1+1+1+4+1+4, "header"); err != nil {
		return nil, err
	}
	pkt := ProtocolParser{}
	pkt.Packet = hex.EncodeToString(data)

	pkt.Preamble = bytesToUint32(data[offset:])
	offset += 4
	pkt.DataLength = bytesToUint32(data[offset:])
	offset += 4
	pkt.CodecID = data[offset]
	offset += 1
	pkt.Quantity1 = data[offset]
	offset += 1

	typ := data[offset]
	offset += 1

	if typ != 0x06 {
		return nil, fmt.Errorf("Not a response packet (type != 6)")
	}

	respSize := bytesToUint32(data[offset:])
	offset += 4

	if err := assertCanRead(data, offset, int(respSize), "response"); err != nil {
		return nil, err
	}

	respRaw := data[offset : offset+int(respSize)]
	offset += int(respSize)

	quantity2 := data[offset]
	offset += 1

	crc := bytesToUint32(data[offset:])
	offset += 4

	// Try to decode response as ASCII; if not printable, fallback to hex
	var respStr string
	if isPrintableASCII(respRaw) {
		respStr = string(respRaw)
	} else {
		respStr = "New value " + bytesToHexPairsString(respRaw)
	}

	pkt.CRC = crc
	pkt.Quantity2 = quantity2
	pkt.CodecType = "GPRS messages"
	pkt.Content = GPRS{
		IsResponse:  true,
		Type:        typ,
		ResponseStr: respStr,
	}

	return &pkt, nil
}
