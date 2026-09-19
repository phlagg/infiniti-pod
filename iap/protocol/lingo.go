package protocol

type iAPError string

func (e iAPError) Error() string { return string(e) }

const (
	ErrPacketTooShort iAPError = "packet too short"
	ErrBadChecksum    iAPError = "bad checksum"
	ErrInvalidStart   iAPError = "invalid start byte"
	ErrInvalidCmd     iAPError = "invalid command"
)

const (
	GeneralID           = 0x00
	ExtendedInterfaceID = 0x04
)

type GeneralCommand byte

const (
	GeneralRequestIdentify     GeneralCommand = 0x00
	GeneralIdentify            GeneralCommand = 0x01
	GeneralACK                 GeneralCommand = 0x02
	GeneralRequestRemoteUIMode GeneralCommand = 0x03
	GeneralReturnRemoteUIMode  GeneralCommand = 0x04
	GeneralEnterRemoteUIMode   GeneralCommand = 0x05
	GeneralExitRemoteUIMode    GeneralCommand = 0x06
	GeneralRequestiPodName     GeneralCommand = 0x07
	GeneralReturniPodName      GeneralCommand = 0x08

	GeneralRequestiPodSoftwareVersion  GeneralCommand = 0x09
	GeneralReturniPodSoftwareVersion   GeneralCommand = 0x0A
	GeneralRequestiPodSerialNum        GeneralCommand = 0x0B
	GeneralReturniPodSerialNum         GeneralCommand = 0x0C
	GeneralRequestiPodModelNum         GeneralCommand = 0x0D
	GeneralReturniPodModelNum          GeneralCommand = 0x0E
	GeneralRequestLingoProtocolVersion GeneralCommand = 0x0F
	GeneralReturnLingoProtocolVersion  GeneralCommand = 0x10

	GeneralIdentifyDeviceLingoes         GeneralCommand = 0x13
	GeneralGetDevAuthenticationInfo      GeneralCommand = 0x14
	GeneralRetDevAuthenticationInfo      GeneralCommand = 0x15
	GeneralAckDevAuthenticationInfo      GeneralCommand = 0x16
	GeneralGetDevAuthenticationSignature GeneralCommand = 0x17
	GeneralRetDevAuthenticationSignature GeneralCommand = 0x18
	GeneralAckDevAuthenticationStatus    GeneralCommand = 0x19

	GeneralGetiPodAuthenticationInfo      GeneralCommand = 0x1A
	GeneralRetiPodAuthenticationInfo      GeneralCommand = 0x1B
	GeneralAckiPodAuthenticationInfo      GeneralCommand = 0x1C
	GeneralGetiPodAuthenticationSignature GeneralCommand = 0x1D
	GeneralRetiPodAuthenticationSignature GeneralCommand = 0x1E
	GeneralAckiPodAuthenticationStatus    GeneralCommand = 0x1F

	GeneralNotifyiPodStateChange GeneralCommand = 0x23
	GeneralGetiPodOptions        GeneralCommand = 0x24
	GeneralRetiPodOptions        GeneralCommand = 0x25

	GeneralGetAccessoryInfo GeneralCommand = 0x27
	GeneralRetAccessoryInfo GeneralCommand = 0x28

	GeneralGetiPodPreferences GeneralCommand = 0x29
	GeneralRetiPodPreferences GeneralCommand = 0x2A
	GeneralSetiPodPreferences GeneralCommand = 0x2B

	GeneralStartIDPS            GeneralCommand = 0x38
	GeneralSetFIDTokenValues    GeneralCommand = 0x39
	GeneralRetFIDTokenValueACKs GeneralCommand = 0x3A
	GeneralEndIDPS              GeneralCommand = 0x3B
	GeneralIDPSStatus           GeneralCommand = 0x3C

	GeneralOpenDataSessionForProtocol GeneralCommand = 0x3F
	GeneralCloseDataSession           GeneralCommand = 0x40
	GeneralDevACK                     GeneralCommand = 0x41
	GeneralDevDataTransfer            GeneralCommand = 0x42
	GeneraliPodDataTransfer           GeneralCommand = 0x43

	GeneralSetEventNotification          GeneralCommand = 0x49
	GeneraliPodNotification              GeneralCommand = 0x4A
	GeneralGetiPodOptionsForLingo        GeneralCommand = 0x4B
	GeneralRetiPodOptionsForLingo        GeneralCommand = 0x4C
	GeneralGetEventNotification          GeneralCommand = 0x4D
	GeneralRetEventNotification          GeneralCommand = 0x4E
	GeneralGetSupportedEventNotification GeneralCommand = 0x4F
	GeneralRetSupportedEventNotification GeneralCommand = 0x51
)

// General Lingo ACK status values (Table 2-13)
type AckStatus byte

const (
	AckOK                          AckStatus = 0x00
	AckErrorUnknownCategory        AckStatus = 0x01
	AckErrorCommandFailed          AckStatus = 0x02
	AckErrorOutOfResources         AckStatus = 0x03
	AckErrorBadParameter           AckStatus = 0x04
	AckErrorUnknownID              AckStatus = 0x05
	AckCommandPending              AckStatus = 0x06
	AckErrorNotAuthenticated       AckStatus = 0x07
	AckErrorBadAuthVersion         AckStatus = 0x08
	AckErrorPowerModeFailed        AckStatus = 0x09
	AckErrorCertificateInvalid     AckStatus = 0x0A
	AckErrorCertificatePermissions AckStatus = 0x0B
	AckErrorFileInUse              AckStatus = 0x0C
	AckErrorInvalidFileHandle      AckStatus = 0x0D
	AckErrorDirectoryNotEmpty      AckStatus = 0x0E
	AckErrorTimedOut               AckStatus = 0x0F
	AckErrorUnavailableInMode      AckStatus = 0x10
	AckErrorInvalidResistorID      AckStatus = 0x11
	// 0x12–0x14 reserved
	AckErrorMaxAccessoryConnections AckStatus = 0x15
	// 0x16–0xFF reserved
)

// pg. 523
func handleIdentifyRequest(payload []byte) error {
	packet := BuildPacket(0x00, 0x00, payload)
	err := SendPacket(packet)
	if err != nil {
		return err
	}
	return nil
}

func HandleGeneral(commandID byte, payload []byte) {
	println("0x00", commandID, len(payload), payload[0])
}

func HandleExtended(commandID byte, payload []byte) {
	var err error = nil
	println("0x04", commandID, len(payload), payload[0])
	switch commandID {

	case ExtIfaceRequestProtocolVersion:
		err = handleRequestProtocolVersion()
		break

	default:
		err = ErrInvalidCmd

	}
	if err != nil {
		println(err)
	}

}
