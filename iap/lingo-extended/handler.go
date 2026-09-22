package extended

// func Dispatch(payload []byte) (uint16, []byte, error) {
// 	if len(payload) < 2 {
// 		println(ErrPacketTooShort.Error())
// 		return 0, nil, ErrPacketTooShort
// 	}

// 	cmdID := uint16(payload[0])<<8 | uint16(payload[1])
// 	cmdData := payload[2:]

// 	switch cmdID {
// 	case ExtIfaceRequestProtocolVersion:
// 		return handleRequestProtocolVersion()

// 	case ExtIfaceRequestiPodName:
// 		return handleRequestiPodName()

// 	case ExtIfaceGetPlayStatus:
// 		return handleGetPlayStatus()

// 	case ExtIfaceSetPlayStatusChangeNotification:
// 		return handleSetPlayStatusChangeNotification(cmdData)

// 	default:
// 		return 0, nil, ErrInvalidCmd
// 	}
// }

// pg 528
// func handleRequestProtocolVersion() (uint16, []byte, error) {
// 	return ExtIfaceReturnProtocolVersion, []byte{ipod.MajorVersionNumber, ipod.MinorVersionNumber}, nil
// }
// func handleRequestiPodName() (uint16, []byte, error) {
// 	iPodName := "Michael's Phone"
// 	return ExtIfaceReturniPodName, []byte(iPodName), nil
// }

// func handleGetPlayStatus() (uint16, []byte, error) {
// 	return ExtIfaceReturnPlayStatus, []byte{0x00}, nil
// }

// func handleSetPlayStatusChangeNotification(statusChangeEvent []byte) (uint16, []byte, error) {
// 	return ExtIfacePlayStatusChangeNotification, statusChangeEvent, nil
// }

// func handleGetChapterInfo() (uint16, []byte, error) {
// 	// packet := protocol.BuildExtendedPacket(ExtIfaceReturnChapterInfo, nil)
// 	// err := protocol.SendPacket(packet)
// 	// if err != nil {
// 	// 	return err
// 	// }
// 	return 0, []byte{0x00}, nil
// }
