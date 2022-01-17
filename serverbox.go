package serverbox

import (
	"fmt"
	"github.com/ramdrjn/serverbox/pkgs/common"
)

func Initialize(debug bool, confFilePath string) error {
	fmt.Println("initializing Server box")

	logLevel := common.InfoLevel
	if debug {
		logLevel = common.DebugLevel
	}
	log := common.InitializeLogger(logLevel)
	log.Debug("server box logging initialized")

	conf, err := common.InitializeConfFile(log, confFilePath)
	if err != nil {
		return err
	}
	log.Debug(conf)

	return nil
}
