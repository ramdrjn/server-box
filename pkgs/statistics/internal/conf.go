package statistics

import (
	"github.com/BurntSushi/toml"
)

type StatisticsConf struct {
	Host string
	Port uint32
}

func ProcessConfFile(stc *StatsContext, confFilePath string) (*StatisticsConf, error) {
	var statsconf StatisticsConf
	_, err := toml.DecodeFile(confFilePath, &statsconf)
	if err != nil {
		pe, ok := err.(toml.ParseError)
		if ok {
			stc.Log.Error(pe.ErrorWithUsage())
		}
	}
	return &statsconf, err
}
