package mux

import (
	"errors"
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

type Router interface {
	RegisterRoute(pattern string, methods string, handler routeHandler,
		userdata interface{}) error
	GetRoutes() func() (string, route)
}

type router struct {
	routes []route
}

func convertMethod(method string) (string, error) {
	switch method {
	case "get":
		return http.MethodGet, nil
	case "post":
		return http.MethodPost, nil
	case "put":
		return http.MethodPut, nil
	case "delete":
		return http.MethodDelete, nil
	}
	return "", errors.New("invalid server type")
}

func (r *router) RegisterRoute(pattern string, methods string, handler routeHandler, userdata interface{}) error {
	route := route{userdata: userdata, pattern: pattern}
	route.handlers = make(map[string]routeHandler)
	for _, method := range strings.Split(methods, ",") {
		meth, err := convertMethod(method)
		if err == nil {
			route.handlers[meth] = handler
		} else {
			return err
		}
	}
	r.routes = append(r.routes, route)
	return nil
}

func (r *router) GetRoutes() func() (string, route) {
	var i int = 0
	max := len(r.routes)
	return func() (string, route) {
		if i < max {
			rou := r.routes[i]
			i++
			return rou.pattern, rou
		}
		return "", route{}
	}
}

func NewRouter() Router {
	return new(router)
}
