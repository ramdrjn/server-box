package state


import (
	"fmt"
	"github.com/ramdrjn/serverbox/pkgs/common"
	. "github.com/ramdrjn/serverbox/pkgs/state/internal"
)

func Initialize(debug bool, confFilePath string) error {
	var statecontext StateContext
	var err error

	fmt.Println("initializing state")

	logLevel := common.InfoLevel
	if debug {
		logLevel = common.DebugLevel
	}
	statecontext.Log = common.InitializeLogger("state", logLevel)

	//Update internal log so that its accessible everwhere
	Log = log

	statecontext.Log.Debug("state logging initialized")

	err = ProcessConfFile(confFilePath, &statecontext.Conf)
	if err != nil {
		statscontext.Log.Error("configuration file %s failed: ",
                        confFilePath, err)
		return err
	}

	err = InitializeGrpcServer(&statecontext)
	if err != nil {
		return err
	}

	ShutDownGrpcServer(&statecontext)

	return nil
}
