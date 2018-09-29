package controller

import (
	"hotbed/module/env"
	"hotbed/service"
	"hotbed/tool/result"

	"github.com/astaxie/beego/validation"
	macaron "gopkg.in/macaron.v1"
)

var configService service.ConfigService

func controllerConfigInit(m *macaron.Macaron) {
	m.Group("/config", func() {
		m.Any("/get", getConfig)
		m.Any("/set", setConfig)
	})
}

func getConfig(ctx *env.Env) {
	name := ctx.QueryTrim("name")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	if valid.HasErrors() {
		ctx.JSON(200, result.FailByCode(result.INVALID_PARAMS, valid.Errors[0].Message))
		return
	}

	val := configService.GetConfigByName(name)

	ctx.JSON(200, result.OkByCode(result.SUCCESS, val))
}

func setConfig(ctx *env.Env) {
	name := ctx.QueryTrim("name")
	val := ctx.QueryTrim("val")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(val, "val").Message("值不能为空")
	if valid.HasErrors() {
		ctx.JSON(200, result.FailByCode(result.INVALID_PARAMS, valid.Errors[0].Message))
		return
	}

	if configService.SetConfig(name, val) {
		ctx.JSON(200, result.Ok())
		return
	}

	ctx.JSON(200, result.Fail())
}
