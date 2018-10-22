package routers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var router = httprouter.New()

func GetRouter() *httprouter.Router {
	return router
}

/*
Description:
mapping a handler with get method
 * Author: architect.bian
 * Date: 2018/10/22 20:27
 */
func initGet(path string, handle httprouter.Handle) {
	router.GET(path, proxyFun(handle))
}

/*
Description:
mapping a handler with HEAD method
 * Author: architect.bian
 * Date: 2018/10/22 20:31
 */
func initHEAD(path string, handle httprouter.Handle) {
	router.GET(path, proxyFun(handle))
}

/*
Description:
mapping a handler with OPTIONS method
 * Author: architect.bian
 * Date: 2018/10/22 20:31
 */
func initOPTIONS(path string, handle httprouter.Handle) {
	router.OPTIONS(path, proxyFun(handle))
}

/*
Description:
mapping a handler with POST method
 * Author: architect.bian
 * Date: 2018/10/22 20:31
 */
func initPOST(path string, handle httprouter.Handle) {
	router.POST(path, proxyFun(handle))
}

/*
Description:
mapping a handler with PUT method
 * Author: architect.bian
 * Date: 2018/10/22 20:31
 */
func initPUT(path string, handle httprouter.Handle) {
	router.PUT(path, proxyFun(handle))
}

/*
Description:
mapping a handler with PATCH method
 * Author: architect.bian
 * Date: 2018/10/22 20:31
 */
func initPATCH(path string, handle httprouter.Handle) {
	router.PATCH(path, proxyFun(handle))
}

/*
Description:
mapping a handler with DELETE method
 * Author: architect.bian
 * Date: 2018/10/22 20:31
 */
func initDELETE(path string, handle httprouter.Handle) {
	router.DELETE(path, proxyFun(handle))
}

/*
Description:
mapping a handler with special method
 * Author: architect.bian
 * Date: 2018/10/22 20:31
 */
func initHandle(method, path string, handle httprouter.Handle) {
	router.Handle(method, path, proxyFun(handle))
}

func proxyFun(handle httprouter.Handle) httprouter.Handle {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		handle(w, r, ps)
	}
}