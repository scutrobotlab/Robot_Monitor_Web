package filehandle

import "os"

type ConfigT struct {
	IsSaveDataAddr        bool
	IsSaveVariablesToMod  bool
	IsSaveVariablesToRead bool
}

var Config ConfigT

func LoadConfig() {
	if _, err := os.Stat("Config.json"); os.IsNotExist(err) {
		Config.IsSaveDataAddr = true
		Config.IsSaveVariablesToMod = true
		Config.IsSaveVariablesToRead = true
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
	if !Config.IsSaveVariablesToMod {
		os.Remove("VariablesToMod.json")
	}
	if !Config.IsSaveVariablesToRead {
		os.Remove("VariablesToRead.json")
	}
}
