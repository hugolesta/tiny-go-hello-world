package main

import (
	"machine"
)
func blink(){
	led := machine.D2
    led.Configure(machine.PinConfig{Mode: machine.PinOutput})
    ledSwitch := true    

    for {
        led.Set(ledSwitch)
        ledSwitch = !ledSwitch
        delay(500)
    }
}
