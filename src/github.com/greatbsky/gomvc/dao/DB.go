package dao

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"github.com/greatbsky/gomvc/config"
	"time"
)

var DB *gorm.DB

/*
Description: 
initialize DB instance
 * Author: architect.bian
 * Date: 2018/10/29 12:52
 */
func init() {
	var err error
	DB, err = gorm.Open(config.MysqlConf.DriverName, config.MysqlConf.Conn)
	if err != nil {
		log.Fatal(err)
	}
	DB.DB().SetMaxOpenConns(config.MysqlConf.MaxOpenConns)
	DB.DB().SetMaxIdleConns(config.MysqlConf.MaxIdleConns)
	DB.DB().SetConnMaxLifetime(time.Duration(config.MysqlConf.ConnMaxLifetime))
}
