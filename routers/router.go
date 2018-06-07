package routers

import (
	macaron "gopkg.in/macaron.v1"
)

func RouterInit(m *macaron.Macaron) {
	routerIndexInit(m)
	routerDictionnaryInit(m)
}
