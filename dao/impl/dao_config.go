package dao

import (
	"hotbed/model"

	"github.com/go-xorm/xorm"
)

type ConfigDaoImpl struct{}

func (this *ConfigDaoImpl) Set(se *xorm.Session, model *model.Config) (int64, error) {
	return se.Insert(model)
}

func (this *ConfigDaoImpl) GetById(se *xorm.Session, id int64) (*model.Config, error) {
	model := new(model.Config)
	_, err := se.Id(id).Get(model)
	return model, err
}

func (this *ConfigDaoImpl) GetByName(se *xorm.Session, name string) (*model.Config, error) {
	model := new(model.Config)
	_, err := se.Where("name = ?", name).Get(model)
	return model, err
}

func (this *ConfigDaoImpl) DeleteById(se *xorm.Session, id int64) (int64, error) {
	model := new(model.Config)
	return se.Id(id).Delete(model)
}

func (this *ConfigDaoImpl) UpdateById(se *xorm.Session, id int64, model *model.Config) (int64, error) {
	return se.Id(id).Update(model)
}
