package webhandle

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"

	datapack "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/DataPack"
	filehandle "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/FileHandle"

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

func variableReadListWebHandler(w http.ResponseWriter, _ *http.Request) {
	b, _ := json.Marshal(datapack.VariableRead)
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

func variableReadAddWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newVariable datapack.VariableT
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &newVariable) == nil {
		for _, v := range datapack.VariableRead.Variables {
			if v.Addr == newVariable.Addr {
				io.WriteString(w, "{\"status\":23}")
				return
			}
		}
		if serialhandle.SerialSendCmd(datapack.ACT_SUBSCRIBE, newVariable) != nil {
			io.WriteString(w, "{\"status\":22}")
		} else {
			datapack.VariableRead.Variables = append(datapack.VariableRead.Variables, newVariable)
			io.WriteString(w, "{\"status\":0}")
		}
	} else {
		io.WriteString(w, "{\"status\":21}")
	}
}

func variableReadDelWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var oldVariable datapack.VariableT
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &oldVariable) == nil {
		for i, v := range datapack.VariableRead.Variables {
			if v.Addr == oldVariable.Addr {
				if serialhandle.SerialSendCmd(datapack.ACT_UNSUBSCRIBE, oldVariable) != nil {
					io.WriteString(w, "{\"status\":22}")
				} else {
					datapack.VariableRead.Variables = append(datapack.VariableRead.Variables[:i], datapack.VariableRead.Variables[i+1:]...)
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

func variableModiModWebHandler(w http.ResponseWriter, r *http.Request) {
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

func variableModiAddWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var newVariable datapack.VariableT
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &newVariable) == nil {
		for _, v := range datapack.VariableModi.Variables {
			if v.Addr == newVariable.Addr {
				io.WriteString(w, "{\"status\":23}")
				return
			}
		}
		datapack.VariableModi.Variables = append(datapack.VariableModi.Variables, newVariable)
		io.WriteString(w, "{\"status\":0}")
	} else {
		io.WriteString(w, "{\"status\":21}")
	}
}

func variableModiDelWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var oldVariable datapack.VariableT
	postData, _ := ioutil.ReadAll(r.Body)
	if json.Unmarshal(postData, &oldVariable) == nil {
		for i, v := range datapack.VariableModi.Variables {
			if v.Addr == oldVariable.Addr {
				datapack.VariableModi.Variables = append(datapack.VariableModi.Variables[:i], datapack.VariableModi.Variables[i+1:]...)
				io.WriteString(w, "{\"status\":0}")
				return
			}
		}
		io.WriteString(w, "{\"status\":24}")
		return
	}
	io.WriteString(w, "{\"status\":21}")
}

func variableModiListWebHandler(w http.ResponseWriter, _ *http.Request) {
	b, _ := json.Marshal(datapack.VariableModi)
	io.WriteString(w, string(b))
}

func fileUploadWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	defer os.Remove("DataAddr")
	r.ParseMultipartForm(32 << 20)
	file, _, err := r.FormFile("file")
	if err != nil {
		io.WriteString(w, "{\"status\":31}")
		return
	}
	defer file.Close()
	f, err := os.OpenFile("DataAddr", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		io.WriteString(w, "{\"status\":32}")
		return
	}
	defer f.Close()
	io.Copy(f, file)
	err = filehandle.Txt2json()
	if err != nil {
		io.WriteString(w, "{\"status\":33}")
		return
	}
	io.WriteString(w, "{\"status\":0}")
}

func fileVariablesWebHandler(w http.ResponseWriter, _ *http.Request) {
	b, _ := json.Marshal(filehandle.ProjectVariables)
	io.WriteString(w, string(b))
}

func fileConfigWebHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	r.ParseForm()
	if len(r.Form) == 0 {
		b, _ := json.Marshal(filehandle.Config)
		io.WriteString(w, string(b))
	} else {
		if _, ok := r.Form["sda"]; ok {
			filehandle.Config.IsSaveDataAddr, _ = strconv.ParseBool(strings.Join(r.Form["sda"], ""))
		}
		if _, ok := r.Form["svm"]; ok {
			filehandle.Config.IsSaveVariableModi, _ = strconv.ParseBool(strings.Join(r.Form["svm"], ""))
		}
		if _, ok := r.Form["svr"]; ok {
			filehandle.Config.IsSaveVariableRead, _ = strconv.ParseBool(strings.Join(r.Form["svr"], ""))
		}
		filehandle.SaveConfig()
		io.WriteString(w, "{\"status\":0}")
	}

}

func Reg() {
	rxBuff := make(chan []byte, 100)
	jsonWS := make(chan string, 10)
	go serialhandle.SerialTransmitThread()
	go serialhandle.SerialReceiveThread(rxBuff)
	go serialhandle.SerialPraseThread(rxBuff, jsonWS)
	webSocketHandler := makeWebSocketHandler(jsonWS)
	http.HandleFunc("/serial", currentSerialPortWebHandler)
	http.HandleFunc("/serial/list", listSerialPortsWebHandler)
	http.HandleFunc("/serial/open", openSerialPortWebHandler)
	http.HandleFunc("/serial/close", closeSerialPortWebHandler)
	http.HandleFunc("/variable/types", variableTypesWebHandler)
	http.HandleFunc("/variable-read/list", variableReadListWebHandler)
	http.HandleFunc("/variable-read/add", variableReadAddWebHandler)
	http.HandleFunc("/variable-read/del", variableReadDelWebHandler)
	http.HandleFunc("/variable-modi/list", variableModiListWebHandler)
	http.HandleFunc("/variable-modi/add", variableModiAddWebHandler)
	http.HandleFunc("/variable-modi/del", variableModiDelWebHandler)
	http.HandleFunc("/variable-modi/mod", variableModiModWebHandler)
	http.HandleFunc("/file/upload", fileUploadWebHandler)
	http.HandleFunc("/file/variables", fileVariablesWebHandler)
	http.HandleFunc("/file/config", fileConfigWebHandler)
	http.HandleFunc("/ws", webSocketHandler)
}
