package controller

import (
	"fmt"
	"hotbed/module/env"
	"hotbed/tool/record"
	"hotbed/tool/result"
	"os"
	"path/filepath"
	"strings"
	"time"

	macaron "gopkg.in/macaron.v1"
)

const rule = "png|jpeg|bmp|svg|jpg|gif"
const max int64 = 10
const size int64 = max << 20

func controllerIndexInit(m *macaron.Macaron) {
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
		record.GetRecorder().Error(err)
		ctx.JSON(200, result.New(500, "获取失败", nil))
		return
	}

	if len(fh.Filename) < 3 {
		ctx.JSON(200, result.New(400, "文件名不正确", nil))
		return
	}

	pos := strings.LastIndex(fh.Filename, ".")
	ext := fh.Filename[pos+1:]
	if !strings.Contains(ext, rule) {
		ctx.JSON(200, result.New(400, "文件格式不正确", nil))
		return
	}

	if size < fh.Size {
		ctx.JSON(200, result.New(400, fmt.Sprintf("文件不能大于%vMB", max), nil))
		return
	}

	filePath := filepath.Join(macaron.Root, macaron.Config().Section("static").Key("static_path").String(), macaron.Config().Section("upload").Key("temp_path").String())

	os.MkdirAll(filePath, os.ModePerm)

	name := fmt.Sprintf("%v.%v", time.Now().Unix(), ext)
	fileName := filepath.Join(filePath, name)

	err = ctx.SaveToFile("file", fileName)
	if err != nil {
		ctx.JSON(200, result.New(500, "获取失败", nil))
		return
	}

	ctx.JSON(200, result.New(200, "ok", fileName))
}

func uploadMove(name string) error {

	tmpPath := filepath.Join(macaron.Root, macaron.Config().Section("static").Key("static_path").String(), macaron.Config().Section("upload").Key("temp_path").String())
	uplodPath := filepath.Join(macaron.Root, macaron.Config().Section("static").Key("static_path").String(), macaron.Config().Section("upload").Key("upload_path").String())

	tempName := filepath.Join(tmpPath, name)
	uplodName := filepath.Join(uplodPath, name)

	return os.Rename(tempName, uplodName)
}
