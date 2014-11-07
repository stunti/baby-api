package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/stunti/baby-api/api/acl"
	"github.com/stunti/baby-api/api/global"
	"github.com/stunti/baby-api/api/handler"

	"github.com/codegangsta/negroni"
	r "github.com/dancannon/gorethink"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/phyber/negroni-gzip/gzip"
)

var restrictedRouter *mux.Router

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

		//get status of matched route
		status := acl.GetAclStatus(r, restrictedRouter)
		log.Println("route status: " + strconv.Itoa(status))
		if status == acl.Open {
			next(w, r)
		} else {

			token, err := jwt.ParseFromRequest(r, func(token *jwt.Token) (interface{}, error) {
				return global.PublicKey, nil
			})
			if err == nil && token.Valid {
				log.Println("user status: ", strconv.Itoa(int(token.Claims["role"].(float64))))
				//grab the role from token
				if (int(token.Claims["role"].(float64)) & status) != 0 {
					context.Set(r, "tokenUserId", token.Claims["user"])
					next(w, r)
				} else {
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprint(w, "GO HOME SON you need ", status)
				}
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprint(w, "GO HOME SON")
			}
		}
	}
}

func main() {
	restrictedRouter = mux.NewRouter()

	restrictedRouter.HandleFunc("/v1/user/profile", handler.UserProfileHandler).Name("V1UserProfile")
	restrictedRouter.HandleFunc("/v1/user/login", handler.UserLoginHandler).Name("V1UserLogin")

	n := negroni.Classic()
	n.Use(AuthMiddleware())
	n.Use(gzip.Gzip(gzip.DefaultCompression))
	n.UseHandler(restrictedRouter)

	//lets save the handler for use in AuthMiddleware
	n.Run(":8180")
}
