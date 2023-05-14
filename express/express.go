package express

import (
	"fmt"
	"net/http"
)

type Express struct {
	middleware []func(http.HandlerFunc) http.HandlerFunc
	routes     map[string]map[string]http.HandlerFunc
}

func New() *Express {
	return &Express{routes: make(map[string]map[string]http.HandlerFunc)}
}

func (e *Express) Use(middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	e.middleware = append(e.middleware, middleware...)
}

func (e *Express) POST(path string, handler http.HandlerFunc) {
	e.addRoute(path, http.MethodPost, handler) //http.MethodPost => "POST"
}

func (e *Express) GET(path string, handler http.HandlerFunc) {
	e.addRoute(path, http.MethodGet, handler) // http.MethodGet => "GET"
}

func (e *Express) addRoute(path string, method string, handler http.HandlerFunc) {
	if e.routes[path] == nil {
		e.routes[path] = make(map[string]http.HandlerFunc)
	}
	e.routes[path][method] = handler
}

func (e *Express) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Stvorimo novi http.Handler
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ako handler postoji u rutama, izvršimo ga
		if h, ok := e.routes[r.URL.Path][r.Method]; ok {
			h(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	// Primijenimo sve middleware funkcije na handler
	for i := len(e.middleware) - 1; i >= 0; i-- {
		handler = e.middleware[i](handler)
	}

	// Izvršimo konačni http.Handler
	handler.ServeHTTP(w, r)
}

func (e *Express) Listen(port string) {
	http.ListenAndServe(fmt.Sprintf(":%s", port), e)
}
