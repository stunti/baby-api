package main

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"log"

	"github.com/stunti/baby-api/api/model"

	r "github.com/dancannon/gorethink"
)

var session *r.Session

func init() {

	session, _ = r.Connect(r.ConnectOpts{
		Address: "localhost:28015",
	})
	r.DbCreate("api").Run(session)
}

func main() {
	r.Db("api").TableDrop("user").Run(session)
	r.Db("api").TableCreate("user").Run(session)
	/***************************************
	  USER
	 ***************************************/
	h := sha256.New()
	io.WriteString(h, "supertest")
	md := h.Sum(nil)
	password := hex.EncodeToString(md)

	/*
		user1 := model.User{
			Email:    "user1@example.com",
			Password: string(password),
		}
	*/
	_, err := r.Db("api").Table("user").Insert(model.User{Email: "user1@example.com", Password: string(password)}).Run(session)
	log.Println(err)
	log.Println(string(password))

	r.Db("api").Table("user").Insert(model.User{Email: "user2@example.com", Password: string(password)}).Run(session)
	r.Db("api").Table("user").Insert(model.User{Email: "user3@example.com", Password: string(password)}).Run(session)
	r.Table("user").Insert(model.User{Email: "user4@example.com", Password: string(password)}).Run(session)

}
