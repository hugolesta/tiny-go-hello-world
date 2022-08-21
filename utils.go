package main

import "time"

func delay(t uint16) {
	time.Sleep(time.Duration(1000000 * uint32(t)))
}
