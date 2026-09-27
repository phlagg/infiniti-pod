package usb

import (
	"device/nrf"
	"errors"
	"machine/usb"
	"machine/usb/hid/keyboard"
	"sync/atomic"
	"time"
)

var kb = keyboard.Port()

func USBInit() error {
	if kb == nil {
		return errors.New("no keyboard port")
	}
	return nil
}

// USBConnected tracks real-time VBUS power presence on an nRF52840 chip.
func USBConnected(usbDisconnected *uint32) bool {
	// Query the raw NRF Power registers to see if VBUS cable voltage is physically detected
	vbusPresent := (nrf.POWER.USBREGSTATUS.Get() & nrf.POWER_USBREGSTATUS_VBUSDETECT) != 0

	if vbusPresent {
		// Device is plugged into a powered USB host
		if atomic.LoadUint32(usbDisconnected) == 1 {
			// Trigger a state change if we were previously marked disconnected
			atomic.StoreUint32(usbDisconnected, 0)
		}
		return true
	}

	// Cable yanked or plugged into a dead/unpowered port
	atomic.StoreUint32(usbDisconnected, 1)
	return false
}

func SendKey(key keyboard.Keycode) error {
	if kb == nil {
		return errors.New("no keyboard port")
	}
	kb.Press(key)
	return nil
}

func SendConsumerKey(keycode keyboard.Keycode) {
	// 1. Send Press Report: [Report ID 3, Low Byte, High Byte]
	pressReport := []byte{3, byte(keycode & 0xFF), byte(keycode >> 8)}
	kb.Write(pressReport)

	// Brief latch delay so Android registers the transition
	time.Sleep(20 * time.Millisecond)

	// 2. Send Release Report: Clear out data on Report ID 3
	releaseReport := []byte{3, 0x00, 0x00}
	kb.Write(releaseReport)

	time.Sleep(20 * time.Millisecond)
}

func SendAndroidShortcut(modifier keyboard.Keycode, key keyboard.Keycode) {
	// Write raw 8-byte packet structure expected by TinyGo's descriptor
	kb.Write([]byte{byte(modifier), byte(key)})
	time.Sleep(20 * time.Millisecond)
	kb.Write([]byte{0, 0}) // Explicit release packet
	time.Sleep(20 * time.Millisecond)
}
func usbSetupHandler(setup usb.Setup) {}

func usbTxHandler() {}

func usbRxHandler(b []byte) {}
