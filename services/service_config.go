package services

import (
	"fmt"
	"hotbed/daos"
)

var configDao daos.ConfigDao

type ConfigService struct{}

func (this *ConfigService) GetConfigByName(name string) string {
	config := configDao.GetByName(name)
	return config.Value
}

func (this *ConfigService) SetConfig(name string, value interface{}) bool {

	config := configDao.GetByName(name)
	config.Value = fmt.Sprintf("%v", value)

	if config.Id == 0 {
		config.Name = name
		return configDao.Set(config)
	}

	return configDao.UpdateById(config.Id, config)
}
