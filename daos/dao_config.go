package daos

import (
	"hotbed/models"
	"hotbed/modules/engine"
	"hotbed/tools/record"
)

type ConfigDao struct{}

func (this *ConfigDao) Set(model *models.Config) bool {

	orm := engine.GetEngine()

	_, err := orm.Insert(model)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)

	return false
}

func (this *ConfigDao) GetById(id int64) (model *models.Config) {

	orm := engine.GetEngine()

	_, err := orm.Id(id).Get(model)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return model
}

func (this *ConfigDao) GetByName(name string) (model *models.Config) {

	orm := engine.GetEngine()

	_, err := orm.Where("name = ?", name).Get(model)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return model
}

func (this *ConfigDao) DeleteById(id int64) bool {

	orm := engine.GetEngine()

	model := new(models.Config)

	_, err := orm.Id(id).Delete(model)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)

	return false
}

func (this *ConfigDao) UpdateById(id int64, model *models.Config) bool {

	orm := engine.GetEngine()

	_, err := orm.Id(id).Update(model)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)

	return false
}
