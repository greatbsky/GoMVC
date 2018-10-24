package routers

import (
	"fmt"
	"github.com/gobasis/http"
)

func index(ctx *http.Context) interface{} {
	ctx.Input.Data["title"] = "hi world from index"
	return "index"
}

func getJson(_ *http.Context) interface{} {
	//m := make(map[string]interface{})
	//m["k1"] = 123
	//m["k2"] = "v2"
	//m["k3"] = []string{"a", "b", "c"}
	//return m
	return http.ServerConf
}

func header(ctx *http.Context) interface{} {
	ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With, X-MY-API-Version")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
	ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
	ctx.ResponseWriter.WriteHeader(500)
	return "500"
}

func hello(ctx *http.Context) interface{} {
	//time.Sleep(10 * time.Second)
	fmt.Fprintf(ctx.ResponseWriter, "hello, %s!\n", ctx.Input.PathParams.ByName("name"))
	return nil
}

/*
Description:
initialize before application start up
 * Author: architect.bian
 * Date: 2018/10/23 13:39
 */
func init()  {
	http.Router.InitGet("/", index)
	http.Router.InitGet("/json", getJson)
	http.Router.InitGet("/header", header)
	http.Router.InitGet("/hello/:name", hello)
}


