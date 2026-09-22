package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/iap/lingo"
)

var GeneralTable = []commandEntry{
	// Identify
	{
		CmdID:  lingo.GeneralRequestIdentify,
		RespID: lingo.GeneralIdentify,
		Exec: func(cmd iap.Command) []byte {
			return GeneralAck(lingo.GeneralIdentify, lingo.AckOK)
		}},
	{
		CmdID:  lingo.GeneralIdentify,
		RespID: 0,
		Exec:   nil,
	},

	// ACK
	{
		CmdID:  lingo.GeneralACK,
		RespID: 0,
		Exec:   nil,
	},

	// Remote UI Mode
	{
		CmdID:  lingo.GeneralRequestRemoteUIMode,
		RespID: lingo.GeneralReturnRemoteUIMode,
		Exec: func(cmd iap.Command) []byte {
			return GetRemoteUIMode()
		}},
	{
		CmdID:  lingo.GeneralEnterRemoteUIMode,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			EnterRemoteUIMode()
			return nil
		}},
	{
		CmdID:  lingo.GeneralExitRemoteUIMode,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			ExitRemoteUIMode()
			return nil
		}},

	// iPod Name
	{
		CmdID:  lingo.GeneralRequestiPodName,
		RespID: lingo.GeneralReturniPodName,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodName()
		}},

	// Software Version
	{
		CmdID:  lingo.GeneralRequestiPodSoftwareVersion,
		RespID: lingo.GeneralReturniPodSoftwareVersion,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodSoftwareVersion()
		}},

	// Serial Number
	{
		CmdID:  lingo.GeneralRequestiPodSerialNum,
		RespID: lingo.GeneralReturniPodSerialNum,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodSerialNum()
		}},

	// Model Number
	{
		CmdID:  lingo.GeneralRequestiPodModelNum,
		RespID: lingo.GeneralReturniPodModelNum,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodModelNum()
		}},

	// Lingo Protocol Version
	{
		CmdID:  lingo.GeneralRequestLingoProtocolVersion,
		RespID: lingo.GeneralReturnLingoProtocolVersion,
		Exec: func(cmd iap.Command) []byte {
			return GetLingoProtocolVersion()
		}},

	// Device Lingoes
	{
		CmdID:  lingo.GeneralIdentifyDeviceLingoes,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			IdentifyDeviceLingoes()
			return nil
		}},

	// Device Authentication
	{
		CmdID:  lingo.GeneralGetDevAuthenticationInfo,
		RespID: lingo.GeneralRetDevAuthenticationInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetDevAuthenticationInfo()
		}},
	{
		CmdID:  lingo.GeneralAckDevAuthenticationInfo,
		RespID: 0,
		Exec:   nil,
	},
	{
		CmdID:  lingo.GeneralGetDevAuthenticationSignature,
		RespID: lingo.GeneralRetDevAuthenticationSignature,
		Exec: func(cmd iap.Command) []byte {
			return GetDevAuthenticationSignature()
		}},
	{
		CmdID:  lingo.GeneralAckDevAuthenticationStatus,
		RespID: 0,
		Exec:   nil,
	},

	// iPod Authentication
	{
		CmdID:  lingo.GeneralGetiPodAuthenticationInfo,
		RespID: lingo.GeneralRetiPodAuthenticationInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodAuthenticationInfo()
		}},
	{
		CmdID:  lingo.GeneralAckiPodAuthenticationInfo,
		RespID: 0,
		Exec:   nil,
	},
	{
		CmdID:  lingo.GeneralGetiPodAuthenticationSignature,
		RespID: lingo.GeneralRetiPodAuthenticationSignature,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodAuthenticationSignature()
		}},
	{
		CmdID:  lingo.GeneralAckiPodAuthenticationStatus,
		RespID: 0,
		Exec:   nil,
	},

	// State Change Notification
	{
		CmdID:  lingo.GeneralNotifyiPodStateChange,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return HandleStateChange(cmd.CmdData)
		}},

	// iPod Options
	{
		CmdID:  lingo.GeneralGetiPodOptions,
		RespID: lingo.GeneralRetiPodOptions,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodOptions()
		}},

	// Accessory Info
	{
		CmdID:  lingo.GeneralGetAccessoryInfo,
		RespID: lingo.GeneralRetAccessoryInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetAccessoryInfo()
		}},

	// Preferences
	{
		CmdID:  lingo.GeneralGetiPodPreferences,
		RespID: lingo.GeneralRetiPodPreferences,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodPreferences()
		}},
	{
		CmdID:  lingo.GeneralSetiPodPreferences,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SetiPodPreferences(cmd.CmdData)
			return nil
		}},

	// IDPS
	{
		CmdID:  lingo.GeneralStartIDPS,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			StartIDPS()
			return nil
		}},
	{
		CmdID:  lingo.GeneralSetFIDTokenValues,
		RespID: lingo.GeneralRetFIDTokenValueACKs,
		Exec: func(cmd iap.Command) []byte {
			return SetFIDTokenValues(cmd.CmdData)
		}},
	{
		CmdID:  lingo.GeneralEndIDPS,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			EndIDPS()
			return nil
		}},
	{
		CmdID:  lingo.GeneralIDPSStatus,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return GetIDPSStatus()
		}},

	// Data Session
	{
		CmdID:  lingo.GeneralOpenDataSessionForProtocol,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return OpenDataSession(cmd.CmdData)
		}},
	{
		CmdID:  lingo.GeneralCloseDataSession,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return CloseDataSession(cmd.CmdData)
		}},
	{
		CmdID:  lingo.GeneralDevACK,
		RespID: 0,
		Exec:   nil,
	},
	{
		CmdID:  lingo.GeneralDevDataTransfer,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return HandleDevDataTransfer(cmd.CmdData)
		}},
	{
		CmdID:  lingo.GeneraliPodDataTransfer,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return HandleiPodDataTransfer(cmd.CmdData)
		}},

	// Event Notifications
	{
		CmdID:  lingo.GeneralSetEventNotification,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return SetEventNotification(cmd.CmdData)
		}},
	{
		CmdID:  lingo.GeneraliPodNotification,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return HandleiPodNotification(cmd.CmdData)
		}},
	{
		CmdID:  lingo.GeneralGetiPodOptionsForLingo,
		RespID: lingo.GeneralRetiPodOptionsForLingo,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodOptionsForLingo(cmd.CmdData)
		}},
	{
		CmdID:  lingo.GeneralGetEventNotification,
		RespID: lingo.GeneralRetEventNotification,
		Exec: func(cmd iap.Command) []byte {
			return GetEventNotification()
		}},
	{
		CmdID:  lingo.GeneralGetSupportedEventNotification,
		RespID: lingo.GeneralRetSupportedEventNotification,
		Exec: func(cmd iap.Command) []byte {
			return GetSupportedEventNotification()
		}},
}
