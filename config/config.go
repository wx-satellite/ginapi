package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}


func Init(path string) error {
	var (
		c *Config
		err error
	)
	c = &Config{Name:path}

	if err = c.initConfig(); err != nil {
		return err
	}


	// 初始化日志


	// 检测配置文件的修改
	//c.watchConfig()

	return nil
}


func (c *Config) initConfig() error {
	var (
		err error
	)
	// 如果指定了配置文件就直接读取
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		// 没有指定配置文件，则默认取当前目录的配置文件
		//viper.SetConfigName("config")
		//viper.AddConfigPath("./")
		viper.SetConfigFile("./config.yaml")

	}

	// 指定配置文件的后缀
	viper.SetConfigType("yaml")

	// 读取配置
	if err = viper.ReadInConfig(); err != nil {
		return err
	}



	return nil
}


// 热加载，监控配置文件的修改
//func (c *Config) watchConfig() {
//	viper.WatchConfig()
//	viper.OnConfigChange(func(event fsnotify.Event){
//		log.Printf("config file changed: %s", event.Name)
//	})
//}