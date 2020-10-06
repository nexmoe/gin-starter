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
	Cfg.SetConfigName("config") // Read the file of config
	if err := Cfg.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("[warning] no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("[warning] read config error")
		}
		log.Fatal(err) // Deadly error
	}

}
