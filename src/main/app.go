package app

import (
	_ "./filters"
	_ "./routers"
	"github.com/greatbsky/gomvc"
)

func Run() {
	gomvc.Run()
}
