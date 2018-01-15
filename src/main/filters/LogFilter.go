package filters

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

var LogFilter = func(ctx *context.Context) {
	fmt.Println(ctx.Request.RequestURI)
}

func init() {
	beego.InsertFilter("/*", beego.BeforeRouter, LogFilter)
}
