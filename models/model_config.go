package models

import "fmt"

type Config struct {
	Base `xorm:"extends"`
	Name string
	Val  string `xorm:"text"`
}

func (this *Config) Info() string {
	return fmt.Sprintf("%#v", this)
}
