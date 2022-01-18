package serverbox

type Server struct {
	enabled bool
}

func InitializeServer(sbc *SbContext) (server *Server, err error) {
	err = nil
	server = new(Server)
	return server, err
}
