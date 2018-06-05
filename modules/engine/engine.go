package engine

import (
	"sync"

	"github.com/go-xorm/xorm"
)

//var engineMap map[string]*xorm.Engine

var engineMap sync.Map

func GetEngine(keys ...string) *xorm.Engine {

	// if len(keys) == 0 {
	// 	return engineMap["default"]
	// }

	// return engineMap[keys[0]]

	if len(keys) == 0 {
		vv, ok := engineMap.Load("default")
		if ok {
			return vv.(*xorm.Engine)
		}
	}

	vv, ok := engineMap.Load(keys[0])
	if ok {
		return vv.(*xorm.Engine)
	}

	return nil
}

func SetEngine(key string, e *xorm.Engine) {
	//engineMap[key] = e

	engineMap.Store(key, e)
}
