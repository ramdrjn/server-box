package serverbox

import (
	"fmt"
	"github.com/ramdrjn/serverbox/pkgs/common"
)

func Initialize(debug bool, confFilePath string) {
	fmt.Println("initializing Server box")

	logLevel := common.InfoLevel
	if debug {
		logLevel = common.DebugLevel
	}
	common.InitializeLogger(logLevel)
	log := common.GetLogger()
	log.Debug("server box logging initialized")

	//common.InitializeConfFile(confFilePath)
}
