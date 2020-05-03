package filehandle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type VariableT struct {
	Addr string
	Size string
	Name string
	Type string
}

type jsonSaveVariables struct {
	Variables []VariableT
}

var SaveVariables jsonSaveVariables

func Txt2json() error {
	file, err := os.Open("DataAddr")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
		return err
	}
	reg := regexp.MustCompile(`(0x[0-9a-f]{8})\s{2}(0x[0-9a-f]+)\s+((\*\s)?[a-zA-Z0-9_\.]+)\s+([a-zA-Z0-9_\.\s]+?)[\n|\r]`)
	match := reg.FindAllStringSubmatch(string(content), -1)
	for _, v := range match {
		SaveVariables.Variables = append(SaveVariables.Variables, VariableT{Addr: v[1], Size: v[2], Name: v[3], Type: v[5]})
	}
	jsonTxt, err := json.Marshal(SaveVariables)
	if err != nil {
		log.Fatal(err)
		return err
	}
	f, err := os.OpenFile("DataAddr.json", os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}
	f.Write(jsonTxt)
	return nil
}
