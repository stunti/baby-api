package model

import (
	"time"
)

type User struct {
	Id          string `gorethink:"id,omitempty"`
	Dateofbirth time.Time
	Email       string
	Phone       string
  Password    string
	Created     time.Time
	Updated     time.Time
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
