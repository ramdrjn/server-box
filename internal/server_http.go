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

	return nil
}
