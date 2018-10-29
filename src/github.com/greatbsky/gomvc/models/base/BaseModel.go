package base

import "time"

/*
Description:
ModelBase definition for all models parent type, without id specify
 * Author: architect.bian
 * Date: 2018/10/29 16:32
 */
type ModelBase struct {
	Status     int32
	Createtime time.Time
	Updatetime time.Time
}
