package config

import (
	"github.com/BurntSushi/toml"
	"github.com/gobasis/log"
	"github.com/greatbsky/gomvc/global"
)

var AppConf appConfType

type appConfType struct {
	Title string `toml:"title"`
	API apiType `toml:"api"`
}

type apiType struct {
	BTCAPI string `toml:"btcAPI"`
}

/*
Description:
initialize parse conf file before application start up
 * Author: architect.bian
 * Date: 2018/10/23 13:38
*/
func init() {
	_, err := toml.DecodeFile(global.EnvConf.Root + "/conf/application.conf", &AppConf)
	if err != nil {
		log.Fatal("application conf file parse failed", "error", err)
	}
}
