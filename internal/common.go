package serverbox

import (
	"github.com/ramdrjn/serverbox/pkgs/common"
)

type SbContext struct {
	Log    common.Logger
	Conf   *ServerBoxConf
	Stats  interface{}
	State  interface{}
	Server interface{}
}
