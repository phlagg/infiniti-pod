package main

import (
	"sync/atomic"
	"time"

	"github.com/phlagg/infiniti-pod/ipod"
	"github.com/phlagg/infiniti-pod/transport/ble"
	"github.com/phlagg/infiniti-pod/transport/serial"
	"github.com/phlagg/infiniti-pod/transport/usb"
)

var bleDisconnected uint32

func main() {
	time.Sleep(time.Second * 2)
	logInfo("Initializing Infiniti-Pod Protocol Systems...")

	// 1. Initialize Car Bus Channel
	err := serial.Init()
	must("open serial interface with vehicle", err)

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

		ipod.Run()

		// Sleep minimally
		time.Sleep(time.Millisecond * 1)

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
