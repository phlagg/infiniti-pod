package iap_test

import (
	"encoding/hex"
	"os"
	"strings"
	"testing"
	"time"

	"golang.org/x/sys/unix"
)

var cases = []struct {
	name     string
	payload  string
	expected string
	note     string
}{
	{
		name:     "SendLingoSupport",
		payload:  "FF FF 55 03 00 01 04 F8",
		expected: "ff550400020001f9",
		note:     "Sends Ack",
	},
	{
		name:     "RequestExtendedLingoVersion",
		payload:  "FF 55 03 04 00 12 E7",
		expected: "ff55050400130114cf",
		note:     "Starts Handshake",
	},
	{
		name:     "RequestGeneralLingoVersion",
		payload:  "FF 55 02 00 0F EF",
		expected: "ff550400100111da",
		note:     "Starts Handshake",
	},
	{
		name:     "RequestiPodName",
		payload:  "FF 55 03 04 00 14 E5",
		expected: "ff55120400154d69636861656c27732050686f6e656e",
		note:     "Returns 'Michael's Phone'",
	},
	{
		name:     "GetPlayStatus",
		payload:  "FF 55 03 04 00 1C DD",
		expected: "ff550c04001d000249f00000753002f1",
		note:     "Returns Play Status",
	},

	// Button commands
	{
		name:     "Play",
		payload:  "FF 55 04 04 00 29 0A C5",
		expected: "ff5506040001000029cc",
		note:     "Triggers ble.KeyPlay",
	},
	{
		name:     "Wait",
		payload:  "FF 55 03 00 00 00 FD",
		expected: "ff55020000fe",
		note:     "Waiting...",
	},
	{
		name:     "Next Track",
		payload:  "FF 55 04 04 00 29 03 CC",
		expected: "ff5506040001000029cc",
		note:     "Triggers ble.KeyNext",
	},
	{
		name:     "Wait",
		payload:  "FF 55 03 00 00 00 FD",
		expected: "ff55020000fe",
		note:     "Waiting...",
	},
	{
		name:     "Stop",
		payload:  "FF 55 04 04 00 29 02 CD",
		expected: "ff5506040001000029cc",
		note:     "Triggers ble.KeyPlayPause",
	},
	{
		name:     "Wait",
		payload:  "FF 55 03 00 00 00 FD",
		expected: "ff55020000fe",
		note:     "Waiting...",
	},
	{
		name:     "Previous Track",
		payload:  "FF 55 04 04 00 29 04 CB",
		expected: "ff5506040001000029cc",
		note:     "Triggers ble.KeyPrevious",
	},
	{
		name:     "Wait",
		payload:  "FF 55 03 00 00 00 FD",
		expected: "ff55020000fe",
		note:     "Waiting...",
	},
	{
		name:     "Play / Pause",
		payload:  "FF 55 04 04 00 29 01 CE",
		expected: "ff5506040001000029cc",
		note:     "Triggers ble.KeyPlayPause",
	},
}

func TestHostCommands(t *testing.T) {
	port := "/dev/ttyUSB0"

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

			// strip spaces from payload hex
			hexPayload := strings.ReplaceAll(tc.payload, " ", "")
			payloadBytes, err := hex.DecodeString(hexPayload)
			if err != nil {
				t.Fatalf("bad payload hex: %v", err)
			}

			got := sendAndRecv(t, port, payloadBytes)
			gotHex := hexify(got)

			if gotHex != tc.expected {
				t.Fatalf("got %s, want %s (%s)", gotHex, tc.expected, tc.note)
			}
		})
		time.Sleep(1 * time.Second)
	}
}

func sendAndRecv(t *testing.T, port string, payload []byte) []byte {
	f, err := openSerialRaw(port)
	if err != nil {
		t.Fatalf("open: %v", err)
	}
	defer f.Close()

	// Write raw bytes
	_, err = f.Write(payload)
	if err != nil {
		t.Fatalf("write: %v", err)
	}

	// Allow MCU to respond
	f.SetReadDeadline(time.Now().Add(500 * time.Millisecond))

	buf := make([]byte, 64)
	n, err := f.Read(buf)
	if err != nil {
		t.Fatalf("read: %v", err)
	}

	return buf[:n]
}

func hexify(b []byte) string {
	const h = "0123456789abcdef"
	out := make([]byte, len(b)*2)
	for i, v := range b {
		out[i*2] = h[v>>4]
		out[i*2+1] = h[v&0xF]
	}
	return string(out)
}

func openSerialRaw(path string) (*os.File, error) {
	// Open using unix.Open so we can use proper flags
	fd, err := unix.Open(path, unix.O_RDWR|unix.O_NOCTTY|unix.O_NONBLOCK, 0666)
	if err != nil {
		return nil, err
	}

	// Convert fd to *os.File
	f := os.NewFile(uintptr(fd), path)

	// Get current termios
	tio, err := unix.IoctlGetTermios(fd, unix.TCGETS)
	if err != nil {
		f.Close()
		return nil, err
	}

	// ----- RAW MODE -----
	// Disable input processing
	tio.Iflag &^= unix.IGNBRK | unix.BRKINT | unix.PARMRK |
		unix.ISTRIP | unix.INLCR | unix.IGNCR | unix.ICRNL | unix.IXON

	// Disable output processing
	tio.Oflag &^= unix.OPOST

	// Disable canonical mode, echo, signals
	tio.Lflag &^= unix.ECHO | unix.ECHONL | unix.ICANON |
		unix.ISIG | unix.IEXTEN

	// Disable character processing
	tio.Cflag &^= unix.CSIZE | unix.PARENB
	tio.Cflag |= unix.CS8

	// Minimum read size = 1 byte
	tio.Cc[unix.VMIN] = 1
	// Timeout = 0 (no timeout)
	tio.Cc[unix.VTIME] = 0
	// ---------------------

	// Apply settings
	if err := unix.IoctlSetTermios(fd, unix.TCSETS, tio); err != nil {
		f.Close()
		return nil, err
	}

	return f, nil
}
