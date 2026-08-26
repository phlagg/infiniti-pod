package ble

import "tinygo.org/x/bluetooth"

// ServiceManager is the interface for managing Bluetooth services.
type ServiceManager interface {
	// AddService adds a Bluetooth service with the given UUID and characteristics.
	AddService(uuid string, characteristics []bluetooth.Characteristic) error

	// RemoveService removes a Bluetooth service with the given UUID.
	RemoveService(uuid string) error

	// ListServices lists all available Bluetooth services.
	ListServices() ([]string, error)
}
