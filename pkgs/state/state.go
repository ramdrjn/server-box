package state

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type State interface {
}

func InitializeState(sbc *common.SbContext) (State, error) {
	return nil, nil
}
