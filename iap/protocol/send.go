package protocol

import "github.com/phlagg/infiniti-pod/iap/transport"

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
	_, err := transport.Write(buf)
	return err
}

// func SendPacketUSB(pkt []byte) error {
// 	_, err := usb.Write(pkt)
// 	return err
// }

func CalculateBufferChecksum(buffer []byte) byte {
	var byteSum byte
	for i := 0; i < len(buffer); i++ {
		byteSum += (buffer)[i]
	}
	return -byteSum
}

func SendLargePacket(packet *[]byte, length uint16) error {

	packetBuf := make([]byte, length)
	packetBuf[0] = 0x55                       // iAP Header
	packetBuf[1] = 0x00                       // Extended Packet Flag
	packetBuf[2] = byte((length >> 8) & 0xFF) // Length MSB
	packetBuf[3] = byte(length & 0xFF)        // Length LSB
	packetBuf[4] = 0x04                       // Mode 4 (Extended / Display Remote)

	for i := uint16(0); i < length; i++ {
		packetBuf[5+i] = (*packet)[i]
	}

	packetBuf[5+length] = CalculateBufferChecksum(packetBuf)

	_, err := transport.Write(packetBuf)
	if err != nil {
		return err
	}

	return nil
}

func SendTrackTitlePacket(title string) error {

	var length uint16 = uint16(len(title)) + 3
	packetBuf := make([]byte, length)
	packetBuf[0] = 0x55         // iAP Header
	packetBuf[1] = byte(length) // Packet Length
	packetBuf[2] = 0x04         // Mode 4 (Extended / Display Remote)
	packetBuf[3] = 0x00         // Command MSB
	packetBuf[4] = 0x1A         // Command LSB: Set Title String

	for i := 0; i < len(title); i++ {
		packetBuf[5+i] = title[i]
	}

	packetBuf[5+len(title)] = CalculateBufferChecksum(packetBuf)

	_, err := transport.Write(packetBuf)
	if err != nil {
		return err
	}

	return nil
}
