package statistics

import (
	"fmt"
	"github.com/ramdrjn/serverbox/pkgs/common"
	. "github.com/ramdrjn/serverbox/pkgs/statistics/internal"
)

func Initialize(debug bool, confFilePath string) error {
	var statscontext StatsContext
	var err error

	fmt.Println("initializing statistics")

	logLevel := common.InfoLevel
	if debug {
		logLevel = common.DebugLevel
	}
	statscontext.Log = common.InitializeLogger("statistics", logLevel)
	log := statscontext.Log.(common.Logger)
	log.Debug("statistics logging initialized")

	statscontext.Conf, err = ProcessConfFile(&statscontext, confFilePath)
	if err != nil {
		return err
	}
	log.Debug("configuration read from file ", statscontext.Conf)

	err = InitializeGrpcServer(&statscontext)
	if err != nil {
		return err
	}

	return nil
}
