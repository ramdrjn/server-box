package serverbox

import (
	"fmt"
	"net/http"
)

type ServerHttp struct {
	server *Server
	mux    *http.ServeMux
}

func (s *ServerHttp) InitializeServerInstance() error {
	err := s.server.stats.RegisterForStats()
	if err != nil {
		return err
	}
	err = s.server.state.RegisterForState()
	if err != nil {
		return err
	}

	s.mux = http.NewServeMux()

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
	host := fmt.Sprintf("%s:%d", s.server.bindIp, s.server.bindPort)
	err = http.ListenAndServe(host, s.mux)
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
	return err
}

func (s *ServerHttp) AttachRouter(router *Router) error {
	next := router.GetRoutes()
	pat, obj := next()
	for pat != "" {
		s.mux.Handle(pat, obj)
		pat, obj = next()
	}
	return nil
}
