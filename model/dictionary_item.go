package model

import "fmt"

type DictionaryItem struct {
	Base  `xorm:"extends"`
	Pid   int64
	Code  string
	Name  string
	Val   string
	Class int
	Color string
}

func (this *DictionaryItem) Info() string {
	return fmt.Sprintf("%+v", this)
}
