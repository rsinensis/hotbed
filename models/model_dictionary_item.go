package models

import "fmt"

type DictionaryItem struct {
	Id    int64
	Pid   int64
	Name  string
	Value string
	Level int
	Color string
	Icon  string
}

func (this *DictionaryItem) Info() string {
	return fmt.Sprintf("%#v", this)
}
