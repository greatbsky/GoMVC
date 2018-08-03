package utils

import (
	"github.com/BurntSushi/toml"
)

type tomlType struct {}

var Toml tomlType

func (this tomlType) Load(f string, v interface{}) error {
	_, err := toml.DecodeFile(f, v)
	return err
}