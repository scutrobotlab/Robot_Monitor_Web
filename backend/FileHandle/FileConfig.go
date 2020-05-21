package filehandle

import "os"

type ConfigT struct {
	IsSaveDataAddr     bool
	IsSaveVariableModi bool
	IsSaveVariableRead bool
}

var Config ConfigT

func LoadConfig() {
	if _, err := os.Stat("Config.json"); os.IsNotExist(err) {
		Config.IsSaveDataAddr = true
		Config.IsSaveVariableModi = true
		Config.IsSaveVariableRead = true
	} else {
		jsonLoad("Config.json", &Config)
	}
}

func SaveConfig() {
	jsonSave("Config.json", Config)
	if !Config.IsSaveDataAddr {
		os.Remove("DataAddr.json")
	} else {
		jsonSave("DataAddr.json", ProjectVariables)
	}
	if !Config.IsSaveVariableModi {
		os.Remove("VariablesToMod.json")
	}
	if !Config.IsSaveVariableRead {
		os.Remove("VariablesToRead.json")
	}
}
