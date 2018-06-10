package models

import "fmt"

type DictionaryType struct {
	Base
	Code string
	Name string
}

func (this *DictionaryType) Info() string {
	return fmt.Sprintf("%#v", this)
}
