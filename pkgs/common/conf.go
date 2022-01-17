package common

import (
	"github.com/BurntSushi/toml"
)

type ServerBoxConf struct {
	Servers map[string]Server
}

type Server struct {
	Bind_Ip   string
	Bind_Port uint32
	Debug     bool
}

func InitializeConfFile(sbc *SbContext, confFilePath string) (*ServerBoxConf, error) {
	var sbconf ServerBoxConf
	_, err := toml.DecodeFile(confFilePath, &sbconf)
	if err != nil {
		pe, ok := err.(toml.ParseError)
		if ok {
			sbc.Log.(Logger).Error(pe.ErrorWithUsage())
		}
	}
	return &sbconf, err
}
