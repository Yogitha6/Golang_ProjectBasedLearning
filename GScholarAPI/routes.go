package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type GScholarAPIRoute struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type GScholarAPIRoutes []GScholarAPIRoute

var gscholarAPIroutes = GScholarAPIRoutes{GScholarAPIRoute{"HomePageGet", "GET", "/", HomePageGETHandler},
GScholarAPIRoute{"HomagePagePost", "POST", "/", HomePagePOSTHandler},
GScholarAPIRoute{"GScholarAPI", "GET", "/gscholarapi/{name}", APIHandler}}

func GScholarAPIRouter() *mux.Router {
	gscholarAPIrouter := mux.NewRouter()
	for _, route := range gscholarAPIroutes {
		gscholarAPIrouter.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(route.HandlerFunc)
	}
	return gscholarAPIrouter
}