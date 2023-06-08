package ooo

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

func GetYamlMap[V any](sourcefile string, genericType V) map[string]V {
	yfile, err := ioutil.ReadFile(sourcefile)
	if err != nil {
		fmt.Println(err.Error())
	}
	yml := make(map[string]V)
	err2 := yaml.Unmarshal(yfile, &yml)
	if err2 != nil {
		fmt.Println(err2.Error())
	}
	return yml
}
