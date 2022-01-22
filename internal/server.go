package serverbox

type Server struct {
	enabled bool
}

func InitializeServer(sbc *SbContext) (err error) {
	err = nil
	sbc.Stats.RegisterForStats("test", "server")
	return err
}
