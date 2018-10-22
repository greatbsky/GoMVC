package routers

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"fmt"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func init()  {
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
}

