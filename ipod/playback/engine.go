package playback

import "errors"

const (
	PlayerStateStopped byte = 0x00
	PlayerStatePlaying byte = 0x01
	PlayerStatePaused  byte = 0x02
	PlayerStateError   byte = 0xFF
)

func GetPlayStatus() []byte {
	var trackTimeMs uint32 = 150_000
	var trackPositionMs uint32 = 30_000
	var playerState byte = PlayerStatePaused
	logStatus("GetPlayStatus", "trackTimeMs", trackTimeMs, "trackPositionMs", trackPositionMs, "playerState", playerState)
	return buildPlayStatus(trackTimeMs, trackPositionMs, playerState)
}

func SetPlayStatusChangeNotification(cmd byte) error {
	switch cmd {
	case 0x00:
		println("Disabled all status event notifications")
	case 0x01:
		println("Enabled play status event notifications")
	default:
		return errors.New("[Playback Engine] invalid status cmd")
	}
	return nil
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

func logStatus(status ...any) {
	print("[Playback] ")
	for _, s := range status {
		print(s, " ")
	}
	println()
}
