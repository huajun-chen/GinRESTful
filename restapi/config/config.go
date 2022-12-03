package config

// ServerConfig 服务配置
type ServerConfig struct {
	Name        string      `yaml:"name"`
	Port        int         `yaml:"port"`
	MysqlInfo   MysqlConfig `yaml:"mysql"`
	RedisInfo   RedisConfig `yaml:"redis"`
	LogsAddress string      `yaml:"logsAddress"`
}

// MysqlConfig MySQL配置
type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Name     string `yaml:"name"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbName"`
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
}
