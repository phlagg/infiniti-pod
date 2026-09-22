package iap

import "encoding/hex"

type iAPError string

func (e iAPError) Error() string { return string(e) }

const (
	ErrPacketTooShort iAPError = "packet too short"
	ErrBadChecksum    iAPError = "bad checksum"
	ErrInvalidStart   iAPError = "invalid start byte"
	ErrInvalidCmd     iAPError = "invalid command"
	ErrInvalidLingoID iAPError = "invalid lingo ID"
)

func ProcessFrames(handler func(Command) Response) {
	pkt, ok := ReadPacket()
	if !ok {
		return
	}

	cmd := parseCommand(pkt)
	buf := []byte{byte(cmd.CmdID >> 8), byte(cmd.CmdID & 0xFF)}
	println("[cmd]", hex.EncodeToString(buf))
	println("[cmdData]", hex.EncodeToString(cmd.CmdData))

	resp := handler(cmd)

	SendPacket(buildResponse(&resp))
}
