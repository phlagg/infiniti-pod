package car

import (
	"encoding/hex"
	"machine"
)

const (
	SyncByte        byte = 0xFF
	PacketStartByte byte = 0x55
)

// Configure the hardware UART interface
var Serial = machine.UART0

// InitializeSerialBus sets up the physical 19200 8N1 car line connection
// Standard legacy Apple iAP accessory baud rate
func InitializeSerialBus() error {
	Serial.Configure(machine.UARTConfig{BaudRate: 19200})
	return nil
}

// Packet Details  The sync byte (OxFF) is not considered part of the packet.
// It is sent merely to facilitate automatic baud rate  detection and correction when using a UART serial port link and, in some cases, to power on the iPod. It is  not necessary to send the sync byte when using BT or USB as a link.
// The packet payload length is the number of bytes in the packet, not including the sync byte, packet start  byte, packet payload length byte, or packet payload checksum byte. That is, it is the length of the command  ID, lingo, and command data. Thus, the packet payload data length for a Request Identify command  would be 0x02. The Lingo ID specifies the broad category that the communication falls under.
// The Command  ID is a more specific indication of the significance of the packet and is interpreted differently depending on  the Lingo ID.

// ReadCarPacket scans the incoming UART buffer look for valid iAP protocol patterns
func ReadCarPacket() (byte, bool) {
	// 1. Look for the header sequence by shifting out dead bytes
	for Serial.Buffered() >= 7 {
		sync, _ := Serial.ReadByte()
		if sync != SyncByte {
			println("[CAR_ERR] Bad Sync Byte (Not 0xFF)")
			continue // Skip this junk byte and check the next one
		}
		b0, _ := Serial.ReadByte()
		if b0 == SyncByte {
			println("[CAR_INFO] Second Sync Byte (0xFF)")
			b0, _ = Serial.ReadByte()
		}

		if b0 != PacketStartByte {
			println("[CAR_ERR] Bad SOP Byte (Not 0x55)")
			continue // Skip this junk byte and check the next one
		}

		payloadLength, _ := Serial.ReadByte()

		// 2. Read the bytes into a temporary buffer to preserve them for logging
		lenBuf := 3 + int(payloadLength) // packet start byte + length byte + checksum byte + payloadLength
		packetBuf := make([]byte, lenBuf)
		packetBuf[0] = b0
		packetBuf[1] = payloadLength
		for i := 2; i < lenBuf; i++ {
			packetBuf[i], _ = Serial.ReadByte()
		}
		checksum := packetBuf[lenBuf-1]

		sum := int(payloadLength)
		for i := 2; i < lenBuf-1; i++ {
			sum += int(packetBuf[i])
		}

		// 3. IMMEDIATELY LOG THE RAW HEX (Crucial for debugging/research)
		// Output format example: [UART_RAW] 55 aa 03 02 00 00 01 fa
		println("[UART_RAW]", hex.EncodeToString(packetBuf))

		// 6. Simple Checksum Validation
		calcChecksum := byte(0x100 - (sum)&0xFF)

		if calcChecksum != checksum {
			println("[CAR_ERR] Bad packet checksum match dropped. Expected:", calcChecksum, "Got:", checksum)
			return 0, false
		}

		// 7. Verify it is a true Mode 2 button event
		// if mode == 0x02 && cmdH == 0x00 && cmdL == 0x00 {
		// 	return data, true
		// }
	}

	// println("[CAR_INF] Valid iAP packet received, but not a Mode 2 button event")
	return 0, false
}

func SendExtendedPacket(packet *[]byte, length uint16) error {

	packetBuf := make([]byte, length)
	packetBuf[0] = 0x55                       // iAP Header
	packetBuf[1] = 0x00                       // Extended Packet Flag
	packetBuf[2] = byte((length >> 8) & 0xFF) // Length MSB
	packetBuf[3] = byte(length & 0xFF)        // Length LSB
	packetBuf[4] = 0x04                       // Mode 4 (Extended / Display Remote)

	for i := uint16(0); i < length; i++ {
		packetBuf[5+i] = (*packet)[i]
	}

	packetBuf[5+length] = CalculateBufferChecksum(packetBuf, 5+length)

	_, err := Serial.Write(packetBuf)
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

	packetBuf[5+len(title)] = CalculateBufferChecksum(packetBuf, length)

	_, err := Serial.Write(packetBuf)
	if err != nil {
		return err
	}

	return nil
}

func CalculateBufferChecksum(buffer []byte, length uint16) byte {
	var byteSum byte
	for i := uint16(1); i < length; i++ {
		byteSum += (buffer)[i]
	}
	return -byteSum
}
