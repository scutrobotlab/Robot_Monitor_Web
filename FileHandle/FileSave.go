package filehandle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	datapack "www.scut-robotlab.cn/git/M3chD09/Robot_Monitor_Web/DataPack"
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
	if Config.IsSaveVariablesToMod {
		jsonSave("VariablesToMod.json", datapack.ModVariables)
	}
	if Config.IsSaveVariablesToRead {
		jsonSave("VariablesToRead.json", datapack.CurrentVariables)
	}
}

func LoadSaves() {
	LoadConfig()
	if Config.IsSaveDataAddr {
		jsonLoad("DataAddr.json", &ProjectVariables)
	}
	if Config.IsSaveVariablesToMod {
		jsonLoad("VariablesToMod.json", &datapack.ModVariables)
	}
	if Config.IsSaveVariablesToRead {
		jsonLoad("VariablesToRead.json", &datapack.CurrentVariables)
	}
}
