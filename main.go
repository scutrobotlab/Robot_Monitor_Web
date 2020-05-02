package main

import (
	"net/http"

	serialhandle "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/SerialHandle"
	webhandle "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/WebHandle"
)

func main() {
	defer serialhandle.CloseSerialPort()
	http.Handle("/", http.FileServer(assetFS()))
	webhandle.Start()
}
