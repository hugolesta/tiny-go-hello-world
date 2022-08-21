package main

import (
	"machine"
	"time"
)
func digital_input(){
    button := machine.D10
    button.Configure(machine.PinConfig{Mode: machine.PinInput})

    led := machine.D9
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})

    for {
        led.Set(!button.Get())
        time.Sleep(time.Millisecond * 100)
    }
}

