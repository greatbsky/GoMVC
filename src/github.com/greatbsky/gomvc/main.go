package main

import (
	"github.com/gobasis/log"
	"github.com/gobasis/log/zapimpl"
	_ "github.com/greatbsky/gomvc/routers"
	"github.com/gobasis/http"
)

func main() {
	log.UseLog(&zapimpl.Logger{}) // use zap log
	log.SetLevel(log.DevDebugLevel)

	http.Run()
}