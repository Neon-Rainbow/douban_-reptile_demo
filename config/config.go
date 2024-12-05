package config

import (
	"fmt"
	"sync"

	"github.com/spf13/viper"
)

var (
	// cfg 配置
	cfg Cfg

	// once 保证只读取一次配置文件
	once sync.Once
)

// Cfg 配置
type Cfg struct {
	Url        string `mapstructure:"url"`
	Goroutines int    `mapstructure:"goroutines"`
	OutputPath string `mapstructure:"outputPath"`
}

// initConfig 读取配置文件
func initConfig() {
	viper.SetConfigFile("config/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file")
	}

	viper.SetDefault("url", "https://movie.douban.com/cinema/nowplaying/shanghai/")
	viper.SetDefault("goroutines", 10)
	viper.SetDefault("outputPath", "./output")

	cfg = Cfg{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		fmt.Println("Error unmarshalling config")
	}
}

// GetConfig 获取配置
// 通过once.Do保证只读取一次配置文件
func GetConfig() *Cfg {
	once.Do(func() {
		initConfig()
	})
	return &cfg
}
