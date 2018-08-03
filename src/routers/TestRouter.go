package routers

import (
	"../base"
	"github.com/astaxie/beego"
	"strconv"
)

type TestRouter struct {
	base.Router
}


func (this *TestRouter) List() {
	type JType struct {
		A string
		B string
		C string
		D string
		E string
		F string
		G string
		H string
	}
	data := make(map[string]interface{})
	data["total"] = 10
	list := make([]*JType, 0)
	for i := 0; i < 10 ; i++ {
		u := JType{"a" + strconv.Itoa(i), "b" + strconv.Itoa(i), "c" + strconv.Itoa(i), "d" + strconv.Itoa(i), "e" + strconv.Itoa(i), "f" + strconv.Itoa(i), "g" + strconv.Itoa(i), "h" + strconv.Itoa(i)}
		list = append(list, &u)
	}
	data["rows"] = list
	this.Data["json"] = data
	this.ServeJSON()
}


func (this *TestRouter) CUD() {
	data := make(map[string]interface{})
	data["success"] = true
	this.Data["json"] = data
	this.ServeJSON()

}

func (this *TestRouter) Login() {
	this.Redirect("/admin/index", 302)
}

func init() {
	beego.Router("/api/:table([\\w]+)/list", &TestRouter{}, "get:List")
	beego.Router("/api/:channel([\\w]+)/:table([\\w]+)/list", &TestRouter{}, "get:List")
	beego.Router("/api/:table([\\w]+)/add", &TestRouter{}, "post:CUD")
	beego.Router("/api/:table([\\w]+)/delete", &TestRouter{}, "get:CUD")
	beego.Router("/api/:table([\\w]+)/edit", &TestRouter{}, "post:CUD")
	beego.Router("/api/:table([\\w]+)/status/:status([\\d]+)", &TestRouter{}, "post:CUD")
	beego.Router("/api/admin/login", &TestRouter{}, "post:Login")
}
