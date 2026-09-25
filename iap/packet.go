package iap

import (
	"encoding/hex"

	"github.com/phlagg/infiniti-pod/transport/serial"
)

const MaxPayloadSize = 255

const (
	SyncByte        byte = 0xFF
	PacketStartByte byte = 0x55
)

type RecievedPacket struct {
	LingoID byte
	Payload []byte
}

var pkt *RecievedPacket

type RecieveState uint8

const (
	AwaitSync = iota
	AwaitStart
	AwaitLength
	AwaitData
	AwaitChecksum
	Reset
)

var state RecieveState = AwaitSync

var dataBuf = make([]byte, MaxPayloadSize)
var idx uint8 = 0
var dataLen = 0
var payloadLength uint8 = 0
var sum = 0

// ReadCarPacket scans the incoming UART buffer look for valid iAP protocol patterns
func ReadPacket() (*RecievedPacket, bool) {
	if serial.Buffered() <= 0 {
		return nil, false
	}

	b, _ := serial.ReadByte()

	switch state {
	case AwaitSync:
		if b != SyncByte {
			return nil, false
		}
		// println("[PARSER] Found Sync Byte")
		state = AwaitStart
		break

	case AwaitStart:
		if b == SyncByte {
			// println("[PARSER] Found Second Sync Byte")
			b, _ = serial.ReadByte()
		}
		if b != PacketStartByte {
			return nil, false
		}
		// println("[PARSER] Found Packet Start Byte")
		state = AwaitLength
		break

	case AwaitLength:
		payloadLength = uint8(b)
		sum += int(payloadLength)
		// println("[PARSER] Found Payload Length:", payloadLength)
		state = AwaitData
		break

	case AwaitData:
		dataBuf[idx] = b
		idx++
		sum += int(b)
		if idx == payloadLength {
			// println("[PARSER] Found Payload")
			state = AwaitChecksum
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

		pkt := RecievedPacket{
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
	state = AwaitSync
}

func SendPacket(pkt *ResponsePacket) error {
	if err := sendPacketSerial(pkt); err != nil {
		return err
	}
	// if err := sendPacketUSB(pkt); err != nil {
	// 	return err
	// }

	return nil
}

func sendPacketSerial(pkt *ResponsePacket) error {
	// prepend sync byte
	buf := append([]byte{0xFF}, *pkt...)
	logSentPacket(&buf)
	_, err := serial.Write(buf)
	return err
}

func logSentPacket(pkt *[]byte) {
	println("[PACKET]", hex.EncodeToString(*pkt))
}

//	func SendPacketUSB(pkt ResponsePacket) error {
//		_, err := usb.Write(pkt)
//		return err
//	}
//

func logPacket(packet *ResponsePacket) {
	println("[PACKET]", hex.EncodeToString(*packet))
}
