package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/oluu-web/hngx-stage2/cmd/api/controllers"
	"github.com/oluu-web/hngx-stage2/cmd/api/middleware"
)

func InitRoutes() http.Handler {
	router := httprouter.New()

	//define routes and their associate controller functions
	router.HandlerFunc(http.MethodPost, "/new", controllers.CreatePerson)
	router.HandlerFunc(http.MethodGet, "/person/:name", controllers.GetPerson)
	router.HandlerFunc(http.MethodPut, "/edit/:name", controllers.UpdatePerson)
	router.HandlerFunc(http.MethodDelete, "/delete/:name", controllers.DeletePerson)

	//Enable CORS (Cross_Origin Resource Sharing) middleware to allow cross-origin requests
	return middleware.EnableCORS(router)
}
