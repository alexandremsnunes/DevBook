package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

//Route's model
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequires bool
}

//Put all routes in router
func Setup(r *mux.Router) *mux.Router {
	routes := usersRoutes

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
