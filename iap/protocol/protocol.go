package protocol

const (
	MajorVersionNumber = 0x01
	MinorVersionNumber = 0x14
)
const MaxPayloadSize = 255

const (
	SyncByte        byte = 0xFF
	PacketStartByte byte = 0x55
)

type Packet struct {
	LingoID byte
	Payload []byte
}

func Dispatch(pkt *Packet) {
	switch pkt.LingoID {
	case GeneralLingoID:
		HandleGeneral(pkt.Payload)
	case ExtendedInterfaceID:
		HandleExtended(pkt.Payload)
	}
}
