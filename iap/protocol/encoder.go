package protocol

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
	packetBuf[pktPayloadLen] = CalculateBufferChecksum(packetBuf)
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
	packetBuf[len(cmdData)+5] = CalculateBufferChecksum(packetBuf)
	return packetBuf
}
