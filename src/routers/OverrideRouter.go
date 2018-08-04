package routers

import (
	"../base"
	"github.com/astaxie/beego"
	"../global"
	"../utils"
	"html/template"
)

type OverrideRouter struct {
	base.Router
}

func (this *OverrideRouter) Override() {
	ext := this.Ctx.Input.Param(":ext")
	path := global.Conf.GetDoPagePath("override." + ext)
	this.Data["ext"] = ext
	if utils.File.Exist(path) {
		this.Data["body"] = template.HTML(utils.Template.Execute(path, this.Ctx.Input.Data()))
	}
	this.Show("override.tpl")
}

func init() {
	beego.Router("/override.:ext([\\w]+)", &OverrideRouter{}, "get:Override")
}
