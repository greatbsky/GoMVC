package admindao

import (
	"github.com/greatbsky/gomvc/models"
	"github.com/greatbsky/gomvc/dao"
	"fmt"
)

func List(params map[string]interface{}) *[]models.Admin {
	var result []models.Admin
	var db = dao.DB.Table("admin")
	for k, v := range params {
		db = db.Where(k + " = ?", fmt.Sprint(v))
	}
	db.Find(&result)
	return &result
}