package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type plugin struct {
	pluginname string
	pluginpath string
	pluginlog  string
}
type Config struct {
	logfile string
	plugins []plugin
}

func main() {

}

func readconfig(cfg *Config, confdir string, confname string) bool {
	viper.SetConfigName(confname)
	viper.AddConfigPath(confdir)
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Printf("Config file not found...%s\n", err.Error())
		return false
	}

	err := viper.Unmarshal("Groucho", &cfg)

}
