package routes

import "net/http"

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("GetUsers"))
}

func PostUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PostUsers"))
}

func PutUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("PutUsers"))
}

func Branko(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Cap Branlp"))
}
