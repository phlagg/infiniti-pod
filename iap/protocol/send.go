package protocol

import (
	"encoding/hex"

	"github.com/phlagg/infiniti-pod/iap/transport"
)

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
	println("[SEND]", hex.EncodeToString(*pkt))
}

// func SendPacketUSB(pkt []byte) error {
// 	_, err := usb.Write(pkt)
// 	return err
// }
