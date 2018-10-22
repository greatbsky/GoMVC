package config

import (
	"github.com/BurntSushi/toml"
	"github.com/gobasis/log"
)

/*
Description: 
server config file
 * Author: architect.bian
 * Date: 2018/10/22 19:11
 */
type ServerConfType struct {
	Address string `toml:"address"`
	Port    int32
}

var ServerConf ServerConfType

func init() {
	_, err := toml.DecodeFile("conf/server.conf", &ServerConf)
	if err != nil {
		log.Fatal("", "error", err)
	}
}