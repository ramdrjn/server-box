package serverbox

import (
	"context"
	"fmt"
	"github.com/ramdrjn/serverbox/pkgs/mux"
	"net/http"
)

type ServerHttp struct {
	server     *Server
	httpServer http.Server
}

func (s *ServerHttp) InitializeServerInstance() (err error) {
	s.httpServer.Addr = fmt.Sprintf("%s:%d", s.server.bindIp,
		s.server.bindPort)
	s.httpServer.Handler = http.NewServeMux()

	err = s.server.stats.RegisterForStats()
	if err != nil {
		return err
	}
	err = s.server.state.RegisterForState()
	if err != nil {
		return err
	}

	err = s.server.state.ReportState("maintanence")
	if err != nil {
		return err
	}

	return nil
}

func (s *ServerHttp) RunServerInstance() error {
	err := s.server.state.ReportState("up")
	if err != nil {
		return err
	}

	err = s.httpServer.ListenAndServe()
	if err != nil {
		Log.Error(err)
		s.server.state.ReportState("down")
	}
	return err
}

func (s *ServerHttp) ShutDownServerInstance() error {
	err := s.server.state.ReportState("down")
	if err != nil {
		return err
	}
	err = s.httpServer.Shutdown(context.TODO())
	return err
}

func (s *ServerHttp) AbortServerInstance() error {
	err := s.server.state.ReportState("down")
	if err != nil {
		return err
	}
	err = s.httpServer.Close()
	return err
}

func (s *ServerHttp) AttachRouterServerInstance(router mux.Router) error {
	mux := s.httpServer.Handler.(*http.ServeMux)
	next := router.GetRoutes()
	pat, obj := next()
	for pat != "" {
		mux.Handle(pat, obj)
		pat, obj = next()
	}
	return nil
}
