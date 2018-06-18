package daos

import (
	"hotbed/models"
	"hotbed/modules/engine"
	"hotbed/tools/record"
)

type DictionaryTypeDao struct{}

func (this *DictionaryTypeDao) Set(dt *models.DictionaryType) bool {

	orm := engine.GetEngine()

	_, err := orm.Insert(dt)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}

func (this *DictionaryTypeDao) GetById(id int64) *models.DictionaryType {

	orm := engine.GetEngine()

	dt := new(models.DictionaryType)
	_, err := orm.Id(id).Get(dt)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return dt
}

func (this *DictionaryTypeDao) GetByCode(code string) *models.DictionaryType {

	orm := engine.GetEngine()

	dt := new(models.DictionaryType)
	_, err := orm.Where("code = ?", code).Get(dt)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return dt
}

func (this *DictionaryTypeDao) DeleteById(id int64) bool {

	orm := engine.GetEngine()

	dt := new(models.DictionaryType)

	_, err := orm.Id(id).Delete(dt)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}

func (this *DictionaryTypeDao) UpdateById(id int64, dt *models.DictionaryType) bool {

	orm := engine.GetEngine()

	_, err := orm.Id(id).Update(dt)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}
