package serverbox

import (
	"errors"
	"fmt"
)

type ServerType uint8

const (
	INVALID = iota
	HTTP_SERVER
)

type Server struct {
	name     string
	sType    ServerType
	uuid     string
	bindIp   string
	bindPort uint16
	stats    Statistics
	state    State
	enabled  bool
}

func convertServerType(sType string) (ServerType, error) {
	switch sType {
	case "http":
		return HTTP_SERVER, nil
	}
	return INVALID, errors.New("invalid server type")
}

func generateUuid(name string, ip string, port uint16) (string, error) {
	return fmt.Sprintf("%s@%s:%d", name, ip, port), nil
}

func InitializeServers(sbc *SbContext) (err error) {
	sbc.Servers = make(map[string]*Server)

	for serverName, serverConf := range sbc.Conf.Servers {
		server := new(Server)
		server.name = serverName
		server.bindIp = serverConf.Bind_ip
		server.bindPort = serverConf.Bind_port
		server.sType, _ = convertServerType(serverConf.Type)

		server.uuid, _ = generateUuid(serverName, server.bindIp,
			server.bindPort)

		statsConf := serverConf.Statistics
		if statsConf.Enabled {
			host := fmt.Sprintf("%s:%d", statsConf.Host,
				statsConf.Port)
			err = InitializeStatistics(server.uuid,
				host, &server.stats)
			if err == nil {
				server.stats.RegisterForStats()
			}
		}
		stateConf := serverConf.State
		if stateConf.Enabled {
			host := fmt.Sprintf("%s:%d", stateConf.Host,
				stateConf.Port)
			err = InitializeState(server.uuid,
				host, &server.state)
			if err == nil {
				server.state.RegisterForState()
			}
		}
		sbc.Servers[serverName] = server
	}
	return err
}

func ShutDownServers(sbc *SbContext) (err error) {
	for _, server := range sbc.Servers {
		ShutDownStatistics(&server.stats)
		ShutDownState(&server.state)
	}
	return nil
}
