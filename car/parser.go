package car

import (
	"encoding/hex"
	"machine"
)

// Configure the hardware UART interface
var Serial = machine.UART0

// InitializeSerialBus sets up the physical 19200 8N1 car line connection
// Standard legacy Apple iAP accessory baud rate
func InitializeSerialBus() error {
	Serial.Configure(machine.UARTConfig{BaudRate: 19200})
	return nil
}

// ReadCarPacket scans the incoming UART buffer look for valid iAP protocol patterns
func ReadCarPacket() (byte, bool) {
	// 1. Look for the header sequence by shifting out dead bytes
	for Serial.Buffered() >= 8 {
		b1, _ := Serial.ReadByte()
		print(b1)
		if b1 != 0x55 {
			println("[CAR_ERR] Bad packet b1 (Not 0x55)")
			continue // Skip this junk byte and check the next one
		}
		// Look ahead at the next byte without removing it yet
		b2, _ := Serial.ReadByte()
		print(b2)
		if b2 != 0xAA {
			println("[CAR_ERR] Bad packet b2 (Not 0xAA)")
			// Not a match! But b2 might be 0x55, so check again next loop
			continue
		}

		// 2. Read the fixed 8 bytes into a temporary buffer to preserve them for logging
		packetBuf := make([]byte, 8)
		packetBuf[0] = b1
		packetBuf[1] = b2
		for i := 2; i < 8; i++ {
			packetBuf[i], _ = Serial.ReadByte()
		}

		// 3. IMMEDIATELY LOG THE RAW HEX (Crucial for debugging/research)
		// Output format example: [UART_RAW] 55 aa 03 02 00 00 01 fa
		println("[UART_RAW]", hex.EncodeToString(packetBuf))

		// 4. Map variables to our safe buffer fields
		length := packetBuf[2]
		mode := packetBuf[3]
		cmdH := packetBuf[4]
		cmdL := packetBuf[5]
		data := packetBuf[6]
		checksum := packetBuf[7]

		// 6. Simple Checksum Validation
		calcChecksum := byte(0x100 - (int(length)+int(mode)+int(cmdH)+int(cmdL)+int(data))&0xFF)
		if calcChecksum != checksum {
			println("[CAR_ERR] Bad packet checksum match dropped. Expected:", calcChecksum, "Got:", checksum)
			return 0, false
		}

		// 7. Verify it is a true Mode 2 button event
		if mode == 0x02 && cmdH == 0x00 && cmdL == 0x00 {
			return data, true
		}
	}

	// println("[CAR_INF] Valid iAP packet received, but not a Mode 2 button event")
	return 0, false
}
