package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	analyse "logAnalysis/Service"
)

type Tconfig struct {
	Path     string
	Quantity int
}

func main() {
	config := Tconfig{}
	load("./run.conf", &config)
	res := analyse.Analyze_requests(config.Path, config.Quantity)
	fmt.Println("result:", res)
}
func load(path string, v interface{}) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		return
	}
}
