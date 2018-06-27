package daos

import (
	"hotbed/models"
	"hotbed/tools/record"

	"github.com/go-xorm/xorm"
)

type DictionaryTypeDao struct{}

func (this *DictionaryTypeDao) Set(se *xorm.Session, dt *models.DictionaryType) bool {

	_, err := se.Insert(dt)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}

func (this *DictionaryTypeDao) GetAll(se *xorm.Session) []models.DictionaryType {

	var dts []models.DictionaryType
	_, err := se.Get(&dts)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return dts
}

func (this *DictionaryTypeDao) GetById(se *xorm.Session, id int64) *models.DictionaryType {

	dt := new(models.DictionaryType)
	_, err := se.Id(id).Get(dt)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return dt
}

func (this *DictionaryTypeDao) GetByCode(se *xorm.Session, code string) *models.DictionaryType {

	dt := new(models.DictionaryType)
	_, err := se.Where("code = ?", code).Get(dt)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return dt
}

func (this *DictionaryTypeDao) DeleteById(se *xorm.Session, id int64) bool {

	dt := new(models.DictionaryType)

	_, err := se.Id(id).Delete(dt)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}

func (this *DictionaryTypeDao) UpdateById(se *xorm.Session, id int64, dt *models.DictionaryType) bool {

	_, err := se.Id(id).Update(dt)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}
