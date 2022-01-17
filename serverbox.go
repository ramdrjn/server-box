package serverbox

import (
	"fmt"
	"github.com/ramdrjn/serverbox/pkgs/common"
	"github.com/ramdrjn/serverbox/pkgs/server"
	"github.com/ramdrjn/serverbox/pkgs/state"
	"github.com/ramdrjn/serverbox/pkgs/statistics"
)

func Initialize(debug bool, confFilePath string) error {
	var sbcontext common.SbContext
	var err error

	fmt.Println("initializing Server box")

	logLevel := common.InfoLevel
	if debug {
		logLevel = common.DebugLevel
	}
	sbcontext.Log = common.InitializeLogger(&sbcontext, logLevel)
	log := sbcontext.Log.(common.Logger)
	log.Debug("server box logging initialized")

	sbcontext.Conf, err = common.ProcessConfFile(&sbcontext,
		confFilePath)
	if err != nil {
		return err
	}
	log.Debug("configuration read from file ",
		sbcontext.Conf.(*common.ServerBoxConf))

	sbcontext.Stats, err = statistics.InitializeStatistics(&sbcontext)
	if err != nil {
		return err
	}

	sbcontext.State, err = state.InitializeState(&sbcontext)
	if err != nil {
		return err
	}

	sbcontext.Server, err = server.InitializeServer(&sbcontext)
	if err != nil {
		return err
	}

	return nil
}
