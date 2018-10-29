package base

/*
Description:
base.Model definition for all models parent type, UID as primary key
 * Author: architect.bian
 * Date: 2018/10/29 12:08
*/
type UIDModel struct {
	UID        string `gorm:"primary_key"`
	ModelBase
}
