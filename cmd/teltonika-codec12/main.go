package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"teltonika-codec12-cli/parser"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "Usage: teltonika-codec12-cli <hexstring>")
	}

	hexStr := os.Args[1]

	// Decode hex string to []byte
	data, err := hex.DecodeString(hexStr)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Invalid hex string:", err)
		os.Exit(2)
	}

	// Parse packet
	packet, err := parser.ParseCodec12(data)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Parse error", err)
		os.Exit(3)
	}

	// Output as JSON
	json.NewEncoder(os.Stdout).Encode(packet)
}
