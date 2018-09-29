package controller

import (
	macaron "gopkg.in/macaron.v1"
)

func ControllerInit(m *macaron.Macaron) {
	controllerIndexInit(m)
	controllerDictionnaryInit(m)
	controllerConfigInit(m)
}
