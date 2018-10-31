package main

import (
	"github.com/gobasis/log"
	"github.com/gobasis/log/zapimpl"
	_ "github.com/greatbsky/gomvc/routers"
	_ "github.com/greatbsky/gomvc/config"
	"github.com/gobasis/http"
	"github.com/gobasis/utils"
)

func main() {
	log.UseLog(&zapimpl.Logger{}) // use zap log
	log.SetLevel(log.DevDebugLevel)
	utils.RunEnv.Setup("pprof", 120)
	go http.Run()
	defer log.Info("Shutdown complete")
	<-utils.RunEnv.NewInterruptChan()
}