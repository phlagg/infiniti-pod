package iap

import "github.com/phlagg/infiniti-pod/iap/lingo"

// Command represents an incoming iAP command extracted from a packet.
type Command struct {
	Lingo   byte   // 0x00 general, 0x04 extended, etc.
	CmdID   uint16 // command ID (e.g., 0x001C GetPlayStatus)
	CmdData []byte // raw payload bytes after the command ID
}

func parsePacket(pkt *RecievedPacket) Command {
	switch pkt.LingoID {
	case lingo.LingoGeneralID:
		return Command{
			Lingo:   pkt.LingoID,
			CmdID:   extractSingleByteCmdID(pkt.Payload),
			CmdData: pkt.Payload[1:],
		}
	case lingo.LingoExtendedID:
		return Command{
			Lingo:   pkt.LingoID,
			CmdID:   extractTwoByteCmdID(pkt.Payload),
			CmdData: pkt.Payload[2:],
		}
	default:
		return Command{}
	}
}

func extractSingleByteCmdID(payload []byte) uint16 {
	return uint16(payload[0])
}

func extractTwoByteCmdID(payload []byte) uint16 {
	return uint16(payload[0])<<8 | uint16(payload[1])
}
