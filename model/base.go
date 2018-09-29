package model

import (
	"time"
)

type Base struct {
	Id      int64
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
	Deleted time.Time `xorm:"deleted"`
}

// type Privacy string

// func (this Privacy) String() string {
// 	return "privacy"
// }

type Sensitivity string

func (s Sensitivity) String() string {
	return "[SENSITIVE DATA]"
}

func (s Sensitivity) MarshalJSON() ([]byte, error) {
	return []byte(`"[SENSITIVE DATA]"`), nil
}
