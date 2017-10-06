package config

import (
	"github.com/Unknwon/goconfig"
	"fmt"
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
	stringValue, _ := cfg.GetValue(goconfig.DEFAULT_SECTION, key)
	return stringValue
}

func GetInt(key string)(int)  {
	intValue, _ := cfg.Int(goconfig.DEFAULT_SECTION, key)
	return intValue
}

func GetBool(key string)(bool)  {
	boolValue, _ := cfg.Bool(goconfig.DEFAULT_SECTION, key)
	return boolValue
}



