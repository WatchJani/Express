package express

import (
	"fmt"
	"net/http"
)

type Express struct {
	currentRout string
	middleware  []func(http.HandlerFunc) http.HandlerFunc
	routes      map[string]map[string]http.HandlerFunc
}

// Initialization our framework
func New() *Express {
	return &Express{routes: make(map[string]map[string]http.HandlerFunc)}
}

// Adds new middleware at our req
func (e *Express) Use(middleware ...func(http.HandlerFunc) http.HandlerFunc) {
	e.middleware = append(e.middleware, middleware...)
}

// if you use the Route method, you only put your function (http.HandlerFuncs) as an argument.
//
// if you don't use it, you must put path as the private argument and your function
// (http.HandlerFuncs) as the second argument!
func (e *Express) POST(args ...interface{}) *Express {
	return methodHead(e, http.MethodPost, args)
}

// if you use the Route method, you only put your function (http.HandlerFuncs) as an argument.
//
// if you don't use it, you must put path as the private argument and your function
// (http.HandlerFuncs) as the second argument!
func (e *Express) GET(args ...interface{}) *Express {
	return methodHead(e, http.MethodGet, args)
}

// if you use the Route method, you only put your function (http.HandlerFuncs) as an argument.
//
// if you don't use it, you must put path as the private argument and your function
// (http.HandlerFuncs) as the second argument!
func (e *Express) PUT(args ...interface{}) *Express {
	return methodHead(e, http.MethodPut, args)
}

// if you use the Route method, you only put your function (http.HandlerFuncs) as an argument.
//
// if you don't use it, you must put path as the private argument and your function
// (http.HandlerFuncs) as the second argument!
func (e *Express) DELETE(args ...interface{}) *Express {
	return methodHead(e, http.MethodDelete, args)
}

// if you use the Route method, you only put your function (http.HandlerFuncs) as an argument.
//
// if you don't use it, you must put path as the private argument and your function
// (http.HandlerFuncs) as the second argument!
func (e *Express) PATCH(args ...interface{}) *Express {
	return methodHead(e, http.MethodPatch, args)
}

// if you use the Route method, you only put your function (http.HandlerFuncs) as an argument.
//
// if you don't use it, you must put path as the private argument and your function
// (http.HandlerFuncs) as the second argument!
func (e *Express) HEAD(args ...interface{}) *Express {
	return methodHead(e, http.MethodHead, args)
}

// funkcija dodana samo zbog previse kopiranja istog koda
func methodHead(e *Express, method string, args []interface{}) *Express {
	path, handler := catch(e.currentRout, args...)
	e.addRoute(path, method, handler)

	return e
}

// since the method can receive anything, no matter how many things we send to it,
// it first checks whether what was intended was sent, and constantly returns the path and function!
//
// it is assumed that only one or two arguments may be sent.
//
// if there is one argument, the function is sent
//
// if two arguments are sent, the first argument is the path and the second is the function.
func catch(curentPath string, args ...interface{}) (string, http.HandlerFunc) {
	if len(args) == 1 {
		// zasto nece da mi prihvati func(http.ResponseWriter, *http.Request) => http.handlerFunc???
		return curentPath, args[0].(func(http.ResponseWriter, *http.Request))
	}

	return args[0].(string), args[1].(func(http.ResponseWriter, *http.Request))
}

// tries to organize the code in such a way that no error occurs when naming the path at the same time
// checks if the path exists and puts it if it doesn't exist
func (e *Express) addRoute(path string, method string, handler http.HandlerFunc) {
	if e.routes[path] == nil {
		e.routes[path] = make(map[string]http.HandlerFunc)
	}
	e.routes[path][method] = handler
}

// add global route for all methods
func (e *Express) Route(path string) *Express {
	e.currentRout = path

	return e
}

func (e *Express) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Stvorimo novi http.HandlerFunc
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Ako handler postoji u rutama, izvrsimo ga
		if h, ok := e.routes[r.URL.Path][r.Method]; ok {
			h(w, r)
		} else {
			http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		}
	})

	// Primijenimo sve middleware funkcije na handler
	for i := len(e.middleware) - 1; i >= 0; i-- {
		//e.middleware[i] => func
		//handler => argument
		handler = e.middleware[i](handler)
	}

	// Izvršimo konacni http.Handler
	handler.ServeHTTP(w, r)
}

func (e *Express) Listen(port string) {
	http.ListenAndServe(fmt.Sprintf(":%s", port), e)
}
