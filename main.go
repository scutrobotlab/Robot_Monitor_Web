package main

import (
	"www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/SerialHandle"
	"www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/WebHandle"
)

func main() {
	defer serialhandle.CloseSerialPort()
	webhandle.WebHandleStart()
}
