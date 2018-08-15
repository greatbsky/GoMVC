package global

import (
	"../utils"
	"path/filepath"
)

var Conf ConfType

type ConfType struct {
	Title   string `toml:"title"`
	Base    string
	BaseCss string
	BaseJs  string
	BaseImg string
	BaseApi string
	UserImg string
	WebImg string
	Domain  string

	API APIType
	Sys SysType

	initData map[interface{}]interface{}
}

type APIType struct {
	Login string
	Logout string
}
type SysType struct {
	RootDir     string
	DoPagesPath string
	LoginToken string
}

/**
构造初始数据
 */
//func (this *ConfType) InitData() map[interface{}]interface{} {
//	if this.initData != nil {
//		return this.initData
//	} else {
//		data := make(map[interface{}]interface{})
//		data["title"] = this.Title
//		data["base"] = this.Base
//		data["basecss"] = this.BaseCss
//		data["basejs"] = this.BaseJs
//		data["baseimg"] = this.BaseImg
//		data["domain"] = this.Domain
//		data["api"] = this.API
//		this.initData = data
//		return this.initData
//	}
//}
func (this *ConfType) TagTplPath(tag string) string {
	return Conf.Sys.RootDir + "/views/tags/" + tag + ".tpl"
}
func (this *ConfType) GetDoPageHtmlPath(name string) string {
	return this.GetDoPagePath(name + ".html")
}
func (this *ConfType) GetDoPagePath(path string) string {
	return filepath.Join(Conf.Sys.DoPagesPath, path)
}

func init() {
	utils.Toml.Load("conf/config.conf", &Conf)
	Conf.Sys.RootDir, _ = filepath.Abs("./")
}
