package hp

import (
	"html/template"
	"../utils"
	"strings"
	"bytes"
	"os"
	"../global"
	"regexp"
	"fmt"
)

type doType struct {}

var Do doType

const (
	attrFragmentKey = "attrFragment"
	attrsKey = "attrs"
	isStartKey = "isStart"
	titleKey = "title"
	tagKey = "tag"
)

/**
返回参数s的head里的html内容
 */
func (this doType) ParseHead(s string) template.HTML {
	return template.HTML(utils.Xml.TagContent(s, "head"))
}

/**
返回参数s的body里的html内容
 */
func (this doType) ParseBody(html string, data map[interface{}]interface{}) template.HTML {
	data[titleKey] = utils.Xml.TagContent(html, "title")
	body := utils.Xml.TagContent(html, "body")
	body = parseTag(body, "search", data)
	body = parseTag(body, "datagrid", data)
	body = parseTag(body, "toolbar", data)
	body = parseTag(body, "tools", data)
	body = parseTag(body, "doaction", data)
	body = parseTag(body, "status", data)
	body = parseTag(body, "viewsection", data)
	body = parseTag(body, "addsection", data)
	body = parseTag(body, "editsection", data)
	return template.HTML(body)
}

func parseTag(s string, tag string, data map[interface{}]interface{}) (string) {
	pos := 0
	for {
		var eleFragment string
		var attrFragment string
		eleFragment, attrFragment, pos = utils.Xml.TagAtrrs(s, tag, pos)
		if len(eleFragment) == 0 {
			break
		}
		tmplPath := global.Conf.TagTplPath(tag)
		tmpl := utils.Template.ParseFiles(tmplPath) //获取模板

		var bufStart bytes.Buffer
		var bufEnd bytes.Buffer
		data[attrFragmentKey] = template.HTMLAttr(attrFragment)
		data[attrsKey] = getAttributes(attrFragment)
		data[isStartKey] = true
		data[tagKey] = tag

		err := tmpl.Execute(&bufStart, data) //解析标签开始部分
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		data[isStartKey] = false
		err = tmpl.Execute(&bufEnd, data) //解析标签结束部分
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		s = strings.Replace(s, eleFragment, bufStart.String(), -1) //替换开始部分
		s = strings.Replace(s, "</"+tag+">", bufEnd.String(), -1)  //替换结束部分
		delete(data, attrFragmentKey)
	}
	return s
}

/**
解析attrFragment返回map类型的attribute，包含默认值
 */
func getAttributes(attrFragment string) map[string]interface{} {
	fields := strings.Split(attrFragment, " ")
	attrs := make(map[string]interface{})
	for _, f := range fields {
		kv := strings.Split(f, "=")
		if len(kv) == 2 {
			var re = regexp.MustCompile(`(^["'])|(["']$)`)
			v := re.ReplaceAllString(kv[1], ``)
			if kv[0] == "id" {
				attrs["id4code"] = template.JS(v)
			}
			attrs[kv[0]] = v
		}
	}
	if attrs["datagrid"] == "" {
		attrs["datagrid"] = "#dg"
	}
	return attrs
}
