package acl

import (
	"net/http"

	"github.com/gorilla/mux"
)

//routeKey

const (
	Open    = 1
	User    = 2
	Admin   = 4
	Unknown = 0
)

var acl = map[string]int{
	"V1UserLogin":   Open,
	"V1UserProfile": User,
}

func GetAclStatus(r *http.Request, router *mux.Router) int {
	//route is not dispatched yet so let's figured out the route which will be used
	var match mux.RouteMatch
	var currentRoute *mux.Route
	if router.Match(r, &match) {
		currentRoute = match.Route
	}

	name := currentRoute.GetName()
	if _, ok := acl[name]; ok {
		return acl[name]
	}
	return Unknown
}
