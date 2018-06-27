package daos

import (
	"hotbed/models"
	"hotbed/tools/record"

	"github.com/go-xorm/xorm"
)

type ConfigDao struct{}

func (this *ConfigDao) Set(se *xorm.Session, model *models.Config) bool {

	_, err := se.Insert(model)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}

func (this *ConfigDao) GetById(se *xorm.Session, id int64) *models.Config {

	model := new(models.Config)
	_, err := se.Id(id).Get(model)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return model
}

func (this *ConfigDao) GetByName(se *xorm.Session, name string) *models.Config {

	model := new(models.Config)
	_, err := se.Where("name = ?", name).Get(model)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return model
}

func (this *ConfigDao) DeleteById(se *xorm.Session, id int64) bool {

	model := new(models.Config)

	_, err := se.Id(id).Delete(model)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}

func (this *ConfigDao) UpdateById(se *xorm.Session, id int64, model *models.Config) bool {

	_, err := se.Id(id).Update(model)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}
