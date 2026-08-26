package main

import (
	"context"
	"time"

	"github.com/phlagg/infiniti-pod/ble"
	"github.com/phlagg/infiniti-pod/car"
	"github.com/phlagg/infiniti-pod/usb"
)

func main() {
	logInfo("Initializing Infiniti-Pod Protocol Systems...")

	// 1. Initialize Car Serial Bus Channel
	err := car.InitializeSerialBus()
	must("open vehicle serial UART link", err)

	// 2. Initialize USB System
	err = usb.InitAudioStream()
	must("initialize USB audio framework", err)

	// 3. Bluetooth configuration
	ctx, cancel := context.WithCancel(context.Background())
	err = ble.InitRemote(cancel)
	must("enable BLE stack infrastructure", err)
	err = ble.StartBeacon()
	must("activate advertising beacons", err)

	logInfo("Bridge listening for car controls...")

	// Real-time loop
	for {
		select {
		case <-ctx.Done():
			logInfo("Connection sequence terminated.")
			return
		default:
			// Process incoming car controls over iAP
			if buttonCode, found := car.ReadCarPacket(); found {
				handleCarSignals(buttonCode)
			}

			// Sleep minimally (1ms) to keep execution fast and real-time
			time.Sleep(time.Millisecond * 1)
		}
	}
}

// handleCarSignals bridges the vehicle iAP commands over to Bluetooth media keys
func handleCarSignals(code byte) {
	switch code {
	case 0x01:
		println("[BRIDGE] Vehicle command: NEXT -> Notifying Phone")
		ble.PressMediaKey(ble.KeyNext)
	case 0x08:
		println("[BRIDGE] Vehicle command: PREVIOUS -> Notifying Phone")
		ble.PressMediaKey(ble.KeyPrevious)
	case 0x02:
		println("[BRIDGE] Vehicle command: PLAY/PAUSE -> Notifying Phone")
		ble.PressMediaKey(ble.KeyPlayPause)
	case 0x00:
		// Button released event, ignore safely
	}
}

func must(action string, err error) {
	if err != nil {
		print("[FATAL] Failed to ")
		print(action)
		print(": ")
		println(err.Error())
		panic(err)
	}
}

func logInfo(msg string) {
	print("[SYS] ")
	println(msg)
}
