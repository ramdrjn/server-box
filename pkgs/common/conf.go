package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
)

func ProcessConfFile(confFile string, confObj interface{}) error {
	_, err := toml.DecodeFile(confFile, confObj)
	if err != nil {
		pe, ok := err.(toml.ParseError)
		if ok {
			fmt.Println(pe.ErrorWithUsage())
		}
	}
	fmt.Println("configuration read from file ", confObj)
	return err
}
