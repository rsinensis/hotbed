package dao

import (
	"hotbed/model"

	"github.com/go-xorm/xorm"
)

type DictionaryTypeDao interface {
	Set(se *xorm.Session, model *model.Config) (int64, error)
	GetById(se *xorm.Session, id int64) (*model.Config, error)
	GetByName(se *xorm.Session, name string) (*model.Config, error)
	DeleteById(se *xorm.Session, id int64) (int64, error)
	UpdateById(se *xorm.Session, id int64, model *model.Config) (int64, error)
}
