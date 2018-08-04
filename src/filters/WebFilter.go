package filters

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
	"../global"
	"strings"
	"fmt"
	"os"
)

var WebFilter = func(ctx *context.Context) {
	initData(ctx)
	if strings.HasPrefix(ctx.Request.RequestURI, "/admin") || strings.HasPrefix(ctx.Request.RequestURI, "/do") {
		cookie, err := ctx.Request.Cookie(global.Conf.Sys.LoginToken)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
		if cookie == nil || len(cookie.Value) == 0 {
			fmt.Fprintf(os.Stderr, "requesting %v, but cookie is empty \n", ctx.Request.RequestURI)
			ctx.Redirect(302, fmt.Sprintf("%s/login", global.Conf.Base))
		}
	}
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
	ctx.Input.SetData("baseapi", global.Conf.BaseApi)
	ctx.Input.SetData("domain", global.Conf.Domain)
	ctx.Input.SetData("api", global.Conf.API)
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, WebFilter)
}
