package main

import (
	"machine"
	"sync/atomic"
	"time"

	"github.com/phlagg/infiniti-pod/ipod"
	"github.com/phlagg/infiniti-pod/transport/ble"
	"github.com/phlagg/infiniti-pod/transport/serial"
	"github.com/phlagg/infiniti-pod/transport/usb"
)

const (
	TransportUSB uint8 = 1
	TransportBLE uint8 = 2
)

type AppContext struct {
	initialized     bool
	btActive        bool
	usbConnected    bool
	transport       uint8
	bleDisconnected uint32
	usbDisconnected uint32
}

var ctx = AppContext{
	initialized:     false,
	btActive:        false,
	usbConnected:    false,
	transport:       0,
	bleDisconnected: 0,
	usbDisconnected: 0,
}

var led = machine.LED
var usbFlag = machine.P0_31

func main() {
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	usbFlag.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	// Initial system boot delay
	time.Sleep(time.Second * 2)

	// One-time startup sequence (never repeats, keeping the main loop flat)
	logInfo("Initializing Infiniti-Pod Protocol Systems...")
	must("open serial interface with vehicle", serial.Init())
	must("initialize USB audio framework", usb.InitAudioStream())
	must("initialize USB system", usb.USBInit())

	must("enable BLE stack infrastructure", ble.InitRemote(func() {
		atomic.StoreUint32(&ctx.bleDisconnected, 1)
	}))

	for {
		// 1. Read real-time hardware status metrics
		ctx.usbConnected = usb.USBConnected(&ctx.usbDisconnected)
		led.Set(ctx.btActive)

		// 2. ENFORCE ABSOLUTE USB PRIORITY RULE
		if ctx.usbConnected && !usbFlag.Get() {
			ctx.transport = TransportUSB

			// If we were running on Bluetooth when the USB cable was inserted, preempt instantly
			if ctx.btActive {
				logInfo("[PREEMPT] USB Host detected. Killing active Bluetooth radio.")
				_ = ble.StopBeacon()
				ctx.btActive = false
			}

			// Clear lingering USB disconnection flags if we are actively connected
			atomic.StoreUint32(&ctx.usbDisconnected, 0)

			// 3. FALLBACK TO BLUETOOTH STREAMING
		} else {
			ctx.transport = TransportBLE

			// If the BLE radio went cold or was never turned on, start advertising
			if !ctx.btActive {
				logInfo("[FALLBACK] USB port idle. Activating BLE advertising...")
				_ = ble.StopBeacon() // Clear state registers
				if err := ble.StartBeacon(); err != nil {
					logInfo("[BT] Warning during advertisement start: " + err.Error())
					time.Sleep(time.Millisecond * 100) // Brief debounce backoff on failure
					continue
				}
				ctx.btActive = true
			}

			// Catch runtime BLE disconnection events
			if atomic.LoadUint32(&ctx.bleDisconnected) == 1 {
				logInfo("[BT] Remote link terminated. Forcing radio reset.")
				atomic.StoreUint32(&ctx.bleDisconnected, 0)
				_ = ble.StopBeacon()
				ctx.btActive = false // Turning this false triggers a fresh StartBeacon on the next pass
				continue
			}
		}

		ipod.Run(ctx.transport, func() bool {
			if ctx.transport == TransportUSB {
				// While on USB, break out immediately if the host disconnects
				return !usb.USBConnected(&ctx.usbDisconnected)
			}
			// While on Bluetooth, break out immediately if a USB host plugs in
			return usb.USBConnected(&ctx.usbDisconnected) && !usbFlag.Get()
		})

		// 4. Handle USB Host cable removal loop cleanup
		if atomic.LoadUint32(&ctx.usbDisconnected) == 1 {
			logInfo("[USB] Host cable detached. Cycling connection transport context.")
			atomic.StoreUint32(&ctx.usbDisconnected, 0)
		}

		// Microcontroller yield sleep to keep CPU power draw stable
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
