package ipod

import "github.com/phlagg/infiniti-pod/iap"

type iPodState struct {
	State uint8
	Name  string
}

const (
	StateInitialize uint8 = iota
	StateIdentify
	StateAuthenticate
	StateConnected
)

var (
	Initialize   = iPodState{State: StateInitialize, Name: "Initialize"}
	Identify     = iPodState{State: StateIdentify, Name: "Identify"}
	Authenticate = iPodState{State: StateAuthenticate, Name: "Authenticate"}
	Connected    = iPodState{State: StateConnected, Name: "Connected"}
)

var nextState iPodState = Initialize
var currentState iPodState

type Response struct {
	ID      uint16
	Payload []byte
}

func Run() {
	if currentState != nextState {
		currentState = nextState
		logiPodStatus(currentState.Name)
	}
	switch currentState {
	case Initialize:
		nextState = Identify

	case Identify:
		nextState = Authenticate

	case Authenticate:
		nextState = Connected

	case Connected:
		iap.ProcessFrames(handleCommand)

	default:
		nextState = Initialize
	}

}

func logiPodStatus(msg string) {
	println("[iPod] ", msg)
}
