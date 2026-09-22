package ipod

const (
	PlayerStateStopped byte = 0x00
	PlayerStatePlaying byte = 0x01
	PlayerStatePaused  byte = 0x02
	PlayerStateError   byte = 0xFF
)

func logStatus(status ...any) {
	print("[Playback] ")
	for _, s := range status {
		print(s, " ")
	}
	println()
}

// General
func GetRemoteUIMode() []byte { return nil }
func EnterRemoteUIMode()      {}
func ExitRemoteUIMode()       {}

// Extended
func GetPlayStatus() []byte {
	var trackTimeMs uint32 = 150_000
	var trackPositionMs uint32 = 30_000
	var playerState byte = PlayerStatePaused
	logStatus("GetPlayStatus", "trackTimeMs", trackTimeMs, "trackPositionMs", trackPositionMs, "playerState", playerState)
	return buildPlayStatus(trackTimeMs, trackPositionMs, playerState)
}

func buildPlayStatus(trackTimeMs, trackPositionMs uint32, playerState byte) []byte {
	b := make([]byte, 9)

	// time
	b[0] = byte(trackTimeMs >> 24)
	b[1] = byte(trackTimeMs >> 16)
	b[2] = byte(trackTimeMs >> 8)
	b[3] = byte(trackTimeMs)

	// position
	b[4] = byte(trackPositionMs >> 24)
	b[5] = byte(trackPositionMs >> 16)
	b[6] = byte(trackPositionMs >> 8)
	b[7] = byte(trackPositionMs)

	// state
	b[8] = playerState

	return b
}

func PlayCurrentSelection()          {}
func PlayControl(data []byte) []byte { return nil }

func GetCurrentPlayingTrackIndex() []byte { return nil }
func SetCurrentPlayingTrack(data []byte)  {}

func GetIndexedPlayingTrackInfo(data []byte) []byte       { return nil }
func GetIndexedPlayingTrackTitle(data []byte) []byte      { return nil }
func GetIndexedPlayingTrackArtistName(data []byte) []byte { return nil }
func GetIndexedPlayingTrackAlbumName(data []byte) []byte  { return nil }

func GetShuffle() []byte     { return nil }
func SetShuffle(data []byte) {}

func GetRepeat() []byte     { return nil }
func SetRepeat(data []byte) {}

func GetTrackArtworkTimes(data []byte) []byte { return nil }

func GetNumPlayingTracks() []byte { return nil }

// Audiobook
func GetCurrentPlayingChapterInfo() []byte       { return nil }
func SetCurrentPlayingChapter(data []byte)       {}
func GetCurrentPlayingChapterPlayStatus() []byte { return nil }
func GetCurrentPlayingChapterName() []byte       { return nil }

func GetAudiobookSpeed() []byte     { return nil }
func SetAudiobookSpeed(data []byte) {}
