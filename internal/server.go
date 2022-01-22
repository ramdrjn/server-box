package serverbox

type Server struct {
	name     string
	bindIp   string
	bindPort uint16
	stats    Statistics
	state    State
}

func InitializeServers(sbc *SbContext) (err error) {
	for serverName, serverConf := range sbc.Conf.Servers {
		server := Server.New()
		server.name = serverName
		server.bindIp = serverConf.Bind_ip
		server.bindPort = serverConf.Bind_port

		err = InitializeStatistics(&server.stats,
			&serverConf.Statistics)
		if err != nil {
			return err
		}

		err = InitializeState(&server.state, &serverConf.State)
		if err != nil {
			return err
		}

		server.stats.RegisterForStats("test", "server")
		server.state.RegisterForState("test", "server")

		sbc.Servers[name] = server
	}
	return err
}

func ShutDownServers(sbc *SbContext) (err error) {
	ShutDownStatistics(&sbcontext)
	ShutDownState(&sbcontext)
}
