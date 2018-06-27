package routers

import (
	"hotbed/models"
	"hotbed/modules/env"
	"hotbed/services"
	"hotbed/tools/result"

	"github.com/astaxie/beego/validation"
	macaron "gopkg.in/macaron.v1"
)

var dictionaryService services.DictionaryService

func routerDictionnaryInit(m *macaron.Macaron) {
	m.Group("/dictionary", func() {
		m.Group("/type", func() {
			m.Any("/list", listDictionnaryType)
			m.Any("/get", getDictionnaryType)
			m.Any("/set", setDictionnaryType)
			m.Any("/update", updateDictionnaryType)
			m.Any("/delete", deleteDictionnaryType)
		})
	})
}

func listDictionnaryType(ctx *env.Env) {
	typs := dictionaryService.GetAllType()

	ctx.JSON(200, result.OkByCode(result.SUCCESS, typs))
}

func getDictionnaryType(ctx *env.Env) {
	id := ctx.QueryInt64("id")

	valid := validation.Validation{}

	valid.Required(id, "id").Message("ID不能为空")

	if valid.HasErrors() {
		ctx.JSON(200, result.FailByCode(result.INVALID_PARAMS, valid.Errors[0].Message))
		return
	}

	ty := dictionaryService.GetTypeById(id)

	ctx.JSON(200, result.OkByCode(result.SUCCESS, ty))
}

func setDictionnaryType(ctx *env.Env) {

	name := ctx.QueryTrim("name")
	code := ctx.QueryTrim("code")

	valid := validation.Validation{}

	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(code, "code").Message("类型不能为空")

	if valid.HasErrors() {
		ctx.JSON(200, result.FailByCode(result.INVALID_PARAMS, valid.Errors[0].Message))
		return
	}

	dictionaryType := dictionaryService.GetTypeByCode(code)
	if dictionaryType.Id != 0 {
		ctx.JSON(200, result.FailByCode(result.EXIST, code))
		return
	}

	dt := new(models.DictionaryType)
	dt.Code = code
	dt.Name = name

	if dictionaryService.SetType(dt) {
		ctx.JSON(200, result.Ok())
		return
	}

	ctx.JSON(200, result.Fail())
}

func updateDictionnaryType(ctx *env.Env) {

	name := ctx.QueryTrim("name")
	code := ctx.QueryTrim("code")
	id := ctx.QueryInt64("id")

	valid := validation.Validation{}

	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(name, "name").Message("名称不能为空")
	valid.Required(code, "code").Message("类型不能为空")

	if valid.HasErrors() {
		ctx.JSON(200, result.FailByCode(result.INVALID_PARAMS, valid.Errors[0].Message))
		return
	}

	idType := dictionaryService.GetTypeById(id)
	if idType.Id == 0 {
		ctx.JSON(200, result.FailByCode(result.NOTEXIST, id))
		return
	}

	codeType := dictionaryService.GetTypeByCode(code)
	if codeType.Id != 0 {
		ctx.JSON(200, result.FailByCode(result.EXIST, code))
		return
	}

	dt := new(models.DictionaryType)
	dt.Code = code
	dt.Name = name

	if dictionaryService.UpdateTypeById(id, dt) {
		ctx.JSON(200, result.Ok())
		return
	}

	ctx.JSON(200, result.Fail())
}

func deleteDictionnaryType(ctx *env.Env) {
	id := ctx.QueryInt64("id")

	valid := validation.Validation{}

	valid.Required(id, "id").Message("ID不能为空")

	if valid.HasErrors() {
		ctx.JSON(200, result.FailByCode(result.INVALID_PARAMS, valid.Errors[0].Message))
		return
	}

	if dictionaryService.DeleteType(id) {
		ctx.JSON(200, result.Ok())
		return
	}

	ctx.JSON(200, result.Fail())
}
