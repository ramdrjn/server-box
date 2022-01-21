package serverbox

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type SbContext struct {
	Log    common.Logger
	Conf   *ServerBoxConf
	Stats  *Statistics
	State  interface{}
	Server interface{}
}

var Log common.Logger
