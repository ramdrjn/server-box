package serverbox

type Server struct {
	enabled bool
}

func InitializeServer(sbc *SbContext) (server *Server, err error) {
	err = nil
	server = new(Server)
	sbc.Stats.RegisterForStats("test", "server")
	return server, err
}
