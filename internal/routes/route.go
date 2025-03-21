package routes

import (
	"fmt"
	"net/http"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
	Middlewares []func(http.HandlerFunc) http.HandlerFunc
}

func NewRoute(method string, path string) *Route {
	return &Route{
		Method: method,
		Path:   path,
	}
}

func (r *Route) SetHandler(h http.HandlerFunc) *Route {
	r.HandlerFunc = h
	return r
}

func (r *Route) AddMiddlewares(ms ...func(http.HandlerFunc) http.HandlerFunc) *Route {
	r.Middlewares = append(r.Middlewares, ms...)
	return r
}

func (r *Route) Register(mux *http.ServeMux) {
	mStack := r.HandlerFunc
    mStack = ApiMiddleware(mStack)

    for i := len(r.Middlewares)-1; i>=0; i-- {
        m := r.Middlewares[i]
		mStack = m(mStack)
    }

	mux.Handle(
		fmt.Sprintf("%s %s", r.Method, r.Path),
		mStack,
	)

}
