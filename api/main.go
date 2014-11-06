package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/stunti/baby-api/api/global"
	"github.com/stunti/baby-api/api/handler"

	"github.com/codegangsta/negroni"
	r "github.com/dancannon/gorethink"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
)

func init() {
	var err error
	var tmp_sess *r.Session

	dir := os.Getenv("SOURCE_PATH")

	log.Println("local path: " + dir)

	global.PrivateKey, _ = ioutil.ReadFile(dir + "/keys/app.rsa")
	global.PublicKey, _ = ioutil.ReadFile(dir + "/keys/app.rsa.pub")
	tmp_sess, err = r.Connect(r.ConnectOpts{
		Address: os.Getenv("HOST_IP") + ":28015",
	})
	r.DbCreate("api").Run(tmp_sess)
	if err != nil {
		log.Println("database already exists")
	}
	r.Db("api").TableCreate("user").Run(tmp_sess)
	if err != nil {
		//log.Println("table user already exists")
		log.Println("error: %v", err)
	}
	global.Session, err = r.Connect(r.ConnectOpts{
		Address:  os.Getenv("HOST_IP") + ":28015",
		Database: "api",
	})
}

func AuthMiddleware() negroni.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
			return global.PublicKey, nil
		})
		if err == nil && token.Valid {
			context.Set(r, "tokenUserId", token.Claims["user"])
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "GO HOME SON")
		}
	}
}

func main() {
	restrictedRouter := mux.NewRouter()
	openRouter := mux.NewRouter()
	n := negroni.Classic()

	//restricted api access
	restrictedRouter.HandleFunc("/user/profile", handler.UserProfileHandler)
	secure := negroni.New()
	secure.Use(AuthMiddleware())
	secure.UseHandler(restrictedRouter)

	openRouter.HandleFunc("/open/login", handler.UserLoginHandler)

	openRouter.Handle("/user/profile", secure)

	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(openRouter)
	http.ListenAndServe(":8180", n)
}
