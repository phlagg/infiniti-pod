package protocol

type iAPError string

func (e iAPError) Error() string { return string(e) }

const (
	ErrPacketTooShort iAPError = "packet too short"
	ErrBadChecksum    iAPError = "bad checksum"
	ErrInvalidStart   iAPError = "invalid start byte"
	ErrInvalidCmd     iAPError = "invalid command"
)

const (
	GeneralLingoID      byte = 0x00
	ExtendedInterfaceID byte = 0x04
)
