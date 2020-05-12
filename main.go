package main

import (
	"fmt"
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
	fmt.Println(`     ____            __              __                                     __                   
    /\  _ \         /\ \            /\ \__      / \_/ \                  __/\ \__                
    \ \ \_\ \    ___\ \ \____    ___\ \  _\    /\      \    ___     ___ /\_\ \  _\   ___   _ __  
     \ \    /   / __ \ \  __ \  / __ \ \ \/    \ \ \__\ \  / __ \ /  _  \/\ \ \ \/  / __ \/\  __\
      \ \ \\ \ /\ \_\ \ \ \_\ \/\ \_\ \ \ \_    \ \ \_/\ \/\ \_\ \/\ \/\ \ \ \ \ \_/\ \_\ \ \ \/ 
       \ \_\ \_\ \____/\ \____/\ \____/\ \__\    \ \_\\ \_\ \____/\ \_\ \_\ \_\ \__\ \____/\ \_\ 
        \/_/\/ /\/___/  \/___/  \/___/  \/__/     \/_/ \/_/\/___/  \/_/\/_/\/_/\/__/\/___/  \/_/ `)
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
