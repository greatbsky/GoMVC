package base

/*
Description:
IDModel definition for all models parent type, ID as primary key
 * Author: architect.bian
 * Date: 2018/10/29 12:08
*/
type IDModel struct {
	ID         uint   `gorm:"primary_key"`
	ModelBase
}
