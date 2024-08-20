package config

import (
	"github.com/spf13/viper"
	"log"
)

const (
	CfgFileName = "antibot.yml"
	CfgFileType = "yml"
)

func init() {
	viper.SetConfigName(CfgFileName)
	viper.SetConfigType(CfgFileType)
	viper.AddConfigPath("/etc")
	//viper.AddConfigPath("./data")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("%s\n", err.Error())
	}
}
