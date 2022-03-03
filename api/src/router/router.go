package router

import (
	"api/src/router/route"

	"github.com/gorilla/mux"
)

//This function generate the router's configurations
func Generate() *mux.Router {
	r := mux.NewRouter()
	return route.Setup(r)
}
