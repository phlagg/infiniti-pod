package ipod

import (
	"errors"

	"github.com/phlagg/infiniti-pod/iap"
)

// General
func HandleStateChange(data []byte) *iap.Command { return nil }

func SetEventNotification(data []byte) *iap.Command   { return nil }
func HandleiPodNotification(data []byte) *iap.Command { return nil }

func GetEventNotification() *iap.Command              { return nil }
func GetSupportedEventNotification() *iap.Command     { return nil }
func GetiPodOptionsForLingo(data []byte) *iap.Command { return nil }

// Extended
func SetPlayStatusChangeNotification(cmd []byte) (*iap.Command, error) {
	switch len(cmd) {
	case 1:
		switch cmd[0] {
		case 0x00:
			println("Disabled all status event notifications")
		case 0x01:
			println("Enabled play status event notifications")
		default:
			return nil, errors.New("[Playback Engine] invalid status cmd")
		}
	case 4:
		maskBits := uint16(cmd[2]<<8 | cmd[3])
		switch maskBits {
		case (1 << 0):
			// Basic play state changes
		case (1 << 1):
			// Extended play state changes
		case (1 << 2):
			// Track index
		case (1 << 3):
			// Track time offset(ms)
		case (1 << 4):
			// Track time offset (sec)
		case (1 << 5):
			// Chapter index
		case (1 << 6):
			// Chapter time offset (ms)
		case (1 << 7):
			// Chapter time offset (sec)
		case (1 << 8):
			// Track identifier
		case (1 << 9):
			// Track media type
		case (1 << 10):
			// Track lyrics ready (if the track has lyrics)
		}
	default:
		return nil, errors.New("[Playback Engine] invalid status cmd")
	}
	return nil, nil
}
