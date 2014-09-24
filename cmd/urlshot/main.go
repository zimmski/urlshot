package main

import (
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/zimmski/nethead/controller"

	"github.com/zimmski/urlshot/controller/url"
)

func addControllerToRoute(c controller.Controller, router *mux.Router) {
	s := router.PathPrefix("/" + c.UID()).Subrouter()

	controller.AddControllerToRoute(c, s)
}

func main() {
	router := mux.NewRouter()

	addControllerToRoute(url.New(), router)

	n := negroni.New()

	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewStatic(http.Dir("public")))

	n.UseHandler(router)

	n.Run(":3000")
}
