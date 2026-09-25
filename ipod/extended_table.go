package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/iap/lingo"
)

var ExtendedTable = []commandEntry{
	// ACK
	{
		CmdID: lingo.ExtIfaceACK,
		Exec: func(cmd iap.Command) *iap.Command {
			return nil
		},
	},

	// Chapter info
	{
		CmdID: lingo.ExtIfaceGetCurrentPlayingChapterInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetCurrentPlayingChapterInfo()
			resp.CmdID = lingo.ExtIfaceReturnCurrentPlayingChapterInfo
			return resp
		},
	},
	{
		CmdID: lingo.ExtIfaceSetCurrentPlayingChapter,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := SetCurrentPlayingChapter(cmd.CmdData)
			resp.CmdID = 0
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceGetCurrentPlayingChapterPlayStatus,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetCurrentPlayingChapterPlayStatus()
			resp.CmdID = lingo.ExtIfaceReturnCurrentPlayingChapterPlayStatus
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceGetCurrentPlayingChapterName,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetCurrentPlayingChapterName()
			resp.CmdID = lingo.ExtIfaceReturnCurrentPlayingChapterName
			return resp
		},
	},

	// Audiobook speed
	{
		CmdID: lingo.ExtIfaceGetAudiobookSpeed,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetAudiobookSpeed()
			resp.CmdID = lingo.ExtIfaceReturnAudiobookSpeed
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceSetAudiobookSpeed,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := SetAudiobookSpeed(cmd.CmdData)
			resp.CmdID = 0
			return resp
		},
	},

	// Indexed track info
	{
		CmdID: lingo.ExtIfaceGetIndexedPlayingTrackInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetIndexedPlayingTrackInfo(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnIndexedPlayingTrackInfo
			return resp
		},
	},

	// Artwork formats
	{
		CmdID: lingo.ExtIfaceGetArtworkFormats,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetArtworkFormats()
			resp.CmdID = lingo.ExtIfaceReturnArtworkFormats
			return resp
		},
	},

	// Artwork data
	{
		CmdID: lingo.ExtIfaceGetTrackArtworkData,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetTrackArtworkData(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnTrackArtworkData
			return resp
		},
	},

	// Protocol version

	{
		CmdID: lingo.ExtIfaceRequestProtocolVersion,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetLingoProtocolVersion(cmd.Lingo)
			return resp
		},
	},

	// iPod name
	{
		CmdID: lingo.ExtIfaceRequestiPodName,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetiPodName()
			resp.CmdID = lingo.ExtIfaceReturniPodName
			return resp
		},
	},

	// DB selection
	{
		CmdID: lingo.ExtIfaceResetDBSelection,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := ResetDBSelection()
			resp.CmdID = 0
			return resp
		},
	},
	{
		CmdID: lingo.ExtIfaceSelectDBRecord,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := SelectDBRecord(cmd.CmdData)
			resp.CmdID = 0
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceGetNumberCategorizedDBRecords,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetNumberCategorizedDBRecords(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnNumberCategorizedDBRecords
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceRetrieveCategorizedDBRecords,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := RetrieveCategorizedDBRecord(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnCategorizedDBRecord
			return resp
		},
	},

	// Play status
	{
		CmdID: lingo.ExtIfaceGetPlayStatus,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetPlayStatus()
			resp.CmdID = lingo.ExtIfaceReturnPlayStatus
			return resp
		},
	},

	// Track index
	{
		CmdID: lingo.ExtIfaceGetCurrentPlayingTrackIndex,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetCurrentPlayingTrackIndex()
			resp.CmdID = lingo.ExtIfaceReturnCurrentPlayingTrackIndex
			return resp
		},
	},

	// Track metadata
	{
		CmdID: lingo.ExtIfaceGetIndexedPlayingTrackTitle,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetIndexedPlayingTrackTitle(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnIndexedPlayingTrackTitle
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceGetIndexedPlayingTrackArtistName,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetIndexedPlayingTrackArtistName(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnIndexedPlayingTrackArtistName
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceGetIndexedPlayingTrackAlbumName,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetIndexedPlayingTrackAlbumName(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnIndexedPlayingTrackAlbumName
			return resp
		},
	},

	// Play status notifications
	{
		CmdID: lingo.ExtIfaceSetPlayStatusChangeNotification,
		Exec: func(cmd iap.Command) *iap.Command {
			resp, err := SetPlayStatusChangeNotification(cmd.CmdData)
			if err != nil {
				return nil
			}
			resp.CmdID = 0
			return resp
		},
	},

	// Playback control
	{
		CmdID: lingo.ExtIfacePlayCurrentSelection,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := PlayCurrentSelection()
			resp.CmdID = 0
			return resp
		},
	},
	{
		CmdID: lingo.ExtIfacePlayControl,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := PlayControl(&cmd)
			return resp
		},
	},

	// Artwork timing
	{
		CmdID: lingo.ExtIfaceGetTrackArtworkTimes,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetTrackArtworkTimes(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnTrackArtworkTimes
			return resp
		},
	},

	// Shuffle
	{
		CmdID: lingo.ExtIfaceGetShuffle,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetShuffle()
			resp.CmdID = lingo.ExtIfaceReturnShuffle
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceSetShuffle,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := SetShuffle(cmd.CmdData)
			resp.CmdID = 0
			return resp
		},
	},

	// Repeat
	{
		CmdID: lingo.ExtIfaceGetRepeat,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetRepeat()
			resp.CmdID = lingo.ExtIfaceReturnRepeat
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceSetRepeat,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := SetRepeat(cmd.CmdData)
			resp.CmdID = 0
			return resp
		},
	},

	// Display image
	{
		CmdID: lingo.ExtIfaceSetDisplayImage,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := SetDisplayImage(cmd.CmdData)
			resp.CmdID = 0
			return resp
		},
	},

	{
		CmdID: lingo.ExtIfaceGetMonoDisplayImageLimits,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetMonoDisplayImageLimits()
			resp.CmdID = lingo.ExtIfaceReturnMonoDisplayImageLimits
			return resp
		},
	},

	// Number of tracks
	{
		CmdID: lingo.ExtIfaceGetNumPlayingTracks,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetNumPlayingTracks()
			resp.CmdID = lingo.ExtIfaceReturnNumPlayingTracks
			return resp
		},
	},

	// Track selection
	{
		CmdID: lingo.ExtIfaceSetCurrentPlayingTrack,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := SetCurrentPlayingTrack(cmd.CmdData)
			resp.CmdID = 0
			return resp
		},
	},
	{
		CmdID: lingo.ExtIfaceSelectSortDBRecord,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := SelectSortDBRecord(cmd.CmdData)
			resp.CmdID = 0
			return resp
		},
	},

	// Color display limits
	{
		CmdID: lingo.ExtIfaceGetColorDisplayImageLimits,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetColorDisplayImageLimits()
			resp.CmdID = lingo.ExtIfaceReturnColorDisplayImageLimits
			return resp
		},
	},

	// DB hierarchy
	{
		CmdID: lingo.ExtIfaceResetDBSelectionHierarchy,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := ResetDBSelectionHierarchy()
			resp.CmdID = 0x00
			return resp
		},
	},

	// iTunes DB info
	{
		CmdID: lingo.ExtIfaceGetDBiTunesInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetDBiTunesInfo(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnDBiTunesInfo
			return resp
		},
	},

	// UID track info
	{
		CmdID: lingo.ExtIfaceGetUIDTrackInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetUIDTrackInfo(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnUIDTrackInfo
			return resp
		},
	},

	// DB track info
	{
		CmdID: lingo.ExtIfaceGetDBTrackInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetDBTrackInfo(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnDBTrackInfo
			return resp
		},
	},

	// PB track info
	{
		CmdID: lingo.ExtIfaceGetPBTrackInfo,
		Exec: func(cmd iap.Command) *iap.Command {
			resp := GetPBTrackInfo(cmd.CmdData)
			resp.CmdID = lingo.ExtIfaceReturnPBTrackInfo
			return resp
		},
	},
}
