package main

import (
	"sync/atomic"
	"time"

	"github.com/phlagg/infiniti-pod/ble"
	"github.com/phlagg/infiniti-pod/iap"
	"github.com/phlagg/infiniti-pod/usb"
)

var bleDisconnected uint32

func main() {
	logInfo("Initializing Infiniti-Pod Protocol Systems...")

	// 1. Initialize Car Bus Channel
	err := iap.Init()
	must("open vehicle iAP interface", err)

	// 2. Initialize USB System
	err = usb.InitAudioStream()
	must("initialize USB audio framework", err)

	// 3. Bluetooth configuration
	err = ble.InitRemote(func() {
		atomic.StoreUint32(&bleDisconnected, 1)
	})

	must("enable BLE stack infrastructure", err)
	err = ble.StartBeacon()
	must("activate advertising beacons", err)

	logInfo("Bridge listening for car controls...")

	for {
		if atomic.LoadUint32(&bleDisconnected) == 1 {
			logInfo("Connection sequence terminated.")
			continue
		}

		iap.ReadLoop()

		// Sleep minimally
		time.Sleep(time.Millisecond * 1)

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
