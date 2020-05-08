package main

import (
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	filehandle "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/FileHandle"
	serialhandle "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/SerialHandle"
	webhandle "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/WebHandle"
)

func main() {
	filehandle.LoadSaves()
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				filehandle.SaveAll()
				serialhandle.CloseSerialPort()
				os.Exit(0)
			default:
				os.Exit(0)
			}
		}
	}()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			filehandle.SaveAll()
		}
	}()
	http.Handle("/", http.FileServer(http.Dir("./WebPage/")))
	webhandle.Start()
}
