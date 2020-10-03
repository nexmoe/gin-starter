package setting

import (
	"log"
	"time"

	"github.com/spf13/viper"
)

var (
	RunMode string

	HTTPPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config") // 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			log.Println("no such config file")
		} else {
			// Config file was found but another error was produced
			log.Println("read config error")
		}
		log.Fatal(err) // 读取配置文件失败致命错误
	}

	LoadBase()
}

func LoadBase() {
	RunMode = viper.GetString("run_mode")
	log.Println(RunMode)
}

func LoadServer() {
	HTTPPort = viper.GetInt("server.port")
	ReadTimeout = viper.GetDuration("server.read_timeout")
	WriteTimeout = viper.GetDuration("server.write_timeout")
}

func GetConfig (key string) {
	viper.Get(key)
}