package service

import (
	"fmt"
	"hotbed/dao"
	"hotbed/module/engine"
	"hotbed/tool/record"
)

var configDao dao.ConfigDao

type ConfigService struct{}

func (this *ConfigService) GetConfigByName(name string) interface{} {

	se := engine.GetSession()
	defer se.Close()

	config, err := configDao.GetByName(se, name)
	if config.Id == 0 || err != nil {
		record.GetRecorder().Errorf("GetConfigByName id:%v error:%v", config.Id, err)
		return nil
	}

	return config.Val
}

func (this *ConfigService) SetConfig(name string, value interface{}) bool {

	se := engine.GetSession()
	defer se.Close()

	config, err := configDao.GetByName(se, name)
	if err != nil {
		record.GetRecorder().Errorf("SetConfig id:%v error:%v", config.Id, err)
		return false
	}

	config.Val = fmt.Sprintf("%v", value)

	if config.Id == 0 {
		config.Name = name
		affected, _ := configDao.Set(se, config)
		return affected > 0
	}

	affected, _ := configDao.UpdateById(se, config.Id, config)
	return affected > 0
}
