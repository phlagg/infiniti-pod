package protocol

const (
	MajorVersionNumber = 1
	MinorVersionNumber = 0
)
const MaxPayloadSize = 255

const (
	SyncByte        byte = 0xFF
	PacketStartByte byte = 0x55
)

type Packet struct {
	LingoID   byte
	CommandID byte
	Payload   []byte
}

func Dispatch(pkt *Packet) {
	switch pkt.LingoID {
	case GeneralID:
		HandleGeneral(pkt.CommandID, pkt.Payload)
	case ExtendedInterfaceID:
		HandleExtended(pkt.CommandID, pkt.Payload)
	}
}
