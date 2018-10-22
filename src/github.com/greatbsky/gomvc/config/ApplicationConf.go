package config

import "path/filepath"


var ApplicationConf ApplicationConfType

type ApplicationConfType struct {
	Title   string `toml:"title"`
	Base    string
	BaseCss string
	BaseJs  string
	BaseImg string
	BaseFile string
	BaseApi string
	UserImg string
	WebImg string
	Domain  string

	//API APIType
	Sys SysType

	initData map[interface{}]interface{}
}


type APIType struct {
	Login string
	Logout string
}

type SysType struct {
	RootDir     string
	DoPagesPath string
	LoginToken string
}

func init() {
	ApplicationConf.Sys.RootDir, _ = filepath.Abs("./")
}