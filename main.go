package main

import (
	"root/cors"
	"root/express"
	"root/routes"
)

func main() {
	app := express.New()

	app.Use(cors.New)

	app.Route("/user").
		GET(routes.GetUser).
		POST(routes.PostUser).
		PUT(routes.PutUser).
		HEAD(routes.PostUser).
		PATCH(routes.PostUser)

	app.GET("/", routes.GetUser)

	app.Listen("5000")
}
