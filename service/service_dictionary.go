package service

import (
	"hotbed/dao"
	"hotbed/model"
	"hotbed/module/engine"
	"hotbed/tool/record"
)

type DictionaryService struct{}

var dictionaryTypeDao dao.DictionaryTypeDao
var dictionaryItemDao dao.DictionaryItemDao

func (this *DictionaryService) SetType(dt *model.DictionaryType) bool {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.Set(se, dt)
}

func (this *DictionaryService) GetAllType() []model.DictionaryType {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.GetAll(se)
}

func (this *DictionaryService) GetTypeByCode(code string) *model.DictionaryType {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.GetByCode(se, code)
}

func (this *DictionaryService) GetTypeById(id int64) *model.DictionaryType {
	se := engine.GetSession()
	defer se.Close()

	return dictionaryTypeDao.GetById(se, id)
}

func (this *DictionaryService) UpdateTypeById(id int64, dt *model.DictionaryType) bool {
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
		record.GetRecorder().Errorf("DeleteType commit error:%v", err)
		return false
	}

	return true
}
