package main

import (
	"net/http"

	"drivers.api/routers"
	"drivers.api/settings"
	"github.com/codegangsta/negroni"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5081", n)
}
