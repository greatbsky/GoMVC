package utils

import (
	"github.com/BurntSushi/toml"
)

type tomlType struct {}

var Toml tomlType

/*
Description: Load from f, then decode with toml type

 * Author: architect.bian
 * Date: 2018/08/06 16:19
 */
func (this tomlType) Load(f string, v interface{}) error {
	_, err := toml.DecodeFile(f, v)
	return err
}