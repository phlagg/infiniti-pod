package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/iap/lingo"
)

var ExtendedTable = []commandEntry{
	// ACK
	{
		CmdID:  lingo.ExtIfaceACK,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return nil
		},
	},

	// Chapter info
	{
		CmdID:  lingo.ExtIfaceGetCurrentPlayingChapterInfo,
		RespID: lingo.ExtIfaceReturnCurrentPlayingChapterInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetCurrentPlayingChapterInfo()
		},
	},
	{
		CmdID:  lingo.ExtIfaceSetCurrentPlayingChapter,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SetCurrentPlayingChapter(cmd.CmdData)
			return nil
		},
	},

	{
		CmdID:  lingo.ExtIfaceGetCurrentPlayingChapterPlayStatus,
		RespID: lingo.ExtIfaceReturnCurrentPlayingChapterPlayStatus,
		Exec: func(cmd iap.Command) []byte {
			return GetCurrentPlayingChapterPlayStatus()
		},
	},

	{
		CmdID:  lingo.ExtIfaceGetCurrentPlayingChapterName,
		RespID: lingo.ExtIfaceReturnCurrentPlayingChapterName,
		Exec: func(cmd iap.Command) []byte {
			return GetCurrentPlayingChapterName()
		},
	},

	// Audiobook speed
	{
		CmdID:  lingo.ExtIfaceGetAudiobookSpeed,
		RespID: lingo.ExtIfaceReturnAudiobookSpeed,
		Exec: func(cmd iap.Command) []byte {
			return GetAudiobookSpeed()
		},
	},

	{
		CmdID:  lingo.ExtIfaceSetAudiobookSpeed,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SetAudiobookSpeed(cmd.CmdData)
			return nil
		},
	},

	// Indexed track info
	{
		CmdID:  lingo.ExtIfaceGetIndexedPlayingTrackInfo,
		RespID: lingo.ExtIfaceReturnIndexedPlayingTrackInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetIndexedPlayingTrackInfo(cmd.CmdData)
		},
	},

	// Artwork formats
	{
		CmdID:  lingo.ExtIfaceGetArtworkFormats,
		RespID: lingo.ExtIfaceReturnArtworkFormats,
		Exec: func(cmd iap.Command) []byte {
			return GetArtworkFormats()
		},
	},

	// Artwork data
	{
		CmdID:  lingo.ExtIfaceGetTrackArtworkData,
		RespID: lingo.ExtIfaceReturnTrackArtworkData,
		Exec: func(cmd iap.Command) []byte {
			return GetTrackArtworkData(cmd.CmdData)
		},
	},

	// Protocol version

	{
		CmdID:  lingo.ExtIfaceRequestProtocolVersion,
		RespID: lingo.ExtIfaceReturnProtocolVersion,
		Exec: func(cmd iap.Command) []byte {
			return GetProtocolVersion()
		},
	},

	// iPod name
	{
		CmdID:  lingo.ExtIfaceRequestiPodName,
		RespID: lingo.ExtIfaceReturniPodName,
		Exec: func(cmd iap.Command) []byte {
			return GetiPodName()
		},
	},

	// DB selection
	{
		CmdID:  lingo.ExtIfaceResetDBSelection,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			ResetDBSelection()
			return nil
		},
	},
	{
		CmdID:  lingo.ExtIfaceSelectDBRecord,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SelectDBRecord(cmd.CmdData)
			return nil
		},
	},

	{
		CmdID:  lingo.ExtIfaceGetNumberCategorizedDBRecords,
		RespID: lingo.ExtIfaceReturnNumberCategorizedDBRecords,
		Exec: func(cmd iap.Command) []byte {
			return GetNumberCategorizedDBRecords(cmd.CmdData)
		},
	},

	{
		CmdID:  lingo.ExtIfaceRetrieveCategorizedDBRecords,
		RespID: lingo.ExtIfaceReturnCategorizedDBRecord,
		Exec: func(cmd iap.Command) []byte {
			return RetrieveCategorizedDBRecord(cmd.CmdData)
		},
	},

	// Play status
	{
		CmdID:  lingo.ExtIfaceGetPlayStatus,
		RespID: lingo.ExtIfaceReturnPlayStatus,
		Exec: func(cmd iap.Command) []byte {
			return GetPlayStatus()
		},
	},

	// Track index
	{
		CmdID:  lingo.ExtIfaceGetCurrentPlayingTrackIndex,
		RespID: lingo.ExtIfaceReturnCurrentPlayingTrackIndex,
		Exec: func(cmd iap.Command) []byte {
			return GetCurrentPlayingTrackIndex()
		},
	},

	// Track metadata
	{
		CmdID:  lingo.ExtIfaceGetIndexedPlayingTrackTitle,
		RespID: lingo.ExtIfaceReturnIndexedPlayingTrackTitle,
		Exec: func(cmd iap.Command) []byte {
			return GetIndexedPlayingTrackTitle(cmd.CmdData)
		},
	},

	{
		CmdID:  lingo.ExtIfaceGetIndexedPlayingTrackArtistName,
		RespID: lingo.ExtIfaceReturnIndexedPlayingTrackArtistName,
		Exec: func(cmd iap.Command) []byte {
			return GetIndexedPlayingTrackArtistName(cmd.CmdData)
		},
	},

	{
		CmdID:  lingo.ExtIfaceGetIndexedPlayingTrackAlbumName,
		RespID: lingo.ExtIfaceReturnIndexedPlayingTrackAlbumName,
		Exec: func(cmd iap.Command) []byte {
			return GetIndexedPlayingTrackAlbumName(cmd.CmdData)
		},
	},

	// Play status notifications
	{
		CmdID:  lingo.ExtIfaceSetPlayStatusChangeNotification,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SetPlayStatusChangeNotification(cmd.CmdData)
			return nil
		},
	},

	// Playback control
	{
		CmdID:  lingo.ExtIfacePlayCurrentSelection,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			PlayCurrentSelection()
			return nil
		},
	},
	{
		CmdID:  lingo.ExtIfacePlayControl,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return PlayControl(cmd.CmdData)
		},
	},

	// Artwork timing
	{
		CmdID:  lingo.ExtIfaceGetTrackArtworkTimes,
		RespID: lingo.ExtIfaceReturnTrackArtworkTimes,
		Exec: func(cmd iap.Command) []byte {
			return GetTrackArtworkTimes(cmd.CmdData)
		},
	},

	// Shuffle
	{
		CmdID:  lingo.ExtIfaceGetShuffle,
		RespID: lingo.ExtIfaceReturnShuffle,
		Exec: func(cmd iap.Command) []byte {
			return GetShuffle()
		},
	},

	{
		CmdID:  lingo.ExtIfaceSetShuffle,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SetShuffle(cmd.CmdData)
			return nil
		},
	},

	// Repeat
	{
		CmdID:  lingo.ExtIfaceGetRepeat,
		RespID: lingo.ExtIfaceReturnRepeat,
		Exec: func(cmd iap.Command) []byte {
			return GetRepeat()
		},
	},

	{
		CmdID:  lingo.ExtIfaceSetRepeat,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SetRepeat(cmd.CmdData)
			return nil
		},
	},

	// Display image
	{
		CmdID:  lingo.ExtIfaceSetDisplayImage,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			return SetDisplayImage(cmd.CmdData)
		},
	},

	{
		CmdID:  lingo.ExtIfaceGetMonoDisplayImageLimits,
		RespID: lingo.ExtIfaceReturnMonoDisplayImageLimits,
		Exec: func(cmd iap.Command) []byte {
			return GetMonoDisplayImageLimits()
		},
	},

	// Number of tracks
	{
		CmdID:  lingo.ExtIfaceGetNumPlayingTracks,
		RespID: lingo.ExtIfaceReturnNumPlayingTracks,
		Exec: func(cmd iap.Command) []byte {
			return GetNumPlayingTracks()
		},
	},

	// Track selection
	{
		CmdID:  lingo.ExtIfaceSetCurrentPlayingTrack,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SetCurrentPlayingTrack(cmd.CmdData)
			return nil
		},
	},
	{
		CmdID:  lingo.ExtIfaceSelectSortDBRecord,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			SelectSortDBRecord(cmd.CmdData)
			return nil
		},
	},

	// Color display limits
	{
		CmdID:  lingo.ExtIfaceGetColorDisplayImageLimits,
		RespID: lingo.ExtIfaceReturnColorDisplayImageLimits,
		Exec: func(cmd iap.Command) []byte {
			return GetColorDisplayImageLimits()
		},
	},

	// DB hierarchy
	{
		CmdID:  lingo.ExtIfaceResetDBSelectionHierarchy,
		RespID: 0,
		Exec: func(cmd iap.Command) []byte {
			ResetDBSelectionHierarchy()
			return nil
		},
	},

	// iTunes DB info
	{
		CmdID:  lingo.ExtIfaceGetDBiTunesInfo,
		RespID: lingo.ExtIfaceReturnDBiTunesInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetDBiTunesInfo(cmd.CmdData)
		},
	},

	// UID track info
	{
		CmdID:  lingo.ExtIfaceGetUIDTrackInfo,
		RespID: lingo.ExtIfaceReturnUIDTrackInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetUIDTrackInfo(cmd.CmdData)
		},
	},

	// DB track info
	{
		CmdID:  lingo.ExtIfaceGetDBTrackInfo,
		RespID: lingo.ExtIfaceReturnDBTrackInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetDBTrackInfo(cmd.CmdData)
		},
	},

	// PB track info
	{
		CmdID:  lingo.ExtIfaceGetPBTrackInfo,
		RespID: lingo.ExtIfaceReturnPBTrackInfo,
		Exec: func(cmd iap.Command) []byte {
			return GetPBTrackInfo(cmd.CmdData)
		},
	},
}
