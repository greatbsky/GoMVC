package base

import (
	"github.com/astaxie/beego"
	"fmt"
	"os"
)

type Router struct {
	beego.Controller
}

func (this *Router) Show(tpl string) {
	this.TplName = tpl
}

func (this *Router) Success() {
	data := make(map[string]interface{})
	data["success"] = true
	this.Data["json"] = data
	this.ServeJSON()
}

func (this *Router) Error404() {
	this.TplName = "error/404.tpl"
}

func (this *Router) Error501() {
	this.TplName = "error/501.tpl"
}

func (this *Router) AbortErr(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	this.Abort("501")
}
