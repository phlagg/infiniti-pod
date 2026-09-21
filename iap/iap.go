package iap

const (
	MajorVersionNumber = 0x01
	MinorVersionNumber = 0x14
)

type iAPError string

func (e iAPError) Error() string { return string(e) }

const (
	ErrPacketTooShort iAPError = "packet too short"
	ErrBadChecksum    iAPError = "bad checksum"
	ErrInvalidStart   iAPError = "invalid start byte"
	ErrInvalidCmd     iAPError = "invalid command"
)
