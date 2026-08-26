package ble

import (
	"context"
	"errors"

	"tinygo.org/x/bluetooth"
)

var (
	adapter    = bluetooth.DefaultAdapter
	hidControl bluetooth.Characteristic
)

// InitRemote configures the BLE stack, registers HID characteristics, and tracks connection states
func InitRemote(cancel context.CancelFunc) error {
	if adapter == nil {
		return errors.New("missing bluetooth hardware radio")
	}

	if err := adapter.Enable(); err != nil {
		return err
	}

	// 1. Define standard Bluetooth SIG Service UUIDs for a media remote
	hidServiceUUID := bluetooth.New16BitUUID(0x1812) // Human Interface Device
	reportCharUUID := bluetooth.New16BitUUID(0x2A4D) // HID Report

	// 2. Register the HID attributes to the local GATT server
	err := adapter.AddService(&bluetooth.Service{
		UUID: hidServiceUUID,
		Characteristics: []bluetooth.CharacteristicConfig{
			{
				Handle: &hidControl,
				UUID:   reportCharUUID,
				Value:  []byte{0x00}, // Initial state: no buttons pressed
				Flags:  bluetooth.CharacteristicReadPermission | bluetooth.CharacteristicNotifyPermission,
			},
		},
	})
	if err != nil {
		return err
	}

	// 3. Keep your exact connection state callback hooks
	adapter.SetConnectHandler(func(device bluetooth.Device, connected bool) {
		if connected {
			print("Device connected: ")
			println(device.Address.String())
			return
		}

		print("Device disconnected: ")
		println(device.Address.String())
		cancel()
	})

	return nil
}

// StartBeacon configures your custom advertising data and launches the radio beacon
func StartBeacon() error {
	adv := adapter.DefaultAdvertisement()

	err := adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: "Go Bluetooth",
		ManufacturerData: []bluetooth.ManufacturerDataElement{
			{CompanyID: 0xffff, Data: []byte{0x01, 0x02}},
		},
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
	_, err := hidControl.Write([]byte{keyMask})
	if err != nil {
		return err
	}
	// Instantly release the key state so it doesn't get stuck in a "long-press" loop
	_, err = hidControl.Write([]byte{0x00})
	return err
}
