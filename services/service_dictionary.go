package services

import (
	"hotbed/daos"
	"hotbed/models"
	"hotbed/modules/engine"
	"hotbed/tools/record"
)

type DictionaryService struct{}

var dictionaryTypeDao daos.DictionaryTypeDao
var dictionaryItemDao daos.DictionaryItemDao

func (this *DictionaryService) SetType(dt *models.DictionaryType) bool {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.Set(se, dt)
}

func (this *DictionaryService) GetAllType() []models.DictionaryType {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.GetAll(se)
}

func (this *DictionaryService) GetTypeByCode(code string) *models.DictionaryType {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.GetByCode(se, code)
}

func (this *DictionaryService) GetTypeById(id int64) *models.DictionaryType {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.GetById(se, id)
}

func (this *DictionaryService) UpdateTypeById(id int64, dt *models.DictionaryType) bool {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.UpdateById(se, id, dt)
}

func (this *DictionaryService) DeleteType(id int64) bool {
	se := engine.GetSession()
	defer se.Close()

	se.Begin()

	if !dictionaryTypeDao.DeleteById(se, id) {
		se.Rollback()
		return false
	}

	if !dictionaryItemDao.DeleteById(se, id) {
		se.Rollback()
		return false
	}

	if err := se.Commit(); err != nil {
		record.GetRecorder().Error(err)
		return false
	}

	return true
}
