package statistics

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type StatsContext struct {
	Log  common.Logger
	Conf *StatisticsConf
}

var Log common.Logger

type StatisticsConf struct {
	Host string
	Port uint32
}
