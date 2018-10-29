package Init

import (
	"testing"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/greatbsky/gomvc/config"
	"github.com/greatbsky/gomvc/config"
	"database/sql"
	"log"
	"io/ioutil"
	"strings"
	"bytes"
	"regexp"
)

var dbNames = []string {
	"AdminDB",
}
var db *sql.DB
/*
Description:
A testor for init database schema & data
 * Author: architect.bian
 * Date: 2018/10/29 11:30
 */
func TestInit(t *testing.T) {
	var err error
	db, err = sql.Open(config.MysqlConf.DriverName, config.MysqlConf.Conn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	for _, name := range dbNames  {
		log.Println("begin to initialize Schema " + name)
		execSql("Schema", name) //init schema
		log.Println("begin to initialize data " + name)
		execSql("InitData", name) //init data
		log.Println("begin to initialize test Data " + name)
		execSql("InitData", name + "_Test") //init test data
	}
}

/*
Description:
exec sql match filename in dirname
 * Author: architect.bian
 * Date: 2018/10/29 11:29
 */
func execSql(dirName, fileName string) {
	var dir = "./DB/" + dirName
	var buffer bytes.Buffer
	sqlBytes, err := ioutil.ReadFile(fmt.Sprintf(dir + "/%s.sql", fileName))
	if err != nil {
		log.Fatal(err)
	}
	buffer.Write(sqlBytes)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	fiArr, err := ioutil.ReadDir(fmt.Sprintf(dir))
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fiArr {
		if strings.HasPrefix(fi.Name(), fmt.Sprintf("%s_Alt", fileName)) {
			bytes, err := ioutil.ReadFile(fmt.Sprintf(dir + "/%s", fi.Name()))
			if err != nil {
				log.Fatal(err)
			}
			buffer.Write(bytes)
		}
	}
	strs := formatSql(buffer.Bytes())
	for _, sql := range strs {
		sql = strings.TrimSpace(sql)
		if len(sql) == 0 {
			continue
		}
		_, err := db.Exec(sql)
		if err != nil {
			log.Fatal(err)
		}
	}
	tx.Commit()
}

/*
Description:
format content, extract sql
 * Author: architect.bian
 * Date: 2018/10/29 11:29
 */
func formatSql(content []byte) []string {
	content = regexp.MustCompile(`\/\*[\s\S]*?\*\/`).ReplaceAll(content, []byte{}) //replace comment /*xxx*/
	content = regexp.MustCompile(`\r`).ReplaceAll(content, []byte("")) //replace /r
	content = regexp.MustCompile(`^[^\w+]*`).ReplaceAll(content, []byte(" ")) //replace prefix space word
	content = regexp.MustCompile("\uFEFF").ReplaceAll(content, []byte("")) //replace BOM(Byte Order Mark)
	return strings.Split(string(content), ";")
}