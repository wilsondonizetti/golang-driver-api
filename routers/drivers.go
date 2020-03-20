package routers

import (
	"drivers.api/controllers"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

func SetDriversRoutes(router *mux.Router) *mux.Router {
	router.Handle("/drivers",
		negroni.New(
			//negroni.HandlerFunc(authentication.RequireTokenAuthentication),
			negroni.HandlerFunc(controllers.DriversController),
		)).Methods("GET")

	return router
}
