package server

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type Server interface {
}

func InitializeServer(sbc *common.SbContext) (Server, error) {
	return nil, nil
}
