package models

import "fmt"

type Diploma struct {
	TokenID  int64  `json:"tokenID"`
	Student  string `json:"student"`
	Provider int64  `json:"-"`
}

func (d *Diploma) ToSQL() string {
	return fmt.Sprintf("(%d, '%s', %d)", d.TokenID, d.Student, d.Provider)
}
