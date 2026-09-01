package ble

import (
	"context"
	"errors"
	"time"

	"tinygo.org/x/bluetooth"
)

var (
	adapter              = bluetooth.DefaultAdapter
	hidControl           bluetooth.Characteristic
	reportCharacteristic bluetooth.Characteristic
)

type BLEEvent struct {
	Connected  bool
	MACAddress [6]byte // A fixed-size array copies data instantly by value
}

// A channel capable of holding 2 events buffer-free
var connectionEventChan = make(chan BLEEvent, 2)

// InitRemote configures the BLE stack, registers HID characteristics, and tracks connection states
func InitRemote(cancel context.CancelFunc) error {
	if adapter == nil {
		return errors.New("missing bluetooth hardware radio")
	}

	if err := adapter.Enable(); err != nil {
		return err
	}

	//  Register the HID attributes to the local GATT server
	err := adapter.AddService(&bluetooth.Service{
		UUID: bluetooth.ServiceUUIDHumanInterfaceDevice,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				UUID:  bluetooth.CharacteristicUUIDHIDInformation,
				Value: []byte{0x11, 0x01, 0x00, 0x03},
				Flags: bluetooth.CharacteristicReadPermission,
			},
			// 2. Report Map (Your custom media ReportDescriptor)
			{
				UUID:         bluetooth.CharacteristicUUIDReportMap,
				Value:        ReportDescriptor,
				Flags:        bluetooth.CharacteristicReadPermission,
				ReadSecurity: bluetooth.SecurityEncrypted,
			},
			// 3. Protocol Mode
			{
				UUID:          bluetooth.CharacteristicUUIDProtocolMode,
				Value:         []byte{0x01},
				Flags:         bluetooth.CharacteristicReadPermission | bluetooth.CharacteristicWriteWithoutResponsePermission,
				ReadSecurity:  bluetooth.SecurityEncrypted,
				WriteSecurity: bluetooth.SecurityEncrypted,
			},
			// 4. HID Control Point
			{
				Handle:        &hidControl,
				UUID:          bluetooth.CharacteristicUUIDHIDControlPoint,
				Value:         []byte{0x00},
				Flags:         bluetooth.CharacteristicWriteWithoutResponsePermission,
				WriteSecurity: bluetooth.SecurityEncrypted,
			},
			// 5. Actual HID Media Data Report (The keypress engine)
			{
				Handle:        &reportCharacteristic,
				UUID:          bluetooth.CharacteristicUUIDReport,
				Value:         []byte{0x00},
				Flags:         bluetooth.CharacteristicReadPermission | bluetooth.CharacteristicNotifyPermission,
				ReadSecurity:  bluetooth.SecurityEncrypted,
				WriteSecurity: bluetooth.SecurityEncrypted,

				Descriptors: []bluetooth.DescriptorConfig{
					{
						UUID:  bluetooth.New16BitUUID(0x2908),
						Value: []byte{0x01, 0x01},
						// Protect the report mapping descriptor metadata
						ReadSecurity: bluetooth.SecurityEncrypted,
					},
				},
			},
		},
	})

	if err != nil {
		return err
	}

	adapter.EnablePairing(bluetooth.PairingParams{
		IOCapabilities: bluetooth.IOCapsNone,
		LESC:           true,
		PairingCompleteHandler: func(device bluetooth.Device, err error) {
			if err != nil {
				println("pairing failed:", err.Error())
			} else {
				println("pairing complete")
			}
		},
	})

	// Start the background logging worker
	go connectionLogWorker()

	// connection state callback hooks
	adapter.SetConnectHandler(func(device bluetooth.Device, connected bool) {
		event := BLEEvent{
			Connected:  connected,
			MACAddress: device.Address.MAC, // Deep-copies the 6 bytes instantly
		}
		select {
		case connectionEventChan <- event:
		default:
			// Channel buffer full, drops message safely to avoid locking the ISR
		}

		if !connected {
			cancel()
		}
	})

	return nil
}

// StartBeacon configures your custom advertising data and launches the radio beacon
func StartBeacon() error {
	adv := adapter.DefaultAdvertisement()

	// Configure the beacon so Android and iOS recognize it as a pairing-ready keyboard remote
	err := adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "Go Bluetooth",

		ServiceUUIDs: []bluetooth.UUID{bluetooth.ServiceUUIDHumanInterfaceDevice},
		Appearance:   961, // keyboard

		// 2. Set an explicit advertising interval (approx. 50ms) so the phone catches the signal fast
		Interval: bluetooth.NewDuration(50 * time.Millisecond),
	})
	if err != nil {
		return err
	}

	return adv.Start()
}

// GetMACAddress pulls the live physical chip string safely
func GetMACAddress() string {
	address, err := adapter.Address()
	if err != nil {
		return "00:00:00:00:00:00"
	}
	return address.MAC.String()
}

// PressMediaKey writes a raw keycode change value directly to the connected phone
func PressMediaKey(keyMask byte) error {
	// Send active key down event
	pressBuffer := []byte{0x01, keyMask}
	_, err := reportCharacteristic.Write(pressBuffer)
	if err != nil {
		return err
	}
	// Instantly release the key state so it doesn't get stuck in a "long-press" loop
	releaseBuffer := []byte{0x01, KeyRelease}
	_, err = reportCharacteristic.Write(releaseBuffer)
	return err
}

func printHexNibble(nibble byte) {
	if nibble < 10 {
		print(string('0' + nibble))
	} else {
		print(string('A' + (nibble - 10)))
	}
}

func connectionLogWorker() {
	for event := range connectionEventChan {
		if event.Connected {
			print("Device connected: ")
		} else {
			print("Device disconnected: ")
		}
		for i := range event.MACAddress {
			if i > 0 {
				print(":")
			}

			b := event.MACAddress[i]
			printHexNibble(b >> 4)   // High nibble
			printHexNibble(b & 0x0F) // Low nibble
		}
		println()
	}
}
