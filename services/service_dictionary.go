package services

import (
	"hotbed/daos"
	"hotbed/models"
)

type DictionaryTypeService struct{}

var dictionaryTypeDao daos.DictionaryTypeDao

func (this *DictionaryTypeService) Set(dt *models.DictionaryType) bool {

	return dictionaryTypeDao.Set(dt)
}

func (this *DictionaryTypeService) GetById(id int64) (dt *models.DictionaryType) {

	return dictionaryTypeDao.GetById(id)
}

func (this *DictionaryTypeService) UpdateById(id int64, dt *models.DictionaryType) bool {

	return dictionaryTypeDao.UpdateById(id, dt)
}

func (this *DictionaryTypeService) DeleteById(id int64) bool {

	return dictionaryTypeDao.DeleteById(id)
}
