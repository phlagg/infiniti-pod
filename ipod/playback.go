package ipod

import (
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/iap/lingo"
	"github.com/phlagg/infiniti-pod/transport/ble"
)

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
func GetRemoteUIMode() *iap.Command { return nil }
func EnterRemoteUIMode()            {}
func ExitRemoteUIMode()             {}

// Extended
func GetPlayStatus() *iap.Command {
	var trackTimeMs uint32 = 150_000
	var trackPositionMs uint32 = 30_000
	var playerState byte = PlayerStatePaused
	logStatus("GetPlayStatus", "trackTimeMs", trackTimeMs, "trackPositionMs", trackPositionMs, "playerState", playerState)
	return &iap.Command{
		CmdData: buildPlayStatus(trackTimeMs, trackPositionMs, playerState),
	}
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

func PlayCurrentSelection() *iap.Command { return nil }

type PlayControlCmd byte

const (
	PlayControlToggle      PlayControlCmd = 0x01
	PlayControlStop        PlayControlCmd = 0x02
	PlayControlNextTrack   PlayControlCmd = 0x03
	PlayControlPrevTrack   PlayControlCmd = 0x04
	PlayControlStartFF     PlayControlCmd = 0x05
	PlayControlStartRew    PlayControlCmd = 0x06
	PlayControlEndFFRew    PlayControlCmd = 0x07
	PlayControlNext        PlayControlCmd = 0x08
	PlayControlPrev        PlayControlCmd = 0x09
	PlayControlPlay        PlayControlCmd = 0x0a
	PlayControlPause       PlayControlCmd = 0x0b
	PlayControlNextChapter PlayControlCmd = 0x0c
	PlayControlPrevChapter PlayControlCmd = 0x0d
)

func PlayControl(c *iap.Command) *iap.Command {
	handlePlaybackCommands(PlayControlCmd(c.CmdData[0]))
	return ExtendedAck(lingo.ExtIfaceACK, c.CmdID, lingo.AckOK)
}

// handlePlaybackCommands bridges the vehicle iAP commands over to Bluetooth media keys
func handlePlaybackCommands(cmd PlayControlCmd) {
	switch cmd {
	case PlayControlToggle:
		println("[BRIDGE] Vehicle command: PLAY/PAUSE -> Notifying Phone")
		ble.PressMediaKey(ble.KeyPlayPause)
	case PlayControlStop:
		println("[BRIDGE] Vehicle command: STOP -> Notifying Phone")
		ble.PressMediaKey(ble.KeyStop)
	case PlayControlNextTrack:
		println("[BRIDGE] Vehicle command: NEXT -> Notifying Phone")
		ble.PressMediaKey(ble.KeyNext)
	case PlayControlPrevTrack:
		println("[BRIDGE] Vehicle command: PREVIOUS -> Notifying Phone")
		ble.PressMediaKey(ble.KeyPrevious)
	case PlayControlPlay:
		println("[BRIDGE] Vehicle command: PLAY -> Notifying Phone")
		ble.PressMediaKey(ble.KeyPlay)
	}
	ble.PressMediaKey(ble.KeyRelease)
}
func GetCurrentPlayingTrackIndex() *iap.Command       { return nil }
func SetCurrentPlayingTrack(data []byte) *iap.Command { return nil }

func GetIndexedPlayingTrackInfo(data []byte) *iap.Command       { return nil }
func GetIndexedPlayingTrackTitle(data []byte) *iap.Command      { return nil }
func GetIndexedPlayingTrackArtistName(data []byte) *iap.Command { return nil }
func GetIndexedPlayingTrackAlbumName(data []byte) *iap.Command  { return nil }

func GetShuffle() *iap.Command            { return nil }
func SetShuffle(data []byte) *iap.Command { return nil }

func GetRepeat() *iap.Command            { return nil }
func SetRepeat(data []byte) *iap.Command { return nil }

func GetTrackArtworkTimes(data []byte) *iap.Command { return nil }

func GetNumPlayingTracks() *iap.Command { return nil }

// Audiobook
func GetCurrentPlayingChapterInfo() *iap.Command        { return nil }
func SetCurrentPlayingChapter(data []byte) *iap.Command { return nil }
func GetCurrentPlayingChapterPlayStatus() *iap.Command  { return nil }
func GetCurrentPlayingChapterName() *iap.Command        { return nil }

func GetAudiobookSpeed() *iap.Command            { return nil }
func SetAudiobookSpeed(data []byte) *iap.Command { return nil }
