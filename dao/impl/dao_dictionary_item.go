package dao

import (
	"hotbed/model"
	"hotbed/tool/record"

	"github.com/go-xorm/xorm"
)

type DictionaryItemDao struct{}

func (this *DictionaryItemDao) Set(se *xorm.Session, di *model.DictionaryItem) bool {

	_, err := se.Insert(di)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}

func (this *DictionaryItemDao) GetAll(se *xorm.Session) []model.DictionaryItem {

	var dis []model.DictionaryItem
	_, err := se.Get(&dis)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return dis
}

func (this *DictionaryItemDao) GetByPid(se *xorm.Session, pid int64) []model.DictionaryItem {

	var dis []model.DictionaryItem
	_, err := se.Where("pid = ?", pid).Get(&dis)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return dis
}

func (this *DictionaryItemDao) GetById(se *xorm.Session, id int64) *model.DictionaryItem {

	di := new(model.DictionaryItem)
	_, err := se.Id(id).Get(di)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return di
}

func (this *DictionaryItemDao) GetByCode(se *xorm.Session, code string) *model.DictionaryItem {

	di := new(model.DictionaryItem)
	_, err := se.Where("code = ?", code).Get(di)

	if err != nil {
		record.GetRecorder().Error(err)
	}

	return di
}

func (this *DictionaryItemDao) DeleteById(se *xorm.Session, id int64) bool {

	di := new(model.DictionaryItem)

	_, err := se.Id(id).Delete(di)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}

func (this *DictionaryItemDao) UpdateById(se *xorm.Session, id int64, di *model.DictionaryItem) bool {

	_, err := se.Id(id).Update(di)

	if err == nil {
		return true
	}

	record.GetRecorder().Error(err)
	return false
}
