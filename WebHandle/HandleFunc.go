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

func CurrentSerialPortsWebHandler(w http.ResponseWriter, _ *http.Request) {
	b, _ := json.Marshal(SerialHandle.CurrentSerialPort)
	io.WriteString(w, string(b))
}

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
			io.WriteString(w, "{\"status\":11}")
		} else {
			io.WriteString(w, "{\"status\":0}")
			SerialHandle.CurrentSerialPort.Name = port
			SerialHandle.CurrentSerialPort.BaudRate = baud
		}
	} else {
		io.WriteString(w, "{\"status\":1}")
	}
}

func CloseSerialPortWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if SerialHandle.CurrentSerialPort.Name != "" {
		if SerialHandle.CloseSerialPort() != nil {
			io.WriteString(w, "{\"status\":13}")
		} else {
			io.WriteString(w, "{\"status\":0}")
			SerialHandle.CurrentSerialPort.Name = ""
			SerialHandle.CurrentSerialPort.BaudRate = 0
		}
	} else {
		io.WriteString(w, "{\"status\":12}")
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
	http.HandleFunc("/serial", CurrentSerialPortsWebHandler)
	http.HandleFunc("/serial/list", ListSerialPortsWebHandler)
	http.HandleFunc("/serial/open", OpenSerialPortWebHandler)
	http.HandleFunc("/serial/close", CloseSerialPortWebHandler)
	http.HandleFunc("/act", VariableActWebHandler)
	http.HandleFunc("/ws", WebSocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
