package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/iap/lingo"
)

type commandEntry struct {
	CmdID uint16
	Exec  func(cmd iap.Command) *iap.Command
}

func handleCommand(cmd iap.Command) *iap.Command {
	var commandTable []commandEntry = nil

	switch cmd.Lingo {
	case lingo.LingoGeneralID:
		commandTable = GeneralTable
	case lingo.LingoExtendedID:
		commandTable = ExtendedTable
	default:

	}
	for _, e := range commandTable {
		if e.CmdID == cmd.CmdID {
			resp := e.Exec(cmd)
			resp.Lingo = cmd.Lingo
			return resp

		}
	}
	return &iap.Command{}
}
