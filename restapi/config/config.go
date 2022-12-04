package config

// ServerConfig 服务配置
type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
	RedisInfo   RedisConfig `mapstructure:"redis"`
	LogsAddress string      `mapstructure:"logsAddress"`
}

// MysqlConfig MySQL配置
type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Name     string `mapstructure:"name"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbName"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}
