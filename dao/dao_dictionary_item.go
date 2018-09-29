/*
 * @Author: Mocos
 * @Date: 2018-09-29 16:54:15
 * @LastEditors: Mocos
 * @LastEditTime: 2018-09-29 16:54:55
 * @Description:
 */

package dao

import (
	"hotbed/model"

	"github.com/go-xorm/xorm"
)

type DictionaryItemDao interface {
	Set(se *xorm.Session, model *model.Config) (int64, error)
	GetById(se *xorm.Session, id int64) (*model.Config, error)
	GetByName(se *xorm.Session, name string) (*model.Config, error)
	DeleteById(se *xorm.Session, id int64) (int64, error)
	UpdateById(se *xorm.Session, id int64, model *model.Config) (int64, error)
}
