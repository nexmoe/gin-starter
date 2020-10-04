package setting

import (
	"log"

	"github.com/spf13/viper"
)

var (
	Cfg *viper.Viper
)

func init() {
	Cfg = viper.New()
	Cfg.AddConfigPath(".")
	Cfg.SetConfigName("config") // 读取配置文件
	if err := Cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}

}
