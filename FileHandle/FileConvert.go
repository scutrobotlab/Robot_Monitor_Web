package filehandle

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

type projectVariablesT struct {
	Addr string
	Size string
	Name string
	Type string
}

type jsonProjectVariables struct {
	Variables []projectVariablesT
}

var ProjectVariables jsonProjectVariables

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
		ProjectVariables.Variables = append(ProjectVariables.Variables, projectVariablesT{Addr: v[1], Size: v[2], Name: v[3], Type: v[5]})
	}
	jsonTxt, err := json.Marshal(ProjectVariables)
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
