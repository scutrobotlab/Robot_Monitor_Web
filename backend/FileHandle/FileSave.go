package filehandle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	datapack "Robot_Monitor_Web/backend/DataPack"
)

func jsonLoad(filename string, v interface{}) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		log.Println(filename, "Found")
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			log.Println(err)
		}
		err = json.Unmarshal(data, v)
		if err != nil {
			log.Println(err)
		}
	}
}

func jsonSave(filename string, v interface{}) {
	jsonTxt, err := json.Marshal(v)
	if err != nil {
		log.Println(err)
	}
	err = ioutil.WriteFile(filename, jsonTxt, 0644)
	if err != nil {
		log.Println(err)
	}
}

func SaveAll() {
	SaveConfig()
	if Config.IsSaveVariableModi {
		jsonSave("VariablesToMod.json", datapack.VariableModi)
	}
	if Config.IsSaveVariableRead {
		jsonSave("VariablesToRead.json", datapack.VariableRead)
	}
}

func LoadSaves() {
	LoadConfig()
	if Config.IsSaveDataAddr {
		jsonLoad("DataAddr.json", &ProjectVariables)
	}
	if Config.IsSaveVariableModi {
		jsonLoad("VariablesToMod.json", &datapack.VariableModi)
	}
	if Config.IsSaveVariableRead {
		jsonLoad("VariablesToRead.json", &datapack.VariableRead)
	}
}
