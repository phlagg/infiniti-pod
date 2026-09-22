package lingo

const (
	LingoGeneralID       = 0x00
	LingoSimpleRemoteID  = 0x02
	LingoDisplayRemoteID = 0x03
	LingoExtendedID      = 0x04
	LingoUSBHostID       = 0x06
	LingoRFTunerID       = 0x07
	LingoEqID            = 0x08
	LingoSportsID        = 0x09
	LingoDigitalAudioID  = 0x0A
	LingoStorageID       = 0x0C
)

// Implementing iAP Extended Interface Mode from pg 342 of specification
const (
	ExtendedPacketFlag = 0x00
	AckStartIdps       = 0x02
)

type AccAckStatus uint8

const (
	AccAckOK                    AccAckStatus = 0x00 // Success (OK)
	AccAckReserved01            AccAckStatus = 0x01 // Reserved
	AccAckErrorFailed           AccAckStatus = 0x02 // ERROR: Command failed
	AccAckErrorOutOfResources   AccAckStatus = 0x03 // ERROR: Out of resources
	AccAckErrorBadParameter     AccAckStatus = 0x04 // ERROR: Bad parameter
	AccAckErrorUnknownID        AccAckStatus = 0x05 // ERROR: Unknown ID
	AccAckReserved06            AccAckStatus = 0x06 // Reserved
	AccAckErrorNotAuthenticated AccAckStatus = 0x07 // ERROR: Accessory not authenticated

	// 0x08–0xFF are reserved.
	AccAckReservedStart AccAckStatus = 0x08
)

type AckCommand uint8

const (
	iPodAck AckCommand = 0x01
)

type RetiPodOption byte

const (
	VideoBrowsing        RetiPodOption = 0x00
	RemoteUiEnhancements RetiPodOption = 0x01
	NestedPlaylists      RetiPodOption = 0x02
	Reserved             RetiPodOption = 0x03
	SetDisplayImage      RetiPodOption = 0x04
	ReservedRange        RetiPodOption = 0x05
)
