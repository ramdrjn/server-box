package state

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type StateContext struct {
	Log  common.Logger
	Conf stateConf
}

var Log common.Logger

type stateConf struct {
	Host string
	Port uint32
}
