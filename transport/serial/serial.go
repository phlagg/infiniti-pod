package serial

import "machine"

// Configure the hardware UART interface
var UART = machine.UART0
var initialized = false

type Serial *machine.UART

// Initialize sets up the physical 19200 8N1 car line connection
// Standard legacy Apple iAP accessory baud rate
func Init() error {
	UART.Configure(machine.UARTConfig{BaudRate: 19200})
	initialized = true
	return nil
}

func Buffered() int {
	if !initialized {
		return 0
	}
	return UART.Buffered()
}

func ReadByte() (byte, error) {
	if !initialized {
		return 0, nil
	}
	return UART.ReadByte()
}

func Write(buf []byte) (int, error) {
	if !initialized {
		return 0, nil
	}
	return UART.Write(buf)
}
