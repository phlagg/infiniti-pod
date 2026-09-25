package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/iap/lingo"
)

const (
	GeneralMajorVersionNumber = 0x01
	GeneralMinorVersionNumber = 0x11

	ExtendedMajorVersionNumber = 0x01
	ExtendedMinorVersionNumber = 0x14

	DisplayRemoteMajorVersionNumber = 0x01
	DisplayRemoteMinorVersionNumber = 0x10

	DigitalAudioMajorVersionNumber = 0x01
	DigitalAudioMinorVersionNumber = 0x12
)

// General
func GeneralAck(ackID byte, cmdIDAckd byte, cmdResultStatus byte) *iap.Command {
	return &iap.Command{
		CmdID:   uint16(ackID),
		CmdData: []byte{cmdResultStatus, cmdIDAckd},
	}
}

func IdentifyDevice() *iap.Command {
	return nil
}

func GetiPodName() *iap.Command {
	return &iap.Command{

		CmdData: []byte("Michael's Phone"),
	}
}
func GetiPodSoftwareVersion() *iap.Command { return nil }
func GetiPodSerialNum() *iap.Command       { return nil }
func GetiPodModelNum() *iap.Command        { return nil }

func GetLingoProtocolVersion(Lingo byte) *iap.Command {
	switch Lingo {
	case lingo.LingoGeneralID:
		return &iap.Command{
			Lingo:   Lingo,
			CmdID:   lingo.GeneralReturnLingoProtocolVersion,
			CmdData: []byte{GeneralMajorVersionNumber, GeneralMinorVersionNumber},
		}
	case lingo.LingoDisplayRemoteID:
		return &iap.Command{
			Lingo:   Lingo,
			CmdID:   lingo.ExtIfaceReturnProtocolVersion,
			CmdData: []byte{DisplayRemoteMajorVersionNumber, DisplayRemoteMinorVersionNumber},
		}
	case lingo.LingoExtendedID:
		return &iap.Command{
			Lingo:   Lingo,
			CmdID:   lingo.ExtIfaceReturnProtocolVersion,
			CmdData: []byte{ExtendedMajorVersionNumber, ExtendedMinorVersionNumber},
		}
	case lingo.LingoDigitalAudioID:
		return &iap.Command{
			Lingo:   Lingo,
			CmdID:   lingo.ExtIfaceReturnProtocolVersion,
			CmdData: []byte{DigitalAudioMajorVersionNumber, DigitalAudioMinorVersionNumber},
		}
	default:
		return &iap.Command{
			Lingo:   Lingo,
			CmdID:   lingo.ExtIfaceReturnProtocolVersion,
			CmdData: []byte{0, 0},
		}
	}
}

func IdentifyDeviceLingoes() {}

func GetAccessoryInfo() *iap.Command { return nil }
func GetiPodOptions() *iap.Command   { return nil }

func GetiPodPreferences() *iap.Command { return nil }
func SetiPodPreferences(data []byte)   {}

// Extended
func ExtendedAck(ackID uint16, cmdIDAckd uint16, cmdResultStatus byte) *iap.Command {
	return &iap.Command{
		CmdID:   ackID,
		CmdData: []byte{cmdResultStatus, byte(cmdIDAckd >> 8), byte(cmdIDAckd & 0xFF)},
	}
}

func GetProtocolVersion() *iap.Command {
	return &iap.Command{
		CmdData: []byte{ExtendedMajorVersionNumber, ExtendedMinorVersionNumber},
	}
}

func GetArtworkFormats() *iap.Command              { return nil }
func GetTrackArtworkData(data []byte) *iap.Command { return nil }

func SetDisplayImage(data []byte) *iap.Command { return nil }
func GetMonoDisplayImageLimits() *iap.Command  { return nil }
func GetColorDisplayImageLimits() *iap.Command { return nil }
