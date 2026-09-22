package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/iap/lingo"
)

type commandEntry struct {
	CmdID  uint16
	RespID uint16
	Exec   func(cmd iap.Command) []byte
}

func handleCommand(cmd iap.Command) iap.Response {

	switch cmd.Lingo {
	case lingo.LingoGeneralID:
		return handleLingo(cmd, GeneralTable)
	case lingo.LingoExtendedID:
		return handleLingo(cmd, ExtendedTable)
	default:

	}

	return iap.Response{}
}

func handleLingo(cmd iap.Command, table []commandEntry) iap.Response {
	for _, e := range table {
		if e.CmdID == cmd.CmdID {
			return iap.Response{
				Lingo:   cmd.Lingo,
				CmdID:   e.RespID,
				CmdData: e.Exec(cmd),
			}
		}
	}
	return iap.Response{}
}

// var ExtendedTable = []commandEntry{
// 	{
// 		CmdID:  lingo.ExtIfaceRequestProtocolVersion,
// 		RespID: lingo.ExtIfaceReturnProtocolVersion,
// 		Exec:   func(cmd iap.Command) []byte { return GetProtocolVersion() },
// 	},
// 	{
// 		CmdID:  lingo.ExtIfaceRequestiPodName,
// 		RespID: lingo.ExtIfaceReturniPodName,
// 		Exec:   func(cmd iap.Command) []byte { return GetiPodName() },
// 	},
// 	{
// 		CmdID:  lingo.ExtIfaceGetPlayStatus,
// 		RespID: lingo.ExtIfaceReturnPlayStatus,
// 		Exec:   func(cmd iap.Command) []byte { return GetPlayStatus() },
// 	},
// 	{
// 		CmdID:  lingo.ExtIfaceSetPlayStatusChangeNotification,
// 		RespID: lingo.ExtIfaceReturnPlayStatus,
// 		Exec: func(cmd iap.Command) []byte {
// 			SetPlayStatusChangeNotification(cmd.CmdData[0])
// 			return []byte{0x00}
// 		},
// 	},
// }
