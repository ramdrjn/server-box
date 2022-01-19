package statistics

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type StatsContext struct {
	Log  common.Logger
	Conf *StatisticsConf
}

var Log common.Logger
