package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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
	defer serialhandle.CloseSerialPort()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			filehandle.SaveAll()
		}
	}()
	http.Handle("/", http.FileServer(http.Dir("./WebPage/")))
	webhandle.Reg()
	port := ""
	if len(os.Args) > 1 {
		p, _ := strconv.Atoi(os.Args[1])
		if p > 0 && p < 65535 {
			port = ":" + os.Args[1]
		} else {
			port = ":8080"
		}
	} else {
		port = ":8080"
	}
	log.Println("Listen on " + port)
	log.Println("Don't close this before you have done")
	var commands = map[string]string{
		"windows": "explorer.exe",
		"darwin":  "open",
		"linux":   "xdg-open",
	}
	run, ok := commands[runtime.GOOS]
	if !ok {
		log.Printf("don't know how to open things on %s platform", runtime.GOOS)
	} else {
		go func() {
			log.Println("Your browser will start in 3 seconds")
			time.Sleep(3 * time.Second)
			exec.Command(run, "http://localhost"+port).Start()
		}()
	}
	log.Fatal(http.ListenAndServe(port, nil))
}
