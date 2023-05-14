package main

import (
	"net/http"
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

	// app.Use()

	app.POST("/", myHandler)
	app.GET("/", myHandlerE)

	app.Listen(PORT)
}
