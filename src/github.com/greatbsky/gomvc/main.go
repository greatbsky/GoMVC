package main

import (
	"fmt"
	"github.com/gobasis/log"
	"github.com/gobasis/log/zapimpl"
	"github.com/greatbsky/gomvc/config"
	"net/http"
	"github.com/greatbsky/gomvc/routers"
	_ "github.com/greatbsky/gomvc/filters"
)

func main() {
	log.UseLog(&zapimpl.Logger{}) // use zap log
	log.SetLevel(log.DevDebugLevel)

	addr := fmt.Sprintf("%s:%d", config.ServerConf.Address, config.ServerConf.Port)
	log.Info("starting http server", "address", addr)
	routers.Router.GetRawRouter().NotFound = http.FileServer(http.Dir("public"))
	err := http.ListenAndServe(addr, routers.Router.GetRawRouter())
	if err != nil {
		log.Fatal("http server start failed", "error", err)
	}
}