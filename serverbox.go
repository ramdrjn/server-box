package serverbox

import (
	"fmt"
	. "github.com/ramdrjn/serverbox/internal"
	"github.com/ramdrjn/serverbox/pkgs/common"
)

func Initialize(debug bool, confFilePath string) (sbcontext *SbContext, err error) {
	sbcontext = new(SbContext)

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
		return nil, err
	}

	err = InitializeServers(sbcontext)

	return sbcontext, err
}

func Run(sbcontext *SbContext) (err error) {
	err = RunServers(sbcontext)
	return err
}

func ShutDown(sbcontext *SbContext) (err error) {
	err = ShutDownServers(sbcontext)
	return err
}

func Abort(sbcontext *SbContext) (err error) {
	err = AbortServers(sbcontext)
	return err
}
