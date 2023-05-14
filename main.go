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

func Setting(hf http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Generirajte novi ID zahtjeva
		requestID := "5"

		// Postavite novi header u zahtjev
		r.Header.Set("X-Request-ID", requestID)

		// Pozovite sljedeći handler
		hf.ServeHTTP(w, r)
	})
}

func CORS(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	app := express.New()

	app.Use(CORS)

	app.POST("/", myHandler)
	app.GET("/", myHandlerE)

	app.Listen(PORT)
}
