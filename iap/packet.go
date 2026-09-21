package iap

import (
	"encoding/hex"

	"github.com/phlagg/infiniti-pod/iap/transport"
)

const MaxPayloadSize = 255

const (
	SyncByte        byte = 0xFF
	PacketStartByte byte = 0x55
)

type Packet struct {
	LingoID byte
	Payload []byte
}

var pkt *Packet

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
func ReadPacket() (*Packet, bool) {
	if transport.Buffered() <= 0 {
		return nil, false
	}

	b, _ := transport.ReadByte()

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
			b, _ = transport.ReadByte()
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
	state = AwaitSync
}

func SendPacket(pkt []byte) error {
	err := sendPacketSerial(pkt)
	if err != nil {
		return err
	}
	// err = sendPacketUSB(pkt)
	// if err != nil {
	// 	return err
	// }

	return nil
}

func sendPacketSerial(pkt []byte) error {
	// prepend sync byte
	buf := append([]byte{0xFF}, pkt...)
	logSentPacket(&buf)
	_, err := transport.Write(buf)
	return err
}

func logSentPacket(pkt *[]byte) {
	println("[PACKET]", hex.EncodeToString(*pkt))
}

//	func SendPacketUSB(pkt []byte) error {
//		_, err := usb.Write(pkt)
//		return err
//	}
func BuildSmallPacket(lingo byte, cmd byte, cmdData []byte) []byte {
	pktPayloadLen := 2 + len(cmdData) // cmd + lingo+ cmdDataLength
	buffLen := 2 + pktPayloadLen + 1  // PacketStartByte + pktPayloadLen byte + []pktPayload + Checksum
	packetBuf := make([]byte, buffLen)
	packetBuf[0] = PacketStartByte
	packetBuf[1] = byte(pktPayloadLen)
	packetBuf[2] = lingo
	packetBuf[3] = cmd
	for i := 0; i < len(cmdData); i++ {
		packetBuf[4+i] = cmdData[i]
	}
	packetBuf[buffLen-1] = CalculateBufferChecksum(packetBuf[1:])
	return packetBuf
}

func BuildSmallExtendedPacket(cmd uint16, cmdData []byte) []byte {
	pktPayloadLen := 3 + len(cmdData) // 2 cmd bytes + lingo +  cmdDatalength
	buffLen := 2 + pktPayloadLen + 1  // PacketStartByte + pktPayloadLen byte + []pktPayload + Checksum
	packetBuf := make([]byte, buffLen)
	packetBuf[0] = PacketStartByte
	packetBuf[1] = byte(pktPayloadLen)
	packetBuf[2] = 0x04
	packetBuf[3] = byte(cmd >> 8)
	packetBuf[4] = byte(cmd & 0xFF)
	for i := 0; i < len(cmdData); i++ {
		packetBuf[5+i] = cmdData[i]
	}
	packetBuf[buffLen-1] = CalculateBufferChecksum(packetBuf[1:])
	// logPacket(&packetBuf)
	return packetBuf
}

func logPacket(packet *[]byte) {
	println("[PACKET]", hex.EncodeToString(*packet))
}

func CalculateBufferChecksum(buffer []byte) byte {
	var byteSum byte
	for i := 0; i < len(buffer); i++ {
		byteSum += (buffer)[i]
	}
	return -byteSum
}
