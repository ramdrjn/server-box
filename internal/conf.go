package serverbox

import (
	"github.com/BurntSushi/toml"
)

type ServerBoxConf struct {
	Servers    map[string]server
	Statistics statistics
}

type server struct {
	Bind_Ip   string
	Bind_Port uint16
	Debug     bool
}

type statistics struct {
	Host    string
	Port    uint16
	Enabled bool
}

func ProcessConfFile(sbc *SbContext, confFilePath string) (*ServerBoxConf, error) {
	var sbconf ServerBoxConf
	_, err := toml.DecodeFile(confFilePath, &sbconf)
	if err != nil {
		pe, ok := err.(toml.ParseError)
		if ok {
			sbc.Log.Error(pe.ErrorWithUsage())
		}
	}
	return &sbconf, err
}
