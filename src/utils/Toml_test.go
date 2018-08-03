package utils

import (
	"testing"
	"fmt"
)


type TestType struct {
	Title string `toml:"title"`
}

var conf TestType

func TestConfig(t *testing.T) {
	Toml.Load("../../conf/config.conf", &conf)
	fmt.Println(conf.Title)
}