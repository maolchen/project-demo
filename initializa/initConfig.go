package initializa

import (
	"github.com/maolchen/krm-backend/config"
	"github.com/spf13/viper"
	"log"
)

//var LogCfg *config.LogConf

// Initializa 初始化配置，根据是否提供了配置文件路径选择从配置文件还是环境变量加载配置
func InitConfig(configFile string) error {

	//LogCfg = &config.LogConf{}

	if configFile != "" {
		viper.SetConfigType("yaml")
		viper.SetConfigFile(configFile)
		log.Printf("尝试读取配置文件: %s\n", configFile)
		if err := viper.ReadInConfig(); err != nil {
			return err
		}

		return parseYamlToCfg()
	} else {
		log.Println("未配置配置文件，从环境变量获取配置。。。。。。")
		return getEnvToCfg()
	}

}

// parseYamlToCfg 解析 YAML 到结构体
func parseYamlToCfg() error {
	if err := viper.Unmarshal(&config.Cfg); err != nil {
		return err
	}
	return nil
}

// getEnvToCfg 从环境变量中获取配置并赋值给结构体，并设置默认值
func getEnvToCfg() error {
	// 启用自动从环境变量读取
	viper.AutomaticEnv()
	// 设置环境变量的默认值
	viper.SetDefault("ADDRESS", ":8000")
	viper.SetDefault("DB_PATH", ".\\data\\app.db")
	viper.SetDefault("SECRET", "default_secret_key")
	viper.SetDefault("LOG_CONF_LOG_FILE", "")
	viper.SetDefault("LOG_CONF_MAX_AGE", 7)
	viper.SetDefault("LOG_CONF_MAX_BACKUPS", 5)
	viper.SetDefault("LOG_CONF_MAX_SIZE", 10)
	viper.SetDefault("LOG_CONF_COMPRESS", true)
	viper.SetDefault("LOG_CONF_LOG_LEVEL", "debug")
	viper.SetDefault("LOG_CONF_LOG_TYPE", "text")
	viper.SetDefault("JWT_EXPIRES", 30)

	config.Cfg.Address = viper.GetString("ADDRESS")
	config.Cfg.DbPath = viper.GetString("DB_PATH")
	config.Cfg.Secret = viper.GetString("SECRET")
	config.Cfg.JwtExpires = viper.GetInt64("JWT_EXPIRES")

	logConf := &config.LogConf{
		LogFile:    viper.GetString("LOG_CONF_LOG_FILE"),
		MaxAge:     viper.GetInt("LOG_CONF_MAX_AGE"),
		MaxBackups: viper.GetInt("LOG_CONF_MAX_BACKUPS"),
		MaxSize:    viper.GetInt("LOG_CONF_MAX_SIZE"),
		Compress:   viper.GetBool("LOG_CONF_COMPRESS"),
		LogLevel:   viper.GetString("LOG_CONF_LOG_LEVEL"),
		LogType:    viper.GetString("LOG_CONF_LOG_TYPE"),
	}

	config.Cfg.LogConf = logConf
	return nil
}
