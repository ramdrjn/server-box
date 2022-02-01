package mux

import (
	"net/http"
	"strings"
)

type HandlerArgs struct {
	HttpRes  http.ResponseWriter
	HttpReq  *http.Request
	UserData interface{}
}

type routeHandler func(*HandlerArgs)

type route struct {
	userdata interface{}
	pattern  string
	handlers map[string]routeHandler
}

func (r route) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	f := r.handlers[req.Method]
	if f != nil {
		f(&HandlerArgs{res, req, r.userdata})
	}
}

type Router struct {
	routes []route
}

func (r *Router) RegisterRoute(pattern string, methods string, handler routeHandler, userdata interface{}) error {
	route := route{userdata: userdata, pattern: pattern}
	route.handlers = make(map[string]routeHandler)
	for _, method := range strings.Split(methods, ",") {
		route.handlers[method] = handler
	}
	r.routes = append(r.routes, route)
	return nil
}

func NewRouter() *Router {
	return new(Router)
}
