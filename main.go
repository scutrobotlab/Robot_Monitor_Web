package main

import (
	"./SerialHandle"
	"./WebHandle"
)

func main() {
	defer SerialHandle.CloseSerialPort()
	WebHandle.WebHandleStart()
}
