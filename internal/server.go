package serverbox

import (
	"errors"
	"fmt"
)

type ServerType uint8

const (
	invalid_server = iota
	http_server
)

type serverInstance interface {
	InitializeServerInstance() error
}

type Server struct {
	name           string
	sType          ServerType
	uuid           string
	bindIp         string
	bindPort       uint16
	stats          Statistics
	state          State
	serverInstance serverInstance
	enabled        bool
}

func convertServerType(sType string) (ServerType, error) {
	switch sType {
	case "http":
		return http_server, nil
	}
	return invalid_server, errors.New("invalid server type")
}

func getServerInstance(s *Server) (serverInstance, error) {
	switch s.sType {
	case http_server:
		return &ServerHttp{server: s}, nil
	}
	return nil, errors.New("invalid server type")
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

		server.sType, err = convertServerType(serverConf.Type)
		if err != nil {
			break
		}

		server.uuid, _ = generateUuid(serverName, server.bindIp,
			server.bindPort)

		statsConf := serverConf.Statistics
		if statsConf.Enabled {
			host := fmt.Sprintf("%s:%d", statsConf.Host,
				statsConf.Port)
			err = InitializeStatistics(server.uuid,
				host, &server.stats)
			if err != nil {
				break
				server.stats.RegisterForStats()
			}
		}
		stateConf := serverConf.State
		if stateConf.Enabled {
			host := fmt.Sprintf("%s:%d", stateConf.Host,
				stateConf.Port)
			err = InitializeState(server.uuid,
				host, &server.state)
			if err != nil {
				break
				server.state.RegisterForState()
			}
		}

		server.serverInstance, err = getServerInstance(server)
		if err != nil {
			break
		}

		err = server.serverInstance.InitializeServerInstance()
		if err != nil {
			break
		}

		server.enabled = true
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
