package iap

import (
	"encoding/hex"

	"github.com/phlagg/infiniti-pod/iap/lingo"
)

type ResponsePacket []byte

func buildResponsePacket(c *Command) *ResponsePacket {
	switch c.Lingo {
	case lingo.LingoGeneralID:
		println("response:", hex.EncodeToString(c.CmdData))
		return buildSmallPacket(c)
	case lingo.LingoExtendedID:
		return buildSmallExtendedPacket(c)

	}
	return &ResponsePacket{}
}

func buildSmallPacket(c *Command) *ResponsePacket {
	pktPayloadLen := 2 + len(c.CmdData) // cmd + lingo+ cmdDataLength
	buffLen := 2 + pktPayloadLen + 1    // PacketStartByte + pktPayloadLen byte + []pktPayload + Checksum
	packetBuf := make(ResponsePacket, buffLen)
	packetBuf[0] = PacketStartByte
	packetBuf[1] = byte(pktPayloadLen)
	packetBuf[2] = c.Lingo
	packetBuf[3] = byte(c.CmdID)
	for i := 0; i < len(c.CmdData); i++ {
		packetBuf[4+i] = c.CmdData[i]
	}
	packetBuf[buffLen-1] = CalculateBufferChecksum(packetBuf[1:])
	return &packetBuf
}

func buildSmallExtendedPacket(c *Command) *ResponsePacket {
	pktPayloadLen := 3 + len(c.CmdData) // 2 cmd bytes + lingo +  cmdDatalength
	buffLen := 2 + pktPayloadLen + 1    // PacketStartByte + pktPayloadLen byte + []pktPayload + Checksum
	packetBuf := make(ResponsePacket, buffLen)
	packetBuf[0] = PacketStartByte
	packetBuf[1] = byte(pktPayloadLen)
	packetBuf[2] = c.Lingo
	packetBuf[3] = byte(c.CmdID >> 8)
	packetBuf[4] = byte(c.CmdID & 0xFF)
	for i := 0; i < len(c.CmdData); i++ {
		packetBuf[5+i] = c.CmdData[i]
	}
	packetBuf[buffLen-1] = CalculateBufferChecksum(packetBuf[1:])
	// logPacket(&packetBuf)
	return &packetBuf
}

func CalculateBufferChecksum(buffer []byte) byte {
	var byteSum byte
	for i := 0; i < len(buffer); i++ {
		byteSum += (buffer)[i]
	}
	return -byteSum
}
