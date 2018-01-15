package routers

import (
	"../base"
	"github.com/astaxie/beego"
)

type MainRouter struct {
	base.Router
}

func (this *MainRouter) Index() {
	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.Show("index.tpl")
}

func (this *MainRouter) Login() {
	// this.Ctx.Output.Body([]byte(this.GetString("name")))
	this.Ctx.Output.Body([]byte(this.Ctx.Input.Query("name")))
}

func (this *MainRouter) Json() {
	type User struct {
		Name string
		Age  int
	}
	u := User{"man", 40}
	this.Data["json"] = &u
	this.ServeJSON()
}

func init() {
	beego.Router("/", &MainRouter{}, "*:Index")
	beego.Router("/login", &MainRouter{}, "*:Login")
	beego.Router("/json", &MainRouter{}, "*:Json")
}

//
// beego.Router("/api/create",&RestRouter{},"post:CreateFood")
// beego.Router("/api/update",&RestRouter{},"put:UpdateFood")
// beego.Router("/api/delete",&RestRouter{},"delete:DeleteFood")
// beego.Router("/api",&RestRouter{},"get,post:ApiFunc")
// beego.Router("/simple",&SimpleRouter{},"get:GetFunc;post:PostFunc")
