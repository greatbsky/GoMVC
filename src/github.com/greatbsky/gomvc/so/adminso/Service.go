package adminso

import (
	"github.com/greatbsky/gomvc/dao/admindao"
	"github.com/greatbsky/gomvc/models"
)

func List() *[]models.Admin {
	params := make(map[string]interface{})
	params["adminid"] = "zhang"
	return admindao.List(params)
}

func ListAll() *[]models.Admin {
	return admindao.List(nil)
}