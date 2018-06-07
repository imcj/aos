package bootstrap

import (
	"fmt"

	aosContainer "github.com/aos-stack/aos/container"
	"github.com/aos-stack/env"
	"github.com/aos-stack/golib/consul"
	"github.com/spf13/viper"
)

// ConsulCommand ...
type ConsulCommand struct {
	ConsulURL         string
	ConsulurlListName string
}

// Execute ...
func (c ConsulCommand) Execute() {
	c.ConsulURL = viper.GetString("consul.url")
	c.ConsulurlListName = viper.GetString("consul.list_name")
	if "" == c.ConsulURL || "" == c.ConsulurlListName {
		fmt.Println("consul.url 或者 consul.list_name 为空")
		return
	}
	envFix := viper.GetStringMapString("consul.env_fix")
	fix := ""
	if val, ok := envFix[env.GetLabel()]; ok {
		fix = val
	}
	c.ConsulURL = fix + c.ConsulURL
	config, err := consul.GetConf(c.ConsulurlListName, c.ConsulURL)
	if err != nil {
		fmt.Println(err)
	}
	if config == nil {
		fmt.Println("值为空")
	}
	aosContainer.ContainerSet("consul", config)
}

// getEnv 获取 consul address fix
func getEnv(dot string) string {

	consuls := map[string]string{
		"local":      "test.",
		"dev":        "test.",
		"test":       "t.",
		"prepare":    "pre.",
		"production": "pro.",
	}
	return consuls[env.GetLabel()]
}
