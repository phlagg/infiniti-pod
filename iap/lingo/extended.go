package lingo

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

	ExtIfaceRequestProtocolVersion = 0x0012
	ExtIfaceReturnProtocolVersion  = 0x0013
	ExtIfaceRequestiPodName        = 0x0014
	ExtIfaceReturniPodName         = 0x0015

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
