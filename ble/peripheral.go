package ble

import (
	"context"

	"tinygo.org/x/bluetooth"
)

var (
	localName = "TG"

	manufacturerName               = "TinyGo"
	manufacturerNameCharacteristic bluetooth.Characteristic

	modelNumber               = "Model 1"
	modelNumberCharacteristic bluetooth.Characteristic

	serialNumber               = "123456"
	serialNumberCharacteristic bluetooth.Characteristic

	// softwareVersion               = "1.0.0"
	// softwareVersionCharacteristic bluetooth.Characteristic

	// firmwareVersion               = "1.0.0"
	// firmwareVersionCharacteristic bluetooth.Characteristic

	// hardwareVersion               = "1.0.0"
	// hardwareVersionCharacteristic bluetooth.Characteristic
)

type Peripheral struct {
	adapter          *bluetooth.Adapter
	Address          bluetooth.MACAddress
	adv              *bluetooth.Advertisement
	ConnectedAddress bluetooth.Address
	deviceConnected  bool
	cancelFunc       context.CancelFunc
	Ctx              context.Context // New fields for deferred connection handling

	PendingConnectionChange struct {
		Device    bluetooth.Device
		Connected bool
		Trigger   bool
	}
}

// Init initializes the Bluetooth adapter and advertisement
func (p *Peripheral) Init() error {
	p.adapter = bluetooth.DefaultAdapter

	if err := p.adapter.Enable(); err != nil {
		panic("failed to enable adapter: %w" + err.Error())
	}
	println("[ble] Bluetooth adapter enabled")
	p.Address, _ = p.adapter.Address()
	p.deviceConnected = false
	return nil
}

func (p *Peripheral) SetConnectHandler() {
	p.Ctx, p.cancelFunc = context.WithCancel(context.Background())

	p.adapter.SetConnectHandler(func(device bluetooth.Device, connected bool) {
		p.deviceConnected = connected
		if connected {
			p.ConnectedAddress = device.Address
		} else {
			if p.cancelFunc != nil {
				p.cancelFunc()
			}
		}
		p.PendingConnectionChange.Device = device
		p.PendingConnectionChange.Connected = connected
		p.PendingConnectionChange.Trigger = true
	})
}

// StartAdvertising configures and starts BLE advertising
func (p *Peripheral) StartAdvertising() error {
	p.adv = p.adapter.DefaultAdvertisement()
	err := p.adv.Configure(bluetooth.AdvertisementOptions{
		LocalName: localName,
		ServiceUUIDs: []bluetooth.UUID{
			bluetooth.NewUUID([16]byte{0x18, 0x0F})},
		ManufacturerData: []bluetooth.ManufacturerDataElement{
			{CompanyID: 0xffff, Data: []byte{0x01}},
		},
	})
	if err != nil {
		panic("failed to configure advertisement: %w" + err.Error())
	}

	err = p.adv.Start()
	if err != nil {
		panic("failed to start advertising: %w" + err.Error())
	}
	println("[ble] Advertising started...")
	return nil
}

// StopAdvertising stops BLE advertising
func (p *Peripheral) StopAdvertising() error {
	return p.adv.Stop()
}

// SetupService creates a BLE service and characteristic
func (p *Peripheral) SetupService() error {
	serviceUUID := bluetooth.NewUUID([16]byte{0x18, 0x0F}) // Example service UUID
	// Define the characteristic
	char := bluetooth.CharacteristicConfig{
		Handle: &manufacturerNameCharacteristic,
		UUID:   bluetooth.CharacteristicUUIDManufacturerNameString,
		Value:  []byte(manufacturerName),
		Flags:  bluetooth.CharacteristicReadPermission,
	}
	char2 := bluetooth.CharacteristicConfig{
		Handle: &manufacturerNameCharacteristic,
		UUID:   bluetooth.CharacteristicUUIDModelNumberString,
		Value:  []byte(modelNumber),
		Flags:  bluetooth.CharacteristicReadPermission,
	}
	char3 := bluetooth.CharacteristicConfig{
		Handle: &serialNumberCharacteristic,
		UUID:   bluetooth.CharacteristicUUIDSerialNumberString,
		Value:  []byte(serialNumber),
		Flags:  bluetooth.CharacteristicReadPermission,
	}
	// Define the service and attach the characteristic
	service := bluetooth.Service{
		UUID:            serviceUUID,
		Characteristics: []bluetooth.CharacteristicConfig{char, char2, char3},
	}

	// Register the service with the adapter
	if err := p.adapter.AddService(&service); err != nil {
		panic("failed to add service: %wl" + err.Error())
	}

	return nil
}
