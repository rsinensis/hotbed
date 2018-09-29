package task

import (
	"os"
	"path/filepath"
	"time"

	macaron "gopkg.in/macaron.v1"
)

func uploadTask(t time.Time) {

	tmpPath := filepath.Join(macaron.Root, macaron.Config().Section("static").Key("static_path").String(), macaron.Config().Section("upload").Key("temp_path").String())

	filepath.Walk(tmpPath, func(file string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && info.ModTime().Unix() < (time.Now().Unix()-60*60*24*7) {
			os.Remove(file)
		}

		return nil
	})
}
