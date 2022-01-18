package serverbox

import (
	"github.com/BurntSushi/toml"
)

type ServerBoxConf struct {
	Servers map[string]Server `toml:"servers"`
}

type ServerConf struct {
	Bind_Ip   string
	Bind_Port uint32
	Debug     bool
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
