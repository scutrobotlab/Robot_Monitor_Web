package WebHandle

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func MakeWebSocketHandler(jsonString chan string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		c, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Print("upgrade:", err)
			return
		}
		defer c.Close()
		for {
			b := <-jsonString
			err = c.WriteMessage(websocket.TextMessage, []byte(b))
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
	}
}
