package config

// LogConf 日志相关配置
type LogConf struct {
	LogFile    string `yaml:"log_file"`
	MaxAge     int    `yaml:"max_age"`
	MaxBackups int    `yaml:"max_backups"` // 最大保留日志数量
	MaxSize    int    `yaml:"max_size"`    // 每个日志文件最大大小（MB）
	Compress   bool   `yaml:"compress"`    // 是否压缩旧日志
	LogLevel   string `yaml:"log_level"`   // 日志级别
	LogType    string `yaml:"log_type"`    // 日志格式
}

// Conf 整体配置结构
type Conf struct {
	Address    string   `yaml:"address"`                          // 服务监听地址
	DbPath     string   `yaml:"db_path"`                          // SQLite数据库路径
	Secret     string   `yaml:"secret"`                           // JWT密钥
	LogConf    *LogConf `yaml:"log_conf" mapstructure:"log_conf"` // 日志配置
	JwtExpires int64    `yaml:"jwt_expires"  //jwt配置`
}

var Cfg = &Conf{}
