package none

import (
	"machine"
	"time"
)

func main2() {
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		go func() {
			led.Low()
			println("Low")
			time.Sleep(time.Millisecond * 200)

			led.High()
			println("High")
			time.Sleep(time.Millisecond * 200)
		}()

		time.Sleep(time.Millisecond * 500)
	}
}
