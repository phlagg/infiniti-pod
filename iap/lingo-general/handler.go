package general

import "github.com/phlagg/infiniti-pod/iap"

type DeviceGeneral interface {
	// UIMode() UIMode
	// SetUIMode(UIMode)
	Name() string
	SoftwareVersion() (major, minor, rev uint8)
	SerialNum() string

	LingoProtocolVersion(lingo uint8) (major, minor uint8)
	LingoOptions(ling uint8) uint64

	PrefSettingID(classID uint8) uint8
	SetPrefSettingID(classID, settingID uint8, restoreOnExit bool)

	// StartIDPS()
	// EndIDPS(status AccEndIDPSStatus)
	// SetToken(token FIDTokenValue) error
	AccAuthCert(cert []byte)

	SetEventNotificationMask(mask uint64)
	EventNotificationMask() uint64
	SupportedEventNotificationMask() uint64

	CancelCommand(lingo uint8, cmd uint16, transaction uint16)

	MaxPayload() uint16
}

func HandleGeneral(payload []byte) {
	var err error = nil

	cmdID := payload[0]
	var cmdData []byte
	if len(payload) > 1 {
		cmdData = payload[1:]
	} else {
		cmdData = []byte{}
	}

	switch cmdID {

	case GeneralIdentify:
		err = handleGeneralIdentify()
		break

	default:
		err = iap.ErrInvalidCmd
		print(cmdData)
		println(err)
	}

}
