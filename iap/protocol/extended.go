package protocol

import "github.com/phlagg/infiniti-pod/iap/playback"

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

// Extended Interface command summary on pg. 368 of specification

const (
	ExtIfaceReserved0000                          = 0x0000 // Reserved
	ExtIfaceACK                                   = 0x0001 // ACK
	ExtIfaceGetCurrentPlayingChapterInfo          = 0x0002
	ExtIfaceReturnCurrentPlayingChapterInfo       = 0x0003
	ExtIfaceSetCurrentPlayingChapter              = 0x0004
	ExtIfaceGetCurrentPlayingChapterPlayStatus    = 0x0005
	ExtIfaceReturnCurrentPlayingChapterPlayStatus = 0x0006
	ExtIfaceGetCurrentPlayingChapterName          = 0x0007
	ExtIfaceReturnCurrentPlayingChapterName       = 0x0008

	ExtIfaceGetAudiobookSpeed    = 0x0009
	ExtIfaceReturnAudiobookSpeed = 0x000A
	ExtIfaceSetAudiobookSpeed    = 0x000B

	ExtIfaceGetIndexedPlayingTrackInfo    = 0x000C
	ExtIfaceReturnIndexedPlayingTrackInfo = 0x000D

	ExtIfaceGetArtworkFormats    = 0x000E
	ExtIfaceReturnArtworkFormats = 0x000F

	ExtIfaceGetTrackArtworkData    = 0x0010
	ExtIfaceReturnTrackArtworkData = 0x0011

	ExtIfaceRequestProtocolVersion = 0x0012 // Deprecated
	ExtIfaceReturnProtocolVersion  = 0x0013 // Deprecated
	ExtIfaceRequestiPodName        = 0x0014 // Deprecated
	ExtIfaceReturniPodName         = 0x0015 // Deprecated

	ExtIfaceResetDBSelection                 = 0x0016
	ExtIfaceSelectDBRecord                   = 0x0017
	ExtIfaceGetNumberCategorizedDBRecords    = 0x0018
	ExtIfaceReturnNumberCategorizedDBRecords = 0x0019
	ExtIfaceRetrieveCategorizedDBRecords     = 0x001A
	ExtIfaceReturnCategorizedDBRecord        = 0x001B

	ExtIfaceGetPlayStatus    = 0x001C
	ExtIfaceReturnPlayStatus = 0x001D

	ExtIfaceGetCurrentPlayingTrackIndex    = 0x001E
	ExtIfaceReturnCurrentPlayingTrackIndex = 0x001F

	ExtIfaceGetIndexedPlayingTrackTitle    = 0x0020
	ExtIfaceReturnIndexedPlayingTrackTitle = 0x0021

	ExtIfaceGetIndexedPlayingTrackArtistName    = 0x0022
	ExtIfaceReturnIndexedPlayingTrackArtistName = 0x0023

	ExtIfaceGetIndexedPlayingTrackAlbumName    = 0x0024
	ExtIfaceReturnIndexedPlayingTrackAlbumName = 0x0025

	ExtIfaceSetPlayStatusChangeNotification = 0x0026
	ExtIfacePlayStatusChangeNotification    = 0x0027

	ExtIfacePlayCurrentSelection = 0x0028
	ExtIfacePlayControl          = 0x0029

	ExtIfaceGetTrackArtworkTimes    = 0x002A
	ExtIfaceReturnTrackArtworkTimes = 0x002B

	ExtIfaceGetShuffle    = 0x002C
	ExtIfaceReturnShuffle = 0x002D
	ExtIfaceSetShuffle    = 0x002E

	ExtIfaceGetRepeat    = 0x002F
	ExtIfaceReturnRepeat = 0x0030
	ExtIfaceSetRepeat    = 0x0031

	ExtIfaceSetDisplayImage              = 0x0032
	ExtIfaceGetMonoDisplayImageLimits    = 0x0033
	ExtIfaceReturnMonoDisplayImageLimits = 0x0034

	ExtIfaceGetNumPlayingTracks    = 0x0035
	ExtIfaceReturnNumPlayingTracks = 0x0036

	ExtIfaceSetCurrentPlayingTrack = 0x0037
	ExtIfaceSelectSortDBRecord     = 0x0038

	ExtIfaceGetColorDisplayImageLimits    = 0x0039
	ExtIfaceReturnColorDisplayImageLimits = 0x003A

	ExtIfaceResetDBSelectionHierarchy = 0x003B
	ExtIfaceGetDBiTunesInfo           = 0x003C
	ExtIfaceReturnDBiTunesInfo        = 0x003D

	ExtIfaceGetUIDTrackInfo    = 0x003E
	ExtIfaceReturnUIDTrackInfo = 0x003F

	ExtIfaceGetDBTrackInfo    = 0x0040
	ExtIfaceReturnDBTrackInfo = 0x0041

	ExtIfaceGetPBTrackInfo    = 0x0042
	ExtIfaceReturnPBTrackInfo = 0x0043

	ExtIfaceReserved0044 = 0x0044 // Reserved

)

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
		err = ErrInvalidCmd

	}
	if err != nil {
		println(err)
	}
}

func SendExtendedACK(cmdIDAckd uint16, cmdResultStatus byte) error {
	err := buildAndSendExtendedPacket(ExtIfaceACK, []byte{cmdResultStatus, byte(cmdIDAckd >> 8), byte(cmdIDAckd & 0xFF)})
	if err != nil {
		return err
	}
	return nil
}

func buildAndSendExtendedPacket(cmd uint16, cmdData []byte) error {
	packet := BuildSmallExtendedPacket(cmd, cmdData)
	err := SendPacket(packet)
	if err != nil {
		return err
	}
	return nil
}

// pg 528
func handleRequestProtocolVersion() error {
	packet := BuildSmallExtendedPacket(ExtIfaceReturnProtocolVersion, []byte{MajorVersionNumber, MinorVersionNumber})
	err := SendPacket(packet)
	if err != nil {
		return err
	}
	return nil
}

func handleRequestiPodName() error {
	iPodName := "Michael's Phone"
	packet := BuildSmallExtendedPacket(ExtIfaceReturniPodName, []byte(iPodName))
	err := SendPacket(packet)
	if err != nil {
		return err
	}
	return nil
}
func handleGetPlayStatus() error {

	playStatus := playback.GetPlayStatus()

	packet := BuildSmallExtendedPacket(ExtIfaceReturnPlayStatus, playStatus)
	err := SendPacket(packet)
	if err != nil {
		return err
	}
	return nil
}

func handleSetPlayStatusChangeNotification() error {

	err := buildAndSendExtendedPacket(ExtIfaceSetPlayStatusChangeNotification, []byte{0x01})
	if err != nil {
		return err
	}
	return nil
}

func handleGetChapterInfo() error {
	// packet := protocol.BuildExtendedPacket(ExtIfaceReturnChapterInfo, nil)
	// err := protocol.SendPacket(packet)
	// if err != nil {
	// 	return err
	// }
	return nil
}
