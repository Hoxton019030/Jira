package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

type MysqlConfig struct {
	Host                  string        `yaml:"host"`
	Port                  int           `yaml:"port"`
	Database              string        `yaml:"database"`
	Username              string        `yaml:"username"`
	Password              string        `yaml:"password"`
	MaxOpenConnections    int           `yaml:"max_open_connections"`
	MaxIdleConnections    int           `yaml:"max_idle_connections"`
	MaxConnectionLifetime time.Duration `yaml:"max_connection_lifetime"`
}

func Init() {
	log.Println("Viper啟動成功")
	viper.SetConfigName("config/config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
