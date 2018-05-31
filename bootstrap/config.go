package bootstrap

import (
	"encoding/json"
	"io/ioutil"
	"fmt"
)

type Application struct {
	Cycle []string
}

type BootstrapConfig struct {
	Consul string
	Application Application
}

type BootstrapConfigCommand struct {
	ConfigPath string
}

var Config BootstrapConfig

func (c BootstrapConfigCommand)Execute() {
	if "" == c.ConfigPath {
		c.ConfigPath = "app/bootstrap.json"
	}
	fmt.Println(c.ConfigPath)

	data, err := ioutil.ReadFile(c.ConfigPath)
	if nil != err {
		fmt.Println(err)
		fmt.Println("Load bootstrap config file failure.")
	}
	err = json.Unmarshal(data, &Config)
	if nil != err {
		fmt.Println("Load bootstrap config file error ")
	}
}