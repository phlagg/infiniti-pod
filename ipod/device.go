package ipod

import (
	"github.com/phlagg/infiniti-pod/iap/lingo"
	"github.com/phlagg/infiniti-pod/transport/ble"
)

const (
	MajorVersionNumber = 0x01
	MinorVersionNumber = 0x14
)

// handleCarSignals bridges the vehicle iAP commands over to Bluetooth media keys
func handleCarSignals(code byte) {
	switch code {
	case 0x01:
		println("[BRIDGE] Vehicle command: NEXT -> Notifying Phone")
		ble.PressMediaKey(ble.KeyNext)
	case 0x08:
		println("[BRIDGE] Vehicle command: PREVIOUS -> Notifying Phone")
		ble.PressMediaKey(ble.KeyPrevious)
	case 0x02:
		println("[BRIDGE] Vehicle command: PLAY/PAUSE -> Notifying Phone")
		ble.PressMediaKey(ble.KeyPlayPause)
	case 0x00:
		// Button released event, ignore safely
	}
}

// General
func GeneralAck(cmdIDAckd byte, cmdResultStatus byte) []byte {
	return []byte{lingo.GeneralACK, cmdResultStatus, cmdIDAckd}
}

func IdentifyDevice() []byte {
	return nil
}

func GetiPodName() []byte {
	return []byte("Michael's Phone")
}
func GetiPodSoftwareVersion() []byte { return nil }
func GetiPodSerialNum() []byte       { return nil }
func GetiPodModelNum() []byte        { return nil }

func GetLingoProtocolVersion() []byte { return nil }

func IdentifyDeviceLingoes() {}

func GetAccessoryInfo() []byte { return nil }
func GetiPodOptions() []byte   { return nil }

func GetiPodPreferences() []byte     { return nil }
func SetiPodPreferences(data []byte) {}

// Extended
func GetProtocolVersion() []byte {
	return []byte{MajorVersionNumber, MinorVersionNumber}
}

func GetArtworkFormats() []byte              { return nil }
func GetTrackArtworkData(data []byte) []byte { return nil }

func SetDisplayImage(data []byte) []byte { return nil }
func GetMonoDisplayImageLimits() []byte  { return nil }
func GetColorDisplayImageLimits() []byte { return nil }
