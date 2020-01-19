package webhandle

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"strconv"
	"strings"

	"../datapack"

	"../serialhandle"
)

func currentSerialPortWebHandler(w http.ResponseWriter, _ *http.Request) {
	b, _ := json.Marshal(serialhandle.CurrentSerialPort)
	io.WriteString(w, string(b))
}

func listSerialPortsWebHandler(w http.ResponseWriter, _ *http.Request) {
	type jsonSerialPort struct {
		Ports []string
	}
	jsonPack := jsonSerialPort{Ports: serialhandle.FindSerialPorts()}
	b, _ := json.Marshal(jsonPack)
	io.WriteString(w, string(b))
}

func openSerialPortWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	baud, err := strconv.Atoi(strings.Join(r.Form["baud"], ""))
	if err != nil {
		baud = 115200
	}
	port := strings.Join(r.Form["port"], "")
	if port != "" {
		if serialhandle.OpenSerialPort(port, baud) != nil {
			io.WriteString(w, "{\"status\":11}")
		} else {
			io.WriteString(w, "{\"status\":0}")
			serialhandle.CurrentSerialPort.Name = port
			serialhandle.CurrentSerialPort.BaudRate = baud
		}
	} else {
		io.WriteString(w, "{\"status\":1}")
	}
}

func closeSerialPortWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	if serialhandle.CurrentSerialPort.Name != "" {
		if serialhandle.CloseSerialPort() != nil {
			io.WriteString(w, "{\"status\":13}")
		} else {
			io.WriteString(w, "{\"status\":0}")
			serialhandle.CurrentSerialPort.Name = ""
			serialhandle.CurrentSerialPort.BaudRate = 0
		}
	} else {
		io.WriteString(w, "{\"status\":12}")
	}
}

func currentVariablesWebHandler(w http.ResponseWriter, _ *http.Request) {
	b, _ := json.Marshal(datapack.DataToRead)
	io.WriteString(w, string(b))
}

func variableTypesWebHandler(w http.ResponseWriter, _ *http.Request) {
	type jsonTypes struct {
		Types []string
	}
	var types jsonTypes
	for k := range datapack.TypeLen {
		types.Types = append(types.Types, k)
	}
	sort.Strings(types.Types)
	b, _ := json.Marshal(types)
	io.WriteString(w, string(b))
}

func variableOptWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var s serialhandle.DataToSerial
	var web datapack.DataFromWeb_t
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &web) == nil {
		s.GetDataFromWeb(&web)
		if s.Send() != nil {
			io.WriteString(w, "Fail: Cannot send data")
		} else {
			if s.Act == datapack.ACT_READ {
				var t datapack.DataToRead_t
				t.GetWebData(&web)
				datapack.DataToRead.Variables = append(datapack.DataToRead.Variables, t)
			} else if s.Act == datapack.ACT_UNREAD {
				for i, v := range datapack.DataToRead.Variables {
					if v.Addr == web.Addr {
						datapack.DataToRead.Variables = append(datapack.DataToRead.Variables[:i], datapack.DataToRead.Variables[i+1:]...)
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
	go serialhandle.SerialParse(jsonWS)
	WebSocketHandler := makeWebSocketHandler(jsonWS)
	http.Handle("/", http.FileServer(http.Dir("./WebPage/")))
	http.HandleFunc("/serial", currentSerialPortWebHandler)
	http.HandleFunc("/serial/list", listSerialPortsWebHandler)
	http.HandleFunc("/serial/open", openSerialPortWebHandler)
	http.HandleFunc("/serial/close", closeSerialPortWebHandler)
	http.HandleFunc("/variable", currentVariablesWebHandler)
	http.HandleFunc("/variable/types", variableTypesWebHandler)
	http.HandleFunc("/variable/opt", variableOptWebHandler)
	http.HandleFunc("/ws", WebSocketHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
