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

	datapack "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/DataPack"

	serialhandle "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/SerialHandle"
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
	b, _ := json.Marshal(datapack.CurrentVariables)
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

func variableAddWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newVariable datapack.VariableT
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &newVariable) == nil {
		for _, v := range datapack.CurrentVariables.Variables {
			if v.Addr == newVariable.Addr {
				io.WriteString(w, "{\"status\":23}")
				return
			}
		}
		if serialhandle.SerialSendCmd(datapack.ACT_READ, newVariable) != nil {
			io.WriteString(w, "{\"status\":22}")
		} else {
			datapack.CurrentVariables.Variables = append(datapack.CurrentVariables.Variables, newVariable)
			io.WriteString(w, "{\"status\":0}")
		}
	} else {
		io.WriteString(w, "{\"status\":21}")
	}
}

func variableDelWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var oldVariable datapack.VariableT
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &oldVariable) == nil {
		for i, v := range datapack.CurrentVariables.Variables {
			if v.Addr == oldVariable.Addr {
				if serialhandle.SerialSendCmd(datapack.ACT_UNREAD, oldVariable) != nil {
					io.WriteString(w, "{\"status\":22}")
				} else {
					datapack.CurrentVariables.Variables = append(datapack.CurrentVariables.Variables[:i], datapack.CurrentVariables.Variables[i+1:]...)
					io.WriteString(w, "{\"status\":0}")
				}
				return
			}
		}
		io.WriteString(w, "{\"status\":24}")
		return
	}
	io.WriteString(w, "{\"status\":21}")
}

func variableModWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var modVariable datapack.VariableT
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &modVariable) == nil {
		if serialhandle.SerialSendCmd(datapack.ACT_WRITE, modVariable) != nil {
			io.WriteString(w, "{\"status\":22}")
		} else {
			io.WriteString(w, "{\"status\":0}")
		}
	} else {
		io.WriteString(w, "{\"status\":21}")
	}
}

func variableModAddWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newVariable datapack.VariableT
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &newVariable) == nil {
		for _, v := range datapack.ModVariables.Variables {
			if v.Addr == newVariable.Addr {
				io.WriteString(w, "{\"status\":23}")
				return
			}
		}
		datapack.ModVariables.Variables = append(datapack.ModVariables.Variables, newVariable)
		io.WriteString(w, "{\"status\":0}")
	} else {
		io.WriteString(w, "{\"status\":21}")
	}
}

func variableModListWebHandler(w http.ResponseWriter, _ *http.Request) {
	b, _ := json.Marshal(datapack.ModVariables)
	io.WriteString(w, string(b))
}

func Start() {
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
	http.HandleFunc("/variable/add", variableAddWebHandler)
	http.HandleFunc("/variable/del", variableDelWebHandler)
	http.HandleFunc("/variable/mod", variableModWebHandler)
	http.HandleFunc("/variable/modadd", variableModAddWebHandler)
	http.HandleFunc("/variable/modlist", variableModListWebHandler)
	http.HandleFunc("/ws", WebSocketHandler)
	addr := ":8080"
	log.Println("Listen on " + addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
