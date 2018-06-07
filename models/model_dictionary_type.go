package models

import "fmt"

type DictionaryType struct {
	Id   int64
	Code string
	Name string
}

func (this *DictionaryType) Info() string {
	return fmt.Sprintf("%#v", this)
}
