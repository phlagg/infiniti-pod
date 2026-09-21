package extended

import "github.com/phlagg/infiniti-pod/iap"

func HandleExtended(payload []byte) {
	var err error = nil
	var commandID uint16 = (uint16(payload[0]) << 8) | uint16(payload[1])
	switch commandID {

	case ExtIfaceRequestProtocolVersion:
		err = handleRequestProtocolVersion()
		break
	case ExtIfaceRequestiPodName:
		err = handleRequestiPodName()
		break
	case ExtIfaceGetPlayStatus:
		err = handleGetPlayStatus()
		break
	case ExtIfaceSetPlayStatusChangeNotification:
		err = handleSetPlayStatusChangeNotification()
	default:
		err = iap.ErrInvalidCmd

	}
	if err != nil {
		println(err)
	}
}
