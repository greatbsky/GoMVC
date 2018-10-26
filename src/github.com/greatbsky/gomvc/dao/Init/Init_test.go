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
	//rows, err := db.Query("SELECT 1 as first, 2 as second;")
	//if err != nil {
	//	log.Println(err)
	//}
	//defer rows.Close()
	////var a string
	////var name string
	////for rows.Next() {
	////	err := rows.Scan(&a, &name)
	////	if err != nil {
	////		log.Fatal(err)
	////	}
	////	log.Println(name)
	////}
	//// Get column names
	//columns, err := rows.Columns()
	//if err != nil {
	//	panic(err.Error())
	//}
	//fmt.Println(columns)

}

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

func formatSql(content []byte) []string {
	content = regexp.MustCompile(`\/\*[\s\S]*?\*\/`).ReplaceAll(content, []byte{})
	content = regexp.MustCompile(`\r`).ReplaceAll(content, []byte(""))
	content = regexp.MustCompile(`^[^\w+]*`).ReplaceAll(content, []byte(" "))
	content = regexp.MustCompile("\uFEFF").ReplaceAll(content, []byte(""))
	return strings.Split(string(content), ";")
}