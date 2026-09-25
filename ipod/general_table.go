package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/iap/lingo"
)

var GeneralTable = []commandEntry{
	// Identify
	{
		CmdID: lingo.GeneralRequestIdentify,
		// RespID: lingo.GeneralIdentify,
		Exec: func(cmd iap.Command) *iap.Command {
			return &iap.Command{}
		}},
	{
		CmdID: lingo.GeneralIdentify,
		// RespID: lingo.GeneralACK,
		Exec: func(cmd iap.Command) *iap.Command {
			return GeneralAck(lingo.GeneralACK, lingo.GeneralIdentify, lingo.AckOK)
		},
	},

	// Remote UI Mode
	{
		CmdID: lingo.GeneralRequestRemoteUIMode,
		// RespID: lingo.GeneralReturnRemoteUIMode,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetRemoteUIMode()
		}},
	{
		CmdID: lingo.GeneralEnterRemoteUIMode,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			EnterRemoteUIMode()
			return nil
		}},
	{
		CmdID: lingo.GeneralExitRemoteUIMode,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			ExitRemoteUIMode()
			return nil
		}},

	// iPod Name
	{
		CmdID: lingo.GeneralRequestiPodName,
		// RespID: lingo.GeneralReturniPodName,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodName()
		}},

	// Software Version
	{
		CmdID: lingo.GeneralRequestiPodSoftwareVersion,
		// RespID: lingo.GeneralReturniPodSoftwareVersion,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodSoftwareVersion()
		}},

	// Serial Number
	{
		CmdID: lingo.GeneralRequestiPodSerialNum,
		// RespID: lingo.GeneralReturniPodSerialNum,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodSerialNum()
		}},

	// Model Number
	{
		CmdID: lingo.GeneralRequestiPodModelNum,
		// RespID: lingo.GeneralReturniPodModelNum,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodModelNum()
		}},

	// Lingo Protocol Version
	{
		CmdID: lingo.GeneralRequestLingoProtocolVersion,
		// RespID: lingo.GeneralReturnLingoProtocolVersion,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetLingoProtocolVersion(cmd.Lingo)
		}},

	// Device Lingoes
	{
		CmdID: lingo.GeneralIdentifyDeviceLingoes,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			IdentifyDeviceLingoes()
			return nil
		}},

	// Device Authentication
	{
		CmdID: lingo.GeneralGetDevAuthenticationInfo,
		// RespID: lingo.GeneralRetDevAuthenticationInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetDevAuthenticationInfo()
		}},
	{
		CmdID: lingo.GeneralAckDevAuthenticationInfo,
		// RespID: 0,
		Exec: nil,
	},
	{
		CmdID: lingo.GeneralGetDevAuthenticationSignature,
		// RespID: lingo.GeneralRetDevAuthenticationSignature,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetDevAuthenticationSignature()
		}},
	{
		CmdID: lingo.GeneralAckDevAuthenticationStatus,
		// RespID: 0,
		Exec: nil,
	},

	// iPod Authentication
	{
		CmdID: lingo.GeneralGetiPodAuthenticationInfo,
		// RespID: lingo.GeneralRetiPodAuthenticationInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodAuthenticationInfo()
		}},
	{
		CmdID: lingo.GeneralAckiPodAuthenticationInfo,
		// RespID: 0,
		Exec: nil,
	},
	{
		CmdID: lingo.GeneralGetiPodAuthenticationSignature,
		// RespID: lingo.GeneralRetiPodAuthenticationSignature,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodAuthenticationSignature()
		}},
	{
		CmdID: lingo.GeneralAckiPodAuthenticationStatus,
		// RespID: 0,
		Exec: nil,
	},

	// State Change Notification
	{
		CmdID: lingo.GeneralNotifyiPodStateChange,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			return HandleStateChange(cmd.CmdData)
		}},

	// iPod Options
	{
		CmdID: lingo.GeneralGetiPodOptions,
		// RespID: lingo.GeneralRetiPodOptions,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodOptions()
		}},

	// Accessory Info
	{
		CmdID: lingo.GeneralGetAccessoryInfo,
		// RespID: lingo.GeneralRetAccessoryInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetAccessoryInfo()
		}},

	// Preferences
	{
		CmdID: lingo.GeneralGetiPodPreferences,
		// RespID: lingo.GeneralRetiPodPreferences,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodPreferences()
		}},
	{
		CmdID: lingo.GeneralSetiPodPreferences,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			SetiPodPreferences(cmd.CmdData)
			return nil
		}},

	// IDPS
	{
		CmdID: lingo.GeneralStartIDPS,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			StartIDPS()
			return nil
		}},
	{
		CmdID: lingo.GeneralSetFIDTokenValues,
		// RespID: lingo.GeneralRetFIDTokenValueACKs,
		Exec: func(cmd iap.Command) *iap.Command {
			return SetFIDTokenValues(cmd.CmdData)
		}},
	{
		CmdID: lingo.GeneralEndIDPS,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			EndIDPS()
			return nil
		}},
	{
		CmdID: lingo.GeneralIDPSStatus,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetIDPSStatus()
		}},

	// Data Session
	{
		CmdID: lingo.GeneralOpenDataSessionForProtocol,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			return OpenDataSession(cmd.CmdData)
		}},
	{
		CmdID: lingo.GeneralCloseDataSession,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			return CloseDataSession(cmd.CmdData)
		}},
	{
		CmdID: lingo.GeneralDevACK,
		// RespID: 0,
		Exec: nil,
	},
	{
		CmdID: lingo.GeneralDevDataTransfer,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			return HandleDevDataTransfer(cmd.CmdData)
		}},
	{
		CmdID: lingo.GeneraliPodDataTransfer,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			return HandleiPodDataTransfer(cmd.CmdData)
		}},

	// Event Notifications
	{
		CmdID: lingo.GeneralSetEventNotification,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			return SetEventNotification(cmd.CmdData)
		}},
	{
		CmdID: lingo.GeneraliPodNotification,
		// RespID: 0,
		Exec: func(cmd iap.Command) *iap.Command {
			return HandleiPodNotification(cmd.CmdData)
		}},
	{
		CmdID: lingo.GeneralGetiPodOptionsForLingo,
		// RespID: lingo.GeneralRetiPodOptionsForLingo,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetiPodOptionsForLingo(cmd.CmdData)
		}},
	{
		CmdID: lingo.GeneralGetEventNotification,
		// RespID: lingo.GeneralRetEventNotification,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetEventNotification()
		}},
	{
		CmdID: lingo.GeneralGetSupportedEventNotification,
		// RespID: lingo.GeneralRetSupportedEventNotification,
		Exec: func(cmd iap.Command) *iap.Command {
			return GetSupportedEventNotification()
		}},
}
