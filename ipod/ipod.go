package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	extended "github.com/phlagg/infiniti-pod/iap/lingo-extended"
	general "github.com/phlagg/infiniti-pod/iap/lingo-general"
	"github.com/phlagg/infiniti-pod/iap/transport"
)

type iPodState uint8

const (
	Initialize = iota
	Identify
	Authenticate
)

var nextState iPodState = Initialize
var currentState iPodState = Initialize

func Init() error {
	transport.Initialize()
	return nil
}

func Run() {
	if currentState != nextState {
		currentState = nextState
	}
	switch currentState {
	case Initialize:

	case Identify:
	case Authenticate:
	}
	processFrames()

}

func processFrames() {
	pkt, ok := iap.ReadPacket()
	if !ok {
		return
	}
	return
}

func handlePacket(pkt *iap.Packet) {
	switch pkt.LingoID {
	case iap.LingoGeneralID:
		general.HandleGeneral(pkt.Payload)
	case iap.LingoExtendedID:
		extended.HandleExtended(pkt.Payload)
	}
}
