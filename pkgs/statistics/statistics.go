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

	//Update internal log so that its accessible everwhere
	Log = statscontext.Log

	statscontext.Log.Debug("statistics logging initialized")
	
	err = common.ProcessConfFile(confFilePath, &statscontext.Conf)
	if err != nil {
		statscontext.Log.Error("configuration file %s failed: ",
                        confFilePath, err)
		return err
	}

	err = InitializeGrpcServer(&statscontext)
	if err != nil {
		return err
	}

	ShutDownGrpcServer(&statscontext)

	return nil
}
