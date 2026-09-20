package protocol

import "encoding/hex"

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
	println("[ENCODER]", hex.EncodeToString(*packet))
}

func CalculateBufferChecksum(buffer []byte) byte {
	var byteSum byte
	for i := 0; i < len(buffer); i++ {
		byteSum += (buffer)[i]
	}
	return -byteSum
}
