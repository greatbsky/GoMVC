package base

import (
	"github.com/astaxie/beego"
	"fmt"
	"os"
)

type Router struct {
	beego.Controller
}

/*
Description: show tpl view

 * Author: architect.bian
 * Date: 2018/08/06 16:23
 */
func (this *Router) Show(tpl string) {
	this.TplName = tpl
}

/*
Description: return success json

 * Author: architect.bian
 * Date: 2018/08/06 16:21
 */
func (this *Router) Success() {
	data := make(map[string]interface{})
	data["success"] = true
	this.Data["json"] = data
	this.ServeJSON()
}

/*
Description: return 404 page

 * Author: architect.bian
 * Date: 2018/08/06 16:22
 */
func (this *Router) Error404() {
	this.TplName = "error/404.tpl"
}

/*
Description: return 501 page

 * Author: architect.bian
 * Date: 2018/08/06 16:22
 */
func (this *Router) Error501() {
	this.TplName = "error/501.tpl"
}

/*
Description: abort execute,print error to console

 * Author: architect.bian
 * Date: 2018/08/06 16:23
 */
func (this *Router) AbortErr(err error) {
	fmt.Fprintf(os.Stderr, "%v\n", err)
	this.Abort("501")
}
