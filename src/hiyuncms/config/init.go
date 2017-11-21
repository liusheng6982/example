package config

import (
	"github.com/Unknwon/goconfig"
	"fmt"
	"log"
)

var cfg * goconfig.ConfigFile

func init()  {
	var err error
	confFile := "webroot/conf/config.ini"
	fmt.Printf( confFile )
	cfg, err = goconfig.LoadConfigFile( confFile )
	if err != nil {
		println("读取配置文件出错", err.Error() )
	}
}

func GetValue(key string) (string) {
	stringValue, err := cfg.GetValue(goconfig.DEFAULT_SECTION, key)
	if err != nil {
		log.Printf("读取配置%s出错%s\n", key, err.Error())
	}
	return stringValue
}

func GetInt(key string)(int)  {
	intValue, err := cfg.Int(goconfig.DEFAULT_SECTION, key)
	if err != nil {
		log.Printf("读取配置%s出错%s\n", key, err.Error())
	}
	return intValue
}

func GetBool(key string)(bool)  {
	boolValue, err := cfg.Bool(goconfig.DEFAULT_SECTION, key)
	if err != nil {
		log.Printf("读取配置%s出错%s\n", key, err.Error())
	}
	return boolValue
}



