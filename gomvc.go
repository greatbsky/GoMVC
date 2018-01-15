package gomvc

import (
	"fmt"

	"github.com/astaxie/beego"
)

type Router struct {
	beego.Controller
}

func (this *Router) Show(tpl string) {
	this.TplName = tpl
}

func Run() {
	beego.LoadAppConfig("ini", "../resources/conf/app.conf")
	fmt.Println("name ", beego.BConfig.AppName)
	fmt.Println("appname", beego.AppConfig.String("appname"))
	beego.BConfig.WebConfig.ViewsPath = "views"
	// beego.SetStaticPath("/css","css") // todo mapper *.html -> public/*.html  *.css->/public/*.css
	beego.Run()
}
