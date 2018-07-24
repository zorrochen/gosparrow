package config

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/configor"
)

type Config struct {
	ServerIP   string `yaml:"server_ip"`
	ServerPort string `yaml:"server_port"`

	EnableErrorLogMonitor bool `yaml:"enable_error_log_monitor"`
}

var Cfg *Config

func Init(confPath string) {
	var path string
	path = confPath

	Cfg = &Config{}
	if err := configor.Load(Cfg, path); err != nil {
		panic(err)
	}

	cfgjson, _ := json.MarshalIndent(Cfg, "", "  ")
	fmt.Printf("[ConfigInit] %s\n", string(cfgjson))
}
