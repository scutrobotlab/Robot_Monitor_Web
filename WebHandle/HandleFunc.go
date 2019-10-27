package WebHandle

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"../DataPack"

	"../SerialHandle"
)

func ListSerialPortsWebHandler(w http.ResponseWriter, _ *http.Request) {
	jsonPack := DataPack.JsonSerialPort{Ports: SerialHandle.FindSerialPorts()}
	b, _ := json.Marshal(jsonPack)
	io.WriteString(w, string(b))
}

func OpenSerialPortWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	baud, err := strconv.Atoi(strings.Join(r.Form["baud"], ""))
	if err != nil {
		baud = 115200
	}
	port := strings.Join(r.Form["port"], "")
	if port != "" {
		if SerialHandle.OpenSerialPort(port, baud) != nil {
			io.WriteString(w, "Fail: Cannot open serial port")
		} else {
			io.WriteString(w, "Success")
		}
	} else {
		io.WriteString(w, "Fail: empty port")
	}
}

func CloseSerialPortWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if SerialHandle.CloseSerialPort() != nil {
		io.WriteString(w, "Fail: Cannot close serial port")
	} else {
		io.WriteString(w, "Success")
	}
}

func VariableActWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var s SerialHandle.DataToSerial
	var web DataPack.DataFromWeb_t
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &web) == nil {
		s.GetDataFromWeb(&web)
		if s.Send() != nil {
			io.WriteString(w, "Fail: Cannot send data")
		} else {
			if s.Act == DataPack.ACT_READ {
				var t DataPack.DataToRead_t
				t.GetWebData(&web)
				DataPack.DataToRead = append(DataPack.DataToRead, t)
			} else if s.Act == DataPack.ACT_UNREAD {
				for i, v := range DataPack.DataToRead {
					if v.Addr == web.Addr {
						DataPack.DataToRead = append(DataPack.DataToRead[:i], DataPack.DataToRead[i+1:]...)
					}
				}
			}
			io.WriteString(w, "Success")
		}
	} else {
		io.WriteString(w, "Fail: Unsupported json format")
	}

}

func WebHandleStart() {
	jsonWS := make(chan string, 10)
	go SerialHandle.SerialParse(jsonWS)
	WebSocketHandler := MakeWebSocketHandler(jsonWS)
	http.Handle("/", http.FileServer(http.Dir("./WebPage/")))
	http.HandleFunc("/list", ListSerialPortsWebHandler)
	http.HandleFunc("/open", OpenSerialPortWebHandler)
	http.HandleFunc("/close", CloseSerialPortWebHandler)
	http.HandleFunc("/act", VariableActWebHandler)
	http.HandleFunc("/ws", WebSocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
