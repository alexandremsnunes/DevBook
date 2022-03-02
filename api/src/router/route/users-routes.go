package route

import (
	"api/src/controllers"
	"net/http"
)

var usersRoutes = []Route{
	{
		URI:                    "/users",
		Method:                 http.MethodPost,
		Function:               controllers.CreateUser,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users",
		Method:                 http.MethodGet,
		Function:               controllers.GetUsers,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodGet,
		Function:               controllers.GetUserById,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodPut,
		Function:               controllers.UpdateUser,
		AuthenticationRequires: false,
	},
	{
		URI:                    "/users/{userId}",
		Method:                 http.MethodDelete,
		Function:               controllers.DeleteUser,
		AuthenticationRequires: false,
	},
}
