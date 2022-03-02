package router

import "github.com/gorilla/mux"

//This function generate the router's configurations
func Generate() *mux.Router {
	return mux.NewRouter()
}
