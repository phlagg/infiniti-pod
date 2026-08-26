package ble

// Consumer Control HID Report Descriptor
// This raw byte map tells Android/iOS that your microcontroller is a media remote control.
var ReportDescriptor = []byte{
	0x05, 0x0C, // Usage Page (Consumer)
	0x09, 0x01, // Usage (Consumer Control)
	0xA1, 0x01, // Collection (Application)
	0x85, 0x01, //   Report ID (1)
	0x15, 0x00, //   Logical Minimum (0)
	0x25, 0x01, //   Logical Maximum (1)
	0x75, 0x01, //   Report Size (1)
	0x95, 0x04, //   Report Count (4) - 4 buttons total
	0x09, 0xB5, //   Usage (Scan Next Track)     - Bit 0
	0x09, 0xB6, //   Usage (Scan Previous Track) - Bit 1
	0x09, 0xCD, //   Usage (Play/Pause)          - Bit 2
	0x09, 0xE9, //   Usage (Volume Increment)    - Bit 3
	0x81, 0x02, //   Input (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0x75, 0x04, //   Report Size (4)             - 4 bits of padding
	0x95, 0x01, //   Report Count (1)            - 1 padding field
	0x81, 0x03, //   Input (Const,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0xC0, // End Collection
}

// Bitmask mappings matching the descriptor positions above
const (
	KeyNext      byte = 0x01 // 0000 0001
	KeyPrevious  byte = 0x02 // 0000 0010
	KeyPlayPause byte = 0x04 // 0000 0100
	KeyVolumeUp  byte = 0x08 // 0000 1000
	KeyRelease   byte = 0x00 // Clear all pressed states
)
