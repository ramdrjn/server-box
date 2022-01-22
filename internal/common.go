package serverbox

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type SbContext struct {
	Log     common.Logger
	Conf    ServerBoxConf
	Servers map[string]*Server
}

var Log common.Logger

type ServerBoxConf struct {
	Servers map[string]server
}

type server struct {
	Bind_ip    string
	Bind_port  uint16
	Debug      bool
	Statistics statistics
	State      state
}

type statistics struct {
	Host    string
	Port    uint16
	Enabled bool
}

type state struct {
	Host    string
	Port    uint16
	Enabled bool
}
