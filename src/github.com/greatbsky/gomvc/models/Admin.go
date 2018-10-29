package models

import "github.com/greatbsky/gomvc/models/base"

type Admin struct {
	base.UIDModel
	Adminid string
	Pwd string
	Dopwd string
	Name string
	Email string
	Mobile string
}
