package handler

import (
	"github.com/stunti/baby-api/api/global"
	"github.com/stunti/baby-api/api/acl"
	"github.com/stunti/baby-api/api/model"

	"crypto/sha256"
	"encoding/hex"
	"fmt"
	r "github.com/dancannon/gorethink"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/context"
	"net/http"
	"time"
  "io"
)

func UserLoginHandler(w http.ResponseWriter, req *http.Request) {
	email := req.FormValue("email")
	password := req.FormValue("password")

	h := sha256.New()
	io.WriteString(h, password)
	md := h.Sum(nil)
	password = hex.EncodeToString(md)

	res, err := r.Table("user").Filter(
		r.Row.Field("Email").Eq(email),
	).Filter(
		r.Row.Field("Password").Eq(password),
	).Run(global.Session)
	// Fetch all the items from the database
	//res, err := r.Table("items").OrderBy(r.Asc("Created")).Run(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user model.User
	err = res.One(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//renderTemplate(w, "index", items)

	token := jwt.New(jwt.GetSigningMethod("RS256"))
	token.Claims["api"] = req.FormValue("api")
	token.Claims["user"] = user.Id
	token.Claims["role"] = acl.User
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	tokenString, _ := token.SignedString(global.PrivateKey)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	resultContent := global.JsonResponse{
		"result": tokenString,
	}

	//str, _ := json.Marshal(resultContent)
	fmt.Fprint(w, resultContent)

}

func UserProfileHandler(w http.ResponseWriter, req *http.Request) {

	//retrieve value from Token
	user_id := context.Get(req, "tokenUserId")

	res, err := r.Table("user").Get(user_id).Run(global.Session)
	// Fetch all the items from the database
	//res, err := r.Table("items").OrderBy(r.Asc("Created")).Run(session)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var user model.User
	err = res.One(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resultContent := global.JsonResponse{
		"result": user,
	}
	fmt.Fprint(w, resultContent)
}
