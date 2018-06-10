package models

import "fmt"

type DictionaryItem struct {
	Base
	Pid   int64
	Code  string
	Name  string
	Val   string
	Class int
	Color string
	Icon  string
}

func (this *DictionaryItem) Info() string {
	return fmt.Sprintf("%#v", this)
}
