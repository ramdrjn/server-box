package serverbox

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
	"os"
)

type SbContext struct {
	Log           common.Logger
	Conf          serverBoxConf
	Servers       map[string]*Server
	SignalChannel chan os.Signal
}

var Log common.Logger

type serverBoxConf struct {
	Servers map[string]server
}

type server struct {
	Bind_ip        string
	Bind_port      uint16
	Debug          bool
	Type           string
	Statistics     statistics
	State          state
	Configurations ServerConfigurations
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

type ServerConfigurations struct {
	Http httpConfigurations
}

type httpConfigurations struct {
	Enabled      bool
	Static_dir   string
	Static_path  string
	Strip_path   string
	Template_dir string
}
