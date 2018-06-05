package routers

import (
	"hotbed/modules/env"
	"path/filepath"

	macaron "gopkg.in/macaron.v1"
)

func routerIndexInit(m *macaron.Macaron) {
	m.Get("/test", test)
	m.Get("/favicon.ico", favicon)
}

func test(ctx *env.Env) {
	ctx.PlainText(200, []byte("ok"))

}

func favicon(ctx *env.Env) {
	ctx.ServeFile(filepath.Join(macaron.Config().Section("static").Key("static_path").String(), "favicon.ico"))
}
