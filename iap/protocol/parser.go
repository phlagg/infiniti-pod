package protocol

import (
	"encoding/hex"

	"github.com/phlagg/infiniti-pod/iap/transport"
)

type RecieveState int

const (
	AwaitSync = iota
	AwaitStart
	AwaitLength
	AwaitData
	AwaitChecksum
	Reset
)

var readState RecieveState = AwaitSync

var dataBuf = make([]byte, MaxPayloadSize)
var idx uint8 = 0
var dataLen = 0
var payloadLength uint8 = 0
var sum = 0

// ReadCarPacket scans the incoming UART buffer look for valid iAP protocol patterns
func ReadPacket() (*Packet, bool) {
	if transport.Buffered() <= 0 {
		return nil, false
	}

	b, _ := transport.ReadByte()

	switch readState {
	case AwaitSync:
		if b != SyncByte {
			return nil, false
		}
		// println("[PARSER] Found Sync Byte")
		readState = AwaitStart
		break

	case AwaitStart:
		if b == SyncByte {
			// println("[PARSER] Found Second Sync Byte")
			b, _ = transport.ReadByte()
		}
		if b != PacketStartByte {
			return nil, false
		}
		// println("[PARSER] Found Packet Start Byte")
		readState = AwaitLength
		break

	case AwaitLength:
		payloadLength = uint8(b)
		sum += int(payloadLength)
		// println("[PARSER] Found Payload Length:", payloadLength)
		readState = AwaitData
		break

	case AwaitData:
		dataBuf[idx] = b
		idx++
		sum += int(b)
		if idx == payloadLength {
			// println("[PARSER] Found Payload")
			readState = AwaitChecksum
		}
		break

	case AwaitChecksum:
		// println("[PARSER] Found Checksum:", b)
		calcChecksum := byte(0x100 - (sum)&0xFF)
		println("[UART_RAW]", hex.EncodeToString(dataBuf[:payloadLength]))
		if b != calcChecksum {
			println("[CAR_ERR] Bad packet checksum match dropped. Calculated:", calcChecksum, "Got:", b)
			resetState()
			break
		}

		payloadCopy := make([]byte, payloadLength-1)
		copy(payloadCopy, dataBuf[1:payloadLength])

		pkt := Packet{
			LingoID: dataBuf[0],
			Payload: payloadCopy,
		}
		resetState()
		return &pkt, true
	}

	return nil, false
}

func resetState() {
	// println("[PARSER] Resetting")
	clear(dataBuf)
	idx = 0
	sum = 0
	dataLen = 0
	payloadLength = 0
	readState = AwaitSync
}
