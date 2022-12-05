package config

// ServerConfig 服务配置
type ServerConfig struct {
	Name      string      `mapstructure:"name"`
	Port      int         `mapstructure:"port"`
	MysqlInfo MysqlConfig `mapstructure:"mysql"`
	RedisInfo RedisConfig `mapstructure:"redis"`
	LogsInfo  LogConfig   `mapstructure:"logs"`
	JWTKey    JWTConfig   `mapstructure:"jwt"`
	AdminInfo AdminConfig `mapstructure:"adminaccount"`
	UserInfo  UserConfig  `mapstructure:"user"`
}

// LogConfig 日志配置
type LogConfig struct {
	Level      string `mapstructure:"level"`      // 日志等级（debug、info、warn、error、dpanic、panic、fatal）
	FileName   string `mapstructure:"fileName"`   // 日志文件的位置
	MaxSize    int    `mapstructure:"maxSize"`    // 在进行切割之前，日志文件的最大大小（以MB为单位）
	MaxAge     int    `mapstructure:"maxAge"`     // 保留旧文件的最大天数
	MaxBackups int    `mapstructure:"maxBackups"` // 保留旧文件的最大个数
	Compress   bool   `mapstructure:"compress"`   // 是否压缩/归档旧文件(默认为false，不压缩)
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

// AdminConfig 管理员账户配置
type AdminConfig struct {
	UserName string `mapstructure:"userName"`
	Password string `mapstructure:"password"`
}

// UserConfig 用户信息配置
type UserConfig struct {
	PwdEncDiff int `mapstructure:"pwdEncDiff"` // 密码加密难度（4~31，默认10）
}
