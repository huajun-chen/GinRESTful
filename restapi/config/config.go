package config

// ServerConfig 服务配置
type ServerConfig struct {
	Name        string      `mapstructure:"name"`
	Port        int         `mapstructure:"port"`
	MysqlInfo   MysqlConfig `mapstructure:"mysql"`
	RedisInfo   RedisConfig `mapstructure:"redis"`
	LogsAddress string      `mapstructure:"logsAddress"`
	JWTKey      JWTConfig   `mapstructure:"jwt"`
}

// MysqlConfig MySQL配置
type MysqlConfig struct {
	Host         string `mapstructure:"host"`         // 数据库地址
	Port         int    `mapstructure:"port"`         // 数据库端口
	Name         string `mapstructure:"name"`         // 用户名
	Password     string `mapstructure:"password"`     // 密码
	DBName       string `mapstructure:"dbName"`       // 数据库名
	MaxIdleConns int    `mapstructure:"maxIdleConns"` // 最大空闲连接数
	MaxOpenConns int    `mapstructure:"maxOpenConns"` // 最大打开连接数
}

// RedisConfig Redis配置
type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
}

// JWTConfig JWT签名密钥
type JWTConfig struct {
	SigningKey string `mapstructure:"key"`
}
