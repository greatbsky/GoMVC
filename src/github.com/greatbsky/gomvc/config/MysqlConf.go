package config

import (
	"github.com/BurntSushi/toml"
	"github.com/gobasis/log"
	"github.com/greatbsky/gomvc/global"
)

/*
Description:
An instance of mysqlConfType, with data from mysql.conf
 * Author: architect.bian
 * Date: 2018/10/26 16:36
 */
var MysqlConf mysqlConfType

type mysqlConfType struct {
	Conn string `toml:"conn"`
	DriverName string `toml:"driverName"`
}

/*
Description:
initialize parse conf file before application start up
 * Author: architect.bian
 * Date: 2018/10/23 13:38
*/
func init() {
	_, err := toml.DecodeFile(global.EnvConf.Root + "/conf/mysql.conf", &MysqlConf)
	if err != nil {
		log.Fatal("application conf file parse failed", "error", err)
	}
}
