package config

import (
	"github.com/spf13/viper"
	"log"
	"path/filepath"
)

type Config struct {
	Domain string
}

var Conf Config

// 加载配置
func Init(){
	path,_:=filepath.Abs(".")
	v := viper.New()
	v.SetConfigFile(path+"/config/config.yaml")
	v.SetConfigType("yaml")
	if err1 := v.ReadInConfig(); err1 != nil {
		log.Panic(err1.Error())
		return
	}
	Conf.Domain = v.GetString("web.domain")
	return
}