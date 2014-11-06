package model

import (
	"time"
)

type User struct {
	Id          string `gorethink:"id,omitempty" json:"id,"`
	Dateofbirth time.Time `json:"dob,omitempty"`
	Email       string `json:"email,omitempty"`
	Phone       string `json:"phone,omitempty"`
  Password    string `json:"-"`
	Created     time.Time `json:"created,omitempty"`
	Updated     time.Time `json:"updated,omitempty"`
}

/*
func (t *User) Completed() bool {
    return t.Status == "complete"
}
*/

func NewUser(email string, phone string, dob time.Time) *User {
	return &User{
		Email:       email,
		Phone:       phone,
		Dateofbirth: dob,
	}
}
