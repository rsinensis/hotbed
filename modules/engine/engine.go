package engine

import (
	"fmt"
	"sync"

	"github.com/go-xorm/xorm"
)

var engineMap sync.Map

func GetEngine(keys ...string) *xorm.Engine {
	if len(keys) == 0 {
		keys = append(keys, "default")
	}

	vv, ok := engineMap.Load(keys[0])
	if !ok {
		panic(fmt.Sprintf("get engine %v err", keys))
	}

	return vv.(*xorm.Engine)
}

func GetSession(keys ...string) *xorm.Session {
	e := GetEngine(keys...)
	return e.NewSession()
}

func SetEngine(key string, e *xorm.Engine) {
	engineMap.Store(key, e)
}
