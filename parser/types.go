package parser

type GPRS struct {
	IsResponse  bool   `json:"isResponse"`
	Type        byte   `json:"type"`        // 6 = response
	ResponseStr string `json:"responseStr"` // Decoded ASCII or "New value..." string
}

type ProtocolParser struct {
	Packet     string `json:"Packet"`
	Preamble   uint32 `json:"Preamble"`
	DataLength uint32 `json:"Data_Length"`
	CodecID    byte   `json:"CodecID"`
	Quantity1  byte   `json:"Quantity1"`
	CRC        uint32 `json:"CRC"`
	Quantity2  byte   `json:"Quantity2"`
	CodecType  string `json:"CodecType"`
	Content    GPRS   `json:"Content"`
}
