package route

import "net/http"

//Route's model
type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	AuthenticationRequires bool
}
