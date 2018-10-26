package global

import (
	"github.com/gobasis/utils"
	"github.com/BurntSushi/toml"
	"github.com/gobasis/log"
	"path/filepath"
)

/*
Description:
An instance of envConfType with initialized data
 * Author: architect.bian
 * Date: 2018/10/26 16:58
 */
var EnvConf envConfType

type envConfType struct {
	Root string `toml:"root"`
}

/*
Description:
initialize global environment configuration before application start up
 * Author: architect.bian
 * Date: 2018/10/26 16:57
 */
func init() {
	EnvConf.Root, _ = filepath.Abs("./")
	if utils.File.Exist(EnvConf.Root + "/conf/env.conf") {
		_, err := toml.DecodeFile(EnvConf.Root + "/conf/env.conf", &EnvConf)
		if err != nil {
			log.Fatal("application conf file parse failed", "error", err)
		}
	}
	EnvConf.Root, _ = filepath.Abs(EnvConf.Root)
}