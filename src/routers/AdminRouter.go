package routers

import (
	"../base"
	"github.com/astaxie/beego"
	"../global"
	"../utils"
	"html/template"
)

type AdminRouter struct {
	base.Router
}

func (this *AdminRouter) Index() {
	this.Show("adminIndex.tpl")
}

func (this *AdminRouter) Menu() {
	doPagesPath := global.Conf.GetDoPageHtmlPath("menu")
	if utils.File.Exist(doPagesPath) {
		this.Data["body"] = template.HTML(utils.Template.Execute(doPagesPath, this.Ctx.Input.Data()))
		this.Show("adminMenu.tpl")
	} else {
		this.Abort("404")
	}
}

func (this *AdminRouter) Top() {
	this.Show("adminTop.tpl")
}

func (this *AdminRouter) Home() {
	this.Show("adminHome.tpl")
}

func init() {
	beego.Router("/admin/index", &AdminRouter{}, "get:Index")
	beego.Router("/admin/menu", &AdminRouter{}, "get:Menu")
	beego.Router("/admin/top", &AdminRouter{}, "get:Top")
	beego.Router("/admin/home", &AdminRouter{}, "get:Home")
}
