package serverbox

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type SbContext struct {
	Log    common.Logger
	Conf   ServerBoxConf
	Stats  Statistics
	State  State
	Server Server
}

var Log common.Logger

type ServerBoxConf struct {
	Servers    map[string]server
	Statistics statistics
}

type server struct {
	Bind_Ip   string
	Bind_Port uint16
	Debug     bool
}

type statistics struct {
	Host    string
	Port    uint16
	Enabled bool
}
