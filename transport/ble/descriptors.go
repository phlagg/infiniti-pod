package ble

// ReportDescriptor for an 8-button remote (Exactly 1 Byte)
var ReportDescriptor = []byte{
	0x05, 0x0C, // Usage Page (Consumer)
	0x09, 0x01, // Usage (Consumer Control)
	0xA1, 0x01, // Collection (Application)
	0x85, 0x01, //   Report ID (1)
	0x15, 0x00, //   Logical Minimum (0)
	0x25, 0x01, //   Logical Maximum (1)
	0x75, 0x01, //   Report Size (1 bit per button flag)
	0x95, 0x07, //   Report Count (7 buttons total)
	0x09, 0xCD, //   Usage (Play/Pause)          - Bit 0
	0x09, 0xB0, //   Usage (Play)                - Bit 1
	0x09, 0xB7, //   Usage (Stop)                - Bit 2
	0x09, 0xB5, //   Usage (Scan Next Track)     - Bit 3 (Skip Forward)
	0x09, 0xB6, //   Usage (Scan Previous Track) - Bit 4 (Skip Backward)
	0x09, 0xB3, //   Usage (Fast Forward)        - Bit 5 (Hold FF)
	0x09, 0xB4, //   Usage (Rewind)              - Bit 6 (Hold RW)
	0x81, 0x02, //   Input (Data,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0x75, 0x01, //   Report Size (1)             - 1 bit of remaining padding
	0x95, 0x01, //   Report Count (1)            - Aligns the packet to exactly 8 bits (1 byte)
	0x81, 0x03, //   Input (Const,Var,Abs,No Wrap,Linear,Preferred State,No Null Position)
	0xC0, // End Collection
}

// Bitmask mappings matching the descriptor positions above.
// To press a button, you send its bitmask byte.
// To release all buttons, you send KeyRelease (0x00).
const (
	KeyPlayPause byte = 1 << 0 // 0x01 (Bit 0)
	KeyPlay      byte = 1 << 1 // 0x02 (Bit 1)
	KeyStop      byte = 1 << 2 // 0x04 (Bit 2)
	KeyNext      byte = 1 << 3 // 0x08 (Bit 3)
	KeyPrevious  byte = 1 << 4 // 0x10 (Bit 4)
	KeyFF        byte = 1 << 5 // 0x20 (Bit 5)
	KeyRW        byte = 1 << 6 // 0x40 (Bit 6)

	KeyRelease byte = 0x00 // 0x00 (Clears all bits -> Releases all buttons)
)
