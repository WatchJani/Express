package main

import (
	"net/http"
	"root/cors"
	"root/express"
)

var PORT string = "5000"

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("kondic"))
}

func myHandlerE(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("janko"))
}

func main() {
	app := express.New()

	app.Use(cors.New)

	app.Route("/").PUT(myHandlerE)
	app.Route("/asdasd").PUT(myHandler)

	app.Route("/film").GET(myHandler).POST(myHandler).PUT(myHandlerE)

	app.PUT("/gre", myHandler)

	app.POST("/", myHandler)
	app.GET("/", myHandlerE)

	app.Listen(PORT)
}
