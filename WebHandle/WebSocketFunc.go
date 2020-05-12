package webhandle

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var chOn = make(chan int)
var chOff = make(chan int)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func makeWebSocketHandler(jsonString chan string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			<-chOn
		Loop:
			for {
				select {
				case <-chOff:
					break Loop
				default:
					b := <-jsonString
					err = c.WriteMessage(websocket.TextMessage, []byte(b))
					if err != nil {
						log.Println("write:", err)
						break
					}
				}
			}
		}
	}
}

func webSocketOnHandler(w http.ResponseWriter, _ *http.Request) {
	chOn <- 1
	io.WriteString(w, "{\"status\":0}")
}

func webSocketOffHandler(w http.ResponseWriter, _ *http.Request) {
	chOff <- 1
	io.WriteString(w, "{\"status\":0}")
}
