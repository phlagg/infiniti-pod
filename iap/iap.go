package iap

import (
	"github.com/phlagg/infiniti-pod/iap/protocol"
	"github.com/phlagg/infiniti-pod/iap/transport"
)

var pkt *protocol.Packet

func Init() error {
	transport.Initialize()
	return nil
}

func ReadLoop() {
	pkt, ok := protocol.ReadPacket()
	if !ok {
		return
	}

	protocol.Dispatch(pkt)
}
