package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"../global"
)

var WebFilter = func(ctx *context.Context) {
	initData(ctx)
}

/**
初始数据
 */
func initData(ctx *context.Context) {
	ctx.Input.SetData("title", global.Conf.Title)
	ctx.Input.SetData("base", global.Conf.Base)
	ctx.Input.SetData("basecss", global.Conf.BaseCss)
	ctx.Input.SetData("basejs", global.Conf.BaseJs)
	ctx.Input.SetData("baseimg", global.Conf.BaseImg)
	ctx.Input.SetData("domain", global.Conf.Domain)
	ctx.Input.SetData("api", global.Conf.API)
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, WebFilter)
}
