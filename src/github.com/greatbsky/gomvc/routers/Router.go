package routers

import (
	"github.com/julienschmidt/httprouter"
)

var router = httprouter.New()

func GetRouter() *httprouter.Router {
	return router
}