package statistics

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type StatsContext struct {
	Log  common.Logger
	Conf statisticsConf
}

var Log common.Logger

type statisticsConf struct {
	Host string
	Port uint32
}
