package config

import (
	"github.com/lexkong/log"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
}

func Init(path string) error {
	var (
		c   *Config
		err error
	)
	c = &Config{Name: path}

	if err = c.initConfig(); err != nil {
		return err
	}

	// 初始化日志
	c.initLog()


	// 检测配置文件的修改
	//c.watchConfig()

	return nil
}

func (c *Config) initLog() {
	var (
		err error
	)
	/**
	writers：输出位置，有两个可选项 —— file 和 stdout。选择 file 会将日志记录到 logger_file 指定的日志文件中，选择 stdout 会将日志输出到标准输出，当然也可以两者同时选择
	logger_level：日志级别，DEBUG、INFO、WARN、ERROR、FATAL
	logger_file：日志文件
	log_format_text：日志的输出格式，JSON 或者 plaintext，true 会输出成非 JSON 格式，false 会输出成 JSON 格式
	rollingPolicy：rotate 依据，可选的有 daily 和 size。如果选 daily 则根据天进行转存，如果是 size 则根据大小进行转存
	log_rotate_date：rotate 转存时间，配 合rollingPolicy: daily 使用
	log_rotate_size：rotate 转存大小，配合 rollingPolicy: size 使用
	log_backup_count：当日志文件达到转存标准时，log 系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数
	 */
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	if err = log.InitWithConfig(&passLagerCfg); err != nil {
		panic(err)
	}
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
