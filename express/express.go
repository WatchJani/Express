package express

import (
	"fmt"
	"net/http"
)

type Express struct {
	middleware []func(http.Handler) http.Handler
	routes     map[string]map[string]http.HandlerFunc
}

func New() *Express {
	return &Express{routes: make(map[string]map[string]http.HandlerFunc)}
}

func (e *Express) Use(middleware ...func(http.Handler) http.Handler) {
	e.middleware = append(e.middleware, middleware...)
}

func (e *Express) POST(path string, handler http.HandlerFunc) {
	e.toRoute(path, http.MethodPost, handler) //http.MethodPost => "POST"
}

func (e *Express) GET(path string, handler http.HandlerFunc) {
	e.toRoute(path, http.MethodGet, handler) // http.MethodGet => "GET"
}

func (e *Express) toRoute(path string, method string, handler http.HandlerFunc) {
	if e.routes[path] == nil {
		e.routes[path] = make(map[string]http.HandlerFunc)
	}
	e.routes[path][method] = handler
}

func (e *Express) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler, ok := e.routes[r.URL.Path][r.Method]; ok {
		handler(w, r)
	} else {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (e *Express) Listen(port string) {
	http.ListenAndServe(fmt.Sprintf(":%s", port), e)
}
