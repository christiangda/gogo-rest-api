package main

import (
	"github.com/christiangda/gogo-rest-api/routers"
	"github.com/urfave/negroni"
)

func main() {
	router := routers.InitRoutes()
	n := negroni.New()

	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())

	n.UseHandler(router)
	n.Run(":8080")
}
