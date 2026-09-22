package iap

import "github.com/phlagg/infiniti-pod/iap/lingo"

// Response represents a device-generated response that iap will frame and send.
type Response struct {
	Lingo   byte
	CmdID   uint16 // response command ID (e.g., 0x001D ReturnPlayStatus)
	CmdData []byte // response payload
}

func buildResponse(resp *Response) []byte {
	switch resp.Lingo {
	case lingo.LingoGeneralID:
		return buildSmallPacket(resp)
	case lingo.LingoExtendedID:
		return buildSmallExtendedPacket(resp)

	}
	return []byte{}
}

func buildSmallPacket(resp *Response) []byte {
	pktPayloadLen := 2 + len(resp.CmdData) // cmd + lingo+ cmdDataLength
	buffLen := 2 + pktPayloadLen + 1       // PacketStartByte + pktPayloadLen byte + []pktPayload + Checksum
	packetBuf := make([]byte, buffLen)
	packetBuf[0] = PacketStartByte
	packetBuf[1] = byte(pktPayloadLen)
	packetBuf[2] = resp.Lingo
	packetBuf[3] = byte(resp.CmdID)
	for i := 0; i < len(resp.CmdData); i++ {
		packetBuf[4+i] = resp.CmdData[i]
	}
	packetBuf[buffLen-1] = CalculateBufferChecksum(packetBuf[1:])
	return packetBuf
}

func buildSmallExtendedPacket(resp *Response) []byte {
	pktPayloadLen := 3 + len(resp.CmdData) // 2 cmd bytes + lingo +  cmdDatalength
	buffLen := 2 + pktPayloadLen + 1       // PacketStartByte + pktPayloadLen byte + []pktPayload + Checksum
	packetBuf := make([]byte, buffLen)
	packetBuf[0] = PacketStartByte
	packetBuf[1] = byte(pktPayloadLen)
	packetBuf[2] = resp.Lingo
	packetBuf[3] = byte(resp.CmdID >> 8)
	packetBuf[4] = byte(resp.CmdID & 0xFF)
	for i := 0; i < len(resp.CmdData); i++ {
		packetBuf[5+i] = resp.CmdData[i]
	}
	packetBuf[buffLen-1] = CalculateBufferChecksum(packetBuf[1:])
	// logPacket(&packetBuf)
	return packetBuf
}

func CalculateBufferChecksum(buffer []byte) byte {
	var byteSum byte
	for i := 0; i < len(buffer); i++ {
		byteSum += (buffer)[i]
	}
	return -byteSum
}

// func buildSmallPacket(lingo byte, cmd byte, cmdData []byte) []byte {
// 	pktPayloadLen := 2 + len(cmdData) // cmd + lingo+ cmdDataLength
// 	buffLen := 2 + pktPayloadLen + 1  // PacketStartByte + pktPayloadLen byte + []pktPayload + Checksum
// 	packetBuf := make([]byte, buffLen)
// 	packetBuf[0] = PacketStartByte
// 	packetBuf[1] = byte(pktPayloadLen)
// 	packetBuf[2] = lingo
// 	packetBuf[3] = cmd
// 	for i := 0; i < len(cmdData); i++ {
// 		packetBuf[4+i] = cmdData[i]
// 	}
// 	packetBuf[buffLen-1] = CalculateBufferChecksum(packetBuf[1:])
// 	return packetBuf
// }

// func buildSmallExtendedPacket(cmd uint16, cmdData []byte) []byte {
// 	pktPayloadLen := 3 + len(cmdData) // 2 cmd bytes + lingo +  cmdDatalength
// 	buffLen := 2 + pktPayloadLen + 1  // PacketStartByte + pktPayloadLen byte + []pktPayload + Checksum
// 	packetBuf := make([]byte, buffLen)
// 	packetBuf[0] = PacketStartByte
// 	packetBuf[1] = byte(pktPayloadLen)
// 	packetBuf[2] = 0x04
// 	packetBuf[3] = byte(cmd >> 8)
// 	packetBuf[4] = byte(cmd & 0xFF)
// 	for i := 0; i < len(cmdData); i++ {
// 		packetBuf[5+i] = cmdData[i]
// 	}
// 	packetBuf[buffLen-1] = CalculateBufferChecksum(packetBuf[1:])
// 	// logPacket(&packetBuf)
// 	return packetBuf
// }

// func buildAndSendExtendedPacket(cmd uint16, cmdData []byte) error {
// 	packet := buildSmallExtendedPacket(cmd, cmdData)
// 	err := SendPacket(packet)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func buildAndSendSmallPacket(lingo byte, cmd byte, cmdData []byte) error {
// 	packet := buildSmallPacket(lingo, cmd, cmdData)
// 	if err := SendPacket(packet); err != nil {
// 		return err
// 	}
// 	return nil
// }
