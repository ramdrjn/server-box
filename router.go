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
	if f {
		f(res, req)
	}
}

type Router struct {
	routes []route
}

func (R *Router) RegisterRoute(pattern string, methods string, handler routeHandler, userdata interface{}) error {
	r := route{userdata: userdata, pattpattern}
	for _, method = range strings.Spit(methods, ",") {
		r.handlers[method] = handler
	}
	//mux.Handle(pattern, r)
	R.routes = append(R.routes, r)
	return nil
}
