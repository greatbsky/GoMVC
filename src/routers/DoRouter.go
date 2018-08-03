package routers

import (
	"../base"
	"github.com/astaxie/beego"
	"../global"
	"../utils"
	"../hp"
)

type DoRouter struct {
	base.Router
}

func (this *DoRouter) Action() {
	//adminuid from cookie
	channel := this.Ctx.Input.Param(":channel")
	table := this.Ctx.Input.Param(":table")
	titleQ := this.Ctx.Input.Query("title")
	if len(titleQ) > 0 {
		this.Ctx.Input.SetData("title", titleQ)
	}
	channelQ := this.Ctx.Input.Query("channel")
	if len(channelQ) > 0 {
		this.Ctx.Input.SetData("channel", channelQ)
	} else {
		this.Ctx.Input.SetData("channel", channel)
	}
	tableQ := this.Ctx.Input.Query("table")
	if len(tableQ) > 0 {
		this.Ctx.Input.SetData("table", tableQ)
	} else {
		this.Ctx.Input.SetData("table", table)
	}
	doPagesPath := global.Conf.DoPagesPath(channel + table)
	if utils.File.Exist(doPagesPath) {
		content := utils.Template.Execute(doPagesPath, this.Ctx.Input.Data())
		this.Data["head"] = hp.Do.ParseHead(content)
		this.Data["body"] = hp.Do.ParseBody(content, this.Ctx.Input.Data())
		this.Show("doTemplate.tpl")
	} else {
		this.Abort("404")
	}
}

func init() {
	beego.Router("/do/:table([\\w]+)", &DoRouter{}, "get:Action")
	beego.Router("/do/:channel([\\w]+)/:table([\\w]+)", &DoRouter{}, "get:Action")
}
