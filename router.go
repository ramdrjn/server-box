package serverbox

import (
	"net/http"
	"strings"
)

type routeHandler func(userdata interface{}, res http.ResponseWriter,
	req *http.Request)

type route struct {
	userdata interface{}
	pattern  string
	handlers map[string]routeHandler
}

func (r route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	f := r.handlers[req.Method]
	if f != nil {
		f(r.userdata, res, req)
	}
}

type router struct {
	routes []route
}

func (r *router) RegisterRoute(pattern string, methods string, handler routeHandler, userdata interface{}) error {
	route := route{userdata: userdata, pattern: pattern}
	route.handlers = make(map[string]routeHandler)
	for _, method := range strings.Split(methods, ",") {
		route.handlers[method] = handler
	}
	//mux.Handle(pattern, r)
	r.routes = append(r.routes, route)
	return nil
}

func NewRouter() *router {
	return new(router)
}
