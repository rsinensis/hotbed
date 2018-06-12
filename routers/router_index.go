package routers

import (
	"fmt"
	"hotbed/modules/env"
	"os"
	"path/filepath"
	"strings"
	"time"

	macaron "gopkg.in/macaron.v1"
)

const rule = "png|jpeg|bmp|svg|jpg|gif"
const size int64 = 10 << 20

func routerIndexInit(m *macaron.Macaron) {
	m.Get("/test", test)
	m.Get("/favicon.ico", favicon)
	m.Get("/robots.txt", robots)

	m.Any("/upload", upload)
}

func test(ctx *env.Env) {
	ctx.PlainText(200, []byte("ok"))

}

func favicon(ctx *env.Env) {
	ctx.ServeFile(filepath.Join(macaron.Config().Section("static").Key("static_path").String(), "favicon.ico"))
}

func robots(ctx *env.Env) {
	ctx.ServeFile(filepath.Join(macaron.Config().Section("static").Key("static_path").String(), "robots.txt"))
}

func upload(ctx *env.Env) {

	_, fh, err := ctx.GetFile("file")
	if err != nil {
		return
	}

	if len(fh.Filename) < 3 {
		return
	}

	pos := strings.LastIndex(fh.Filename, ".")
	ext := fh.Filename[pos+1:]
	if !strings.Contains(ext, rule) {
		return
	}

	if size < fh.Size {
		return
	}

	filePath := filepath.Join(macaron.Root, macaron.Config().Section("static").Key("static_path").String(), macaron.Config().Section("upload").Key("temp_path").String())

	os.MkdirAll(filePath, os.ModePerm)

	name := fmt.Sprintf("%v.%v", time.Now().Unix(), ext)
	fileName := filepath.Join(filePath, name)

	err = ctx.SaveToFile("file", fileName)
	if err != nil {
		return
	}
}

func uploadMove(name string) error {

	tmpPath := filepath.Join(macaron.Root, macaron.Config().Section("static").Key("static_path").String(), macaron.Config().Section("upload").Key("temp_path").String())
	uplodPath := filepath.Join(macaron.Root, macaron.Config().Section("static").Key("static_path").String(), macaron.Config().Section("upload").Key("upload_path").String())

	tempName := filepath.Join(tmpPath, name)
	uplodName := filepath.Join(uplodPath, name)

	return os.Rename(tempName, uplodName)
}
