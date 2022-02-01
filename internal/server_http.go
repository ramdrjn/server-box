package serverbox

type ServerHttp struct {
	server *Server
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
	return err
}

func (s *ServerHttp) ShutDownServerInstance() error {
	err := s.server.state.ReportState("down")
	if err != nil {
		return err
	}
	return err
}

func (s *ServerHttp) AttachRouter(router interface{}) error {
	return nil
}
