package app

import (
	_ "./filters"
	_ "./routers"
	"github.com/astaxie/beego"
	"./base"
	"./utils"
)

func Run() {
	beego.LoadAppConfig("ini", "conf/app.conf")
	// fmt.Println("name ", beego.BConfig.AppName)
	// fmt.Println("appname", beego.AppConfig.String("appname"))
	beego.BConfig.WebConfig.ViewsPath = "views"
	//beego.SetStaticPath("/css","css") // todo mapper *.html -> public/*.html  *.css->/public/*.css
	beego.ErrorController(&base.Router{})
	addFuncMap()
	beego.Run()
}

func addFuncMap() {
	for k, v := range utils.Template.FuncMap {
		beego.AddFuncMap(k, v)
	}
}