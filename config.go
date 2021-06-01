package main

import (
	"fmt"

	"github.com/spf13/viper"
)

var (
	c = viper.New()
)

func Configer(config string) *viper.Viper {
	fmt.Println(config)
	if config != "" {
		c.SetConfigFile(config) // 指定配置文件路径
	} else {
		c.SetConfigName("config")
		c.SetConfigType("toml")
		c.AddConfigPath(".")
	}
	err := c.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	return c
}
