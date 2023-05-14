package express

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
)

type Express struct {
	currentRout string
	middleware  []func(http.HandlerFunc) http.HandlerFunc
	routes      map[string]map[string]http.HandlerFunc
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

func (e *Express) PUT(args ...interface{}) *Express {
	if len(args) == 1 && reflect.TypeOf(args[0]).Kind() == reflect.Func {
		e.addRoute(e.currentRout, http.MethodPut, args[0].(func(http.ResponseWriter, *http.Request)))
	} else if len(args) == 2 {
		path := args[0].(string)
		handler := args[1].(func(http.ResponseWriter, *http.Request))
		e.addRoute(path, http.MethodPut, handler)
	} else {
		log.Fatal("PUT method requires either a handler function or a path and a handler function")
	}
	return e
}

func (e *Express) addRoute(path string, method string, handler http.HandlerFunc) {
	if e.routes[path] == nil {
		e.routes[path] = make(map[string]http.HandlerFunc)
	}
	e.routes[path][method] = handler
}

func (e *Express) Route(path string) *Express {
	e.currentRout = path

	return e
}

// sve se vrti oko ove funkcije
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
