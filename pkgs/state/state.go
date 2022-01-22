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
	log := statecontext.Log.(common.Logger)
	log.Debug("state logging initialized")

	//Update internal log so that its accessible everwhere
	Log = log

	statecontext.Conf, err = ProcessConfFile(&statecontext, confFilePath)
	if err != nil {
		return err
	}
	log.Debug("configuration read from file ", statecontext.Conf)

	err = InitializeGrpcServer(&statecontext)
	if err != nil {
		return err
	}

	return nil
}
