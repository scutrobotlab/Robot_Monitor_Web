package main

import (
	"./serialhandle"
	"./webhandle"
)

func main() {
	defer serialhandle.CloseSerialPort()
	webhandle.WebHandleStart()
}
