package serverbox

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type SbContext struct {
	Log     common.Logger
	Conf    ServerBoxConf
	Servers map[string]Server
}

var Log common.Logger

type ServerBoxConf struct {
	Servers map[string]server
}

type server struct {
	Bind_Ip    string
	Bind_Port  uint16
	Debug      bool
	Statistics StatisticsConf
	State      StateConf
}

type StatisticsConf struct {
	Host    string
	Port    uint16
	Enabled bool
}

type StateConf struct {
	Host    string
	Port    uint16
	Enabled bool
}
