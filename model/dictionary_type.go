package model

import "fmt"

type DictionaryType struct {
	Base `xorm:"extends"`
	Code string
	Name string
}

func (this *DictionaryType) Info() string {
	return fmt.Sprintf("%+v", this)
}
