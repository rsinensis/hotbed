package models

import "fmt"

type Config struct {
	Base
	Name  string
	Value string `xorm:"text"`
}

func (this *Config) Info() string {
	return fmt.Sprintf("%#v", this)
}
