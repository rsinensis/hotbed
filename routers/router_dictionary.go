package routers

import (
	"hotbed/models"
	"hotbed/modules/env"
	"hotbed/services"
	"hotbed/tools/result"

	"github.com/astaxie/beego/validation"
	macaron "gopkg.in/macaron.v1"
)

var dictionaryTypeService services.DictionaryTypeService

func routerDictionnaryInit(m *macaron.Macaron) {
	m.Group("/dictionary", func() {
		m.Group("/type", func() {
			m.Any("/add", addDictionnaryType)
		})
	})
}

func addDictionnaryType(ctx *env.Env) {

	name := ctx.QueryTrim("name")
	code := ctx.QueryTrim("code")

	valid := validation.Validation{}

	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(code, "code").Message("类型不能为空")

	if valid.HasErrors() {
		ctx.JSON(200, result.FailByCode(result.INVALID_PARAMS, valid.Errors[0].Message))
		return
	}

	dictionaryType := dictionaryTypeService.GetByCode(code)
	if dictionaryType.Id != 0 {
		ctx.JSON(200, result.FailByCode(result.EXIST, "code"))
		return
	}

	dt := new(models.DictionaryType)
	dt.Code = code
	dt.Name = name

	if dictionaryTypeService.Set(dt) {
		ctx.JSON(200, result.Ok())
		return
	}

	ctx.JSON(200, result.Fail())
}
