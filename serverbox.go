package serverbox

import (
	"fmt"
	. "github.com/ramdrjn/serverbox/internal"
	"github.com/ramdrjn/serverbox/pkgs/common"
)

func Initialize(debug bool, confFilePath string) error {
	var sbcontext SbContext
	var err error

	fmt.Println("initializing Server box")

	logLevel := common.InfoLevel
	if debug {
		logLevel = common.DebugLevel
	}
	sbcontext.Log = common.InitializeLogger("serverbox", logLevel)

	//Update internal log so that its accessible everwhere
	Log = sbcontext.Log

	sbcontext.Log.Debug("server box logging initialized")

	err = common.ProcessConfFile(confFilePath, &sbcontext.Conf)
	if err != nil {
		sbcontext.Log.Error("configuration file %s failed: ",
			confFilePath, err)
		return err
	}

	err = InitializeServers(&sbcontext)
	if err != nil {
		return err
	}

	ShutDownServers(&sbcontext)

	return nil
}
