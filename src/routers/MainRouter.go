package routers

import (
	"../base"
	"github.com/astaxie/beego"
	"../global"
)

type MainRouter struct {
	base.Router
}

func (this *MainRouter) Index() {
	this.Show("index.tpl")
}

func (this *MainRouter) Login() {
	this.Data["apiLogin"] = global.Conf.API.Login
	this.Data["apiLogout"] = global.Conf.API.Logout
	this.Show("login.tpl")
}

func init() {
	beego.Router("/", &MainRouter{}, "get:Index")
	beego.Router("/login", &MainRouter{}, "get:Login")
}
