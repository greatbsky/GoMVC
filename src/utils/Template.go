package utils

import (
	"strings"
	"html/template"
	"path/filepath"
	"io/ioutil"
	"bytes"
)

type templateType struct {
	FuncMap template.FuncMap
}

var Template = templateType{
	FuncMap: template.FuncMap{
		"title": strings.Title,
		"contains": contains,
		"equal": equal,
	},
}

/*
Description: 应用funcMap读取file做模板返回一个模板

 * Author: architect.bian
 * Date: 2018/08/06 16:17
 */
func (this templateType) ParseFiles(file string) *template.Template {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	s := string(buf)
	name := filepath.Base(file)
	return template.Must(template.New(name).Funcs(this.FuncMap).Parse(s))
}

/*
Description: parse file as template,return result

 * Author: architect.bian
 * Date: 2018/08/06 16:17
 */
func (this templateType) Execute(file string, data interface{}) string {
	var buf bytes.Buffer
	err := this.ParseFiles(file).Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

/*
Description: flag是否存在在s中

 * Author: architect.bian
 * Date: 2018/08/06 16:18
 */
func contains(s interface{}, flag string) bool {
	if s == nil {
		return false
	}
	return strings.Contains(s.(string), flag)
}

/*
Description: 判断字符串是否相等

 * Author: architect.bian
 * Date: 2018/08/06 16:19
 */
func equal(s interface{}, target string) bool {
	if s == nil {
		return false
	}
	return s == target
}
