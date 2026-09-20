package protocol

const (
	GeneralRequestIdentify     = 0x00
	GeneralIdentify            = 0x01
	GeneralACK                 = 0x02
	GeneralRequestRemoteUIMode = 0x03
	GeneralReturnRemoteUIMode  = 0x04
	GeneralEnterRemoteUIMode   = 0x05
	GeneralExitRemoteUIMode    = 0x06
	GeneralRequestiPodName     = 0x07
	GeneralReturniPodName      = 0x08

	GeneralRequestiPodSoftwareVersion  = 0x09
	GeneralReturniPodSoftwareVersion   = 0x0A
	GeneralRequestiPodSerialNum        = 0x0B
	GeneralReturniPodSerialNum         = 0x0C
	GeneralRequestiPodModelNum         = 0x0D
	GeneralReturniPodModelNum          = 0x0E
	GeneralRequestLingoProtocolVersion = 0x0F
	GeneralReturnLingoProtocolVersion  = 0x10

	GeneralIdentifyDeviceLingoes         = 0x13
	GeneralGetDevAuthenticationInfo      = 0x14
	GeneralRetDevAuthenticationInfo      = 0x15
	GeneralAckDevAuthenticationInfo      = 0x16
	GeneralGetDevAuthenticationSignature = 0x17
	GeneralRetDevAuthenticationSignature = 0x18
	GeneralAckDevAuthenticationStatus    = 0x19

	GeneralGetiPodAuthenticationInfo      = 0x1A
	GeneralRetiPodAuthenticationInfo      = 0x1B
	GeneralAckiPodAuthenticationInfo      = 0x1C
	GeneralGetiPodAuthenticationSignature = 0x1D
	GeneralRetiPodAuthenticationSignature = 0x1E
	GeneralAckiPodAuthenticationStatus    = 0x1F

	GeneralNotifyiPodStateChange = 0x23
	GeneralGetiPodOptions        = 0x24
	GeneralRetiPodOptions        = 0x25

	GeneralGetAccessoryInfo = 0x27
	GeneralRetAccessoryInfo = 0x28

	GeneralGetiPodPreferences = 0x29
	GeneralRetiPodPreferences = 0x2A
	GeneralSetiPodPreferences = 0x2B

	GeneralStartIDPS            = 0x38
	GeneralSetFIDTokenValues    = 0x39
	GeneralRetFIDTokenValueACKs = 0x3A
	GeneralEndIDPS              = 0x3B
	GeneralIDPSStatus           = 0x3C

	GeneralOpenDataSessionForProtocol = 0x3F
	GeneralCloseDataSession           = 0x40
	GeneralDevACK                     = 0x41
	GeneralDevDataTransfer            = 0x42
	GeneraliPodDataTransfer           = 0x43

	GeneralSetEventNotification          = 0x49
	GeneraliPodNotification              = 0x4A
	GeneralGetiPodOptionsForLingo        = 0x4B
	GeneralRetiPodOptionsForLingo        = 0x4C
	GeneralGetEventNotification          = 0x4D
	GeneralRetEventNotification          = 0x4E
	GeneralGetSupportedEventNotification = 0x4F
	GeneralRetSupportedEventNotification = 0x51
)

// General Lingo ACK status values (Table 2-13)

const (
	AckOK                          byte = 0x00
	AckErrorUnknownCategory        byte = 0x01
	AckErrorCommandFailed          byte = 0x02
	AckErrorOutOfResources         byte = 0x03
	AckErrorBadParameter           byte = 0x04
	AckErrorUnknownID              byte = 0x05
	AckCommandPending              byte = 0x06
	AckErrorNotAuthenticated       byte = 0x07
	AckErrorBadAuthVersion         byte = 0x08
	AckErrorPowerModeFailed        byte = 0x09
	AckErrorCertificateInvalid     byte = 0x0A
	AckErrorCertificatePermissions byte = 0x0B
	AckErrorFileInUse              byte = 0x0C
	AckErrorInvalidFileHandle      byte = 0x0D
	AckErrorDirectoryNotEmpty      byte = 0x0E
	AckErrorTimedOut               byte = 0x0F
	AckErrorUnavailableInMode      byte = 0x10
	AckErrorInvalidResistorID      byte = 0x11
	// 0x12–0x14 reserved
	AckErrorMaxAccessoryConnections byte = 0x15
	// 0x16–0xFF reserved
)

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
		println("0x00", cmdID, len(cmdData), cmdData[0])
		err = handleGeneralIdentify()
		break

	default:
		err = ErrInvalidCmd

	}
	if err != nil {
		println(err)
	}

}

func SendACK(lingo byte, cmdIDAckd byte, cmdResultStatus byte) error {
	err := buildAndSendSmallPacket(lingo, GeneralACK, []byte{cmdResultStatus, cmdIDAckd})
	if err != nil {
		return err
	}
	return nil
}

// pg. 523
func handleGeneralIdentify() error {
	err := SendACK(GeneralLingoID, GeneralIdentify, AckOK)
	if err != nil {
		return err
	}
	return nil
}

func buildAndSendSmallPacket(lingo byte, cmd byte, cmdData []byte) error {
	packet := BuildSmallPacket(lingo, cmd, cmdData)
	err := SendPacket(packet)
	if err != nil {
		return err
	}
	return nil
}
