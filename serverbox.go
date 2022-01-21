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
	log := sbcontext.Log.(common.Logger)
	log.Debug("server box logging initialized")

	//Update internal log so that its accessible everwhere
	Log = log

	sbcontext.Conf, err = ProcessConfFile(&sbcontext, confFilePath)
	if err != nil {
		return err
	}
	log.Debug("configuration read from file ", sbcontext.Conf)

	sbcontext.Stats, err = InitializeStatistics(&sbcontext)
	if err != nil {
		return err
	}

	sbcontext.State, err = InitializeState(&sbcontext)
	if err != nil {
		return err
	}

	sbcontext.Server, err = InitializeServer(&sbcontext)
	if err != nil {
		return err
	}

	ShutDownStatistics(&sbcontext)

	return nil
}
