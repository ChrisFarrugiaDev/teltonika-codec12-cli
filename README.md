# teltonika-codec12-cli

CLI tool to parse Teltonika Codec 12 response packets  
GitHub: `chrisfarrugia.dev/teltonika-codec12-cli`

## Overview

`teltonika-codec12-cli` is a lightweight Go CLI tool for parsing Teltonika **Codec 12** packets.

It is designed mainly for parsing **GPRS command responses** returned by Teltonika devices, while keeping the output structure compatible with an existing parser format such as:

- `Packet`
- `Preamble`
- `Data_Length`
- `CodecID`
- `Quantity1`
- `CRC`
- `Quantity2`
- `CodecType`
- `Content`

You can:

- use it directly from your terminal/bash
- integrate it with your backend stack such as Node.js, TypeScript, PHP, or Python

## What does it do?

- Accepts a single hex string representing a raw Teltonika Codec 12 packet
- Parses **response packets** such as type `0x06`
- Outputs a parsed JSON representation to `stdout`
- Preserves backward-compatible field naming and structure

## Supported packet type

This tool currently focuses on **Codec 12 responses**.

According to the Teltonika protocol:

- `0x05` = command
- `0x06` = response

Example supported use case:

- parsing device response after sending a GPRS command such as `getinfo`, `getio`, or configuration-related commands

---

## Quick Start

### 1. Convert device data to hex string

Your Teltonika device response should be converted into a hex string before passing it to this tool.

### 2. Parse from CLI

In development:

```bash
go run ./cmd/teltonika-codec12/main.go 00000000000000560c01060000004e4e65772076616c75652034353030323a31323834373b34353233303a313b34353133303a313b34353134303a313b34353136303a313b34353238303a313b35303032303a313b34353138303a313b010000820b
```

After building:

```bash
go build -o teltonika-codec12-cli ./cmd/teltonika-codec12/main.go
./teltonika-codec12-cli 00000000000000560c01060000004e4e65772076616c75652034353030323a31323834373b34353233303a313b34353133303a313b34353134303a313b34353136303a313b34353238303a313b35303032303a313b34353138303a313b010000820b
```

Or, if installed globally:

```bash
teltonika-codec12-cli 00000000000000560c01060000004e4e65772076616c75652034353030323a31323834373b34353233303a313b34353133303a313b34353134303a313b34353136303a313b34353238303a313b35303032303a313b34353138303a313b010000820b
```

---

## Example Output

```json
{
  "Packet": "00000000000000560c01060000004e4e65772076616c75652034353030323a31323834373b34353233303a313b34353133303a313b34353134303a313b34353136303a313b34353238303a313b35303032303a313b34353138303a313b010000820b",
  "Preamble": 0,
  "Data_Length": 86,
  "CodecID": 12,
  "Quantity1": 1,
  "CRC": 33291,
  "Quantity2": 1,
  "CodecType": "GPRS messages",
  "Content": {
    "isResponse": true,
    "type": 6,
    "responseStr": "New value 45002:12847;45230:1;45130:1;45140:1;45160:1;45280:1;50020:1;45180:1;"
  }
}
```

---

## Protocol Note: Codec 12 session flow

Important:

This tool parses **Codec 12 packets only**. It does not manage the full Teltonika TCP session lifecycle.

In a normal Teltonika TCP communication flow:

1. the device connects to the server
2. the device sends IMEI
3. the server replies with one-byte IMEI acknowledgment (`0x01`)
4. the device sends AVL data packets
5. the server replies with AVL record acknowledgment
6. while the GPRS session stays active, Codec 12 commands may be sent
7. the device returns Codec 12 response packets

This CLI tool expects a **raw Codec 12 packet hex string only**.

Your application should handle:

- socket communication
- IMEI handshake
- AVL acknowledgments
- command sending
- session lifecycle

before passing the Codec 12 response packet to this parser.

---

## Teltonika Codec 12 structure

A typical Codec 12 response packet contains:

- Preamble
- Data Size
- Codec ID
- Response Quantity 1
- Type
- Response Size
- Response
- Response Quantity 2
- CRC

For Codec 12 responses:

- `Codec ID` is always `0x0C`
- `Type` is `0x06`

---

## Integration in Other Languages

You can call the CLI from any language that can launch shell commands.

### Node.js

```js
const { execFile } = require("child_process");

execFile("teltonika-codec12-cli", [hexString], (error, stdout) => {
  if (error) {
    console.error(error);
    return;
  }

  const parsed = JSON.parse(stdout);
  console.log(parsed);
});
```

### PHP

```php
$json = shell_exec("teltonika-codec12-cli $hexstring");  
echo $json;
```

### Python

```python
import subprocess  
  
result = subprocess.run(  
["teltonika-codec12-cli", hexstring],  
capture_output=True,  
text=True  
)  
  
print(result.stdout)

```

---

## TypeScript Usage Example

You can use this tool as a child process from Node.js / TypeScript:

```ts
import { execFile } from "node:child_process";

export function teltonikaCodec12Parse(hexString: string): Promise<TeltonikaCodec12Packet> {
    return new Promise((resolve, reject) => {
        execFile("teltonika-codec12-cli", [hexString], (error, stdout, stderr) => {
            if (error) return reject(error);

            try {
                resolve(JSON.parse(stdout));
            } catch (err) {
                reject(new Error("Failed to parse CLI output as JSON: " + err));
            }
        });
    });
}

export type TeltonikaCodec12Packet = {
    Packet: string;
    Preamble: number;
    Data_Length: number;
    CodecID: number;
    Quantity1: number;
    CRC: number;
    Quantity2: number;
    CodecType: string;
    Content: GPRSContent;
};

export type GPRSContent = {
    isResponse: boolean;
    type: number;
    responseStr: string;
};

```

---

## File Structure

.  
├── cmd/  
│   └── teltonika-codec12/  
│       └── main.go  
├── parser/  
│   ├── codec12.go  
│   ├── helpers.go  
│   └── types.go  
├── go.mod  
├── teltonika-codec12-cli  
└── README.md

---

## Building & Installation

Build the binary:

go build -o teltonika-codec12-cli ./cmd/teltonika-codec12/main.go

Optional: move the binary to a system path:

sudo mv teltonika-codec12-cli /usr/local/bin/  
sudo chmod 755 /usr/local/bin/teltonika-codec12-cli  
sudo chown root:root /usr/local/bin/teltonika-codec12-cli

Note: for production, avoid `chmod 777`.  
Use `chmod 755` and proper ownership as shown above.

---

## Notes

- This parser currently focuses on **response packets**
- It is intended to be simple, lightweight, and easy to integrate
- Output field names are intentionally kept compatible with an older parser structure
- CRC is parsed from the packet output, but full CRC validation may be handled separately depending on your implementation