package services

import (
	"fmt"
	"hotbed/daos"
)

var configDao daos.ConfigDao

type ConfigService struct{}

func (this *ConfigService) GetConfigByName(name string) interface{} {
	config := configDao.GetByName(name)
	if config.Id == 0 {
		return nil
	}
	return config.Val
}

func (this *ConfigService) SetConfig(name string, value interface{}) bool {

	config := configDao.GetByName(name)

	config.Val = fmt.Sprintf("%v", value)

	if config.Id == 0 {
		config.Name = name
		return configDao.Set(config)
	}

	return configDao.UpdateById(config.Id, config)
}
