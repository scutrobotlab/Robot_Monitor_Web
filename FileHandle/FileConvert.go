package filehandle

import (
	"io/ioutil"
	"log"
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
	content, err := ioutil.ReadFile("DataAddr")
	if err != nil {
		log.Println(err)
		return err
	}
	reg := regexp.MustCompile(`(0x[0-9a-f]{8})\s{2}(0x[0-9a-f]+)\s+((\*\s)?[a-zA-Z0-9_\.]+)\s+([a-zA-Z0-9_\.\s]+?)[\n|\r]`)
	match := reg.FindAllStringSubmatch(string(content), -1)
	ProjectVariables.Variables = nil
	for _, v := range match {
		ProjectVariables.Variables = append(ProjectVariables.Variables, projectVariablesT{Addr: v[1], Size: v[2], Name: v[3], Type: v[5]})
	}
	if Config.IsSaveDataAddr {
		jsonSave("DataAddr.json", ProjectVariables)
	}
	return nil
}
