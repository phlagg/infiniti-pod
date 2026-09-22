package iap

import (
	"testing"

	"github.com/phlagg/infiniti-pod/transport/serial"
)

var fakeBuf []byte
var fakeIdx int

// fake serial
func fakeBuffered() int {
	return len(fakeBuf) - fakeIdx
}

func fakeReadByte() (byte, error) {
	b := fakeBuf[fakeIdx]
	fakeIdx++
	return b, nil
}
func init() {
	serial.Buffered = fakeBuffered
	serial.ReadByte = fakeReadByte
}

func TestReadPacket_Valid(t *testing.T) {
	resetState()

	// Sync, Start, Length=3, LingoID=0x02, Payload=0xAA, 0xBB, Checksum
	fakeBuf = []byte{
		0xFF, // sync
		0x55, // start
		0x03, // payload length
		0x00, // lingo id
		0x01, // payload[0]
		0x04, // payload[1]
		0xF8, // correct checksum you compute
	}
	fakeIdx = 0

	pkt, ok := ReadPacket()
	if !ok {
		t.Fatalf("expected packet")
	}
	if pkt.LingoID != 0x02 {
		t.Fatalf("wrong lingo")
	}
	if len(pkt.Payload) != 2 {
		t.Fatalf("wrong payload length")
	}
	if pkt.Payload[0] != 0xAA {
		t.Fatalf("wrong payload")
	}
}

func TestReadPacket_BadChecksum(t *testing.T) {
	resetState()

	fakeBuf = []byte{
		0xFF,
		0x55,
		0x03,
		0x02,
		0xAA,
		0xBB,
		0x00, // wrong checksum
	}
	fakeIdx = 0

	pkt, ok := ReadPacket()
	if ok {
		t.Fatalf("expected failure")
	}
}
func TestReadPacket_IgnoresGarbage(t *testing.T) {
	resetState()

	fakeBuf = []byte{
		0x01, 0x02, 0x03, // garbage
		0xFF, // sync
		0x55, // start
		0x03, // length
		0x00, // lingo
		0x00, // cmd
		0xFD, // checksum
	}
	fakeIdx = 0

	pkt, ok := ReadPacket()
	if !ok {
		t.Fatalf("expected packet")
	}
}
