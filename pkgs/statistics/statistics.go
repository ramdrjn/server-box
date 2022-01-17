package statistics

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type Statistics interface {
}

func InitializeStatistics(sbc *common.SbContext) (Statistics, error) {
	return nil, nil
}
