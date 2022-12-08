package config

// ServerConfig 服务配置
type ServerConfig struct {
	Name      string      `mapstructure:"name"`            // 项目名
	Port      int         `mapstructure:"port"`            // 端口
	Page      int         `mapstructure:"page"`            // 页数/第几页
	PageSize  int         `mapstructure:"pageSize"`        // 每页的数量
	RateLimit int         `mapstructure:"rateLimit"`       // 限制并发请求数量（每秒1万次请求）
	RaLiSw    bool        `mapstructure:"rateLimitSwitch"` // 是否开启限流（默认关闭限流）
	Language  LanguageCfg `mapstructure:"language"`        // 语言相关配置
	MysqlInfo MysqlConfig `mapstructure:"mysql"`           // MySQL配置
	RedisInfo RedisConfig `mapstructure:"redis"`           // Redis配置
	LogsInfo  LogConfig   `mapstructure:"logs"`            // 日志配置
	JWTKey    JWTConfig   `mapstructure:"jwt"`             // JWT配置
	AdminInfo AdminConfig `mapstructure:"adminaccount"`    // 管理员账户配置
	UserInfo  UserConfig  `mapstructure:"user"`            // 用户信息配置
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
	Host     string `mapstructure:"host"`     // 缓存数据库地址
	Port     int    `mapstructure:"port"`     // 端口
	Password string `mapstructure:"password"` // 密码
}

// JWTConfig JWT配置
type JWTConfig struct {
	SigningKey      string `mapstructure:"key"`             // Token密钥
	TokenExpiration int    `mapstructure:"tokenExpiration"` // Token过期时间（默认6小时，60*60*6=21600）
}

// AdminConfig 管理员账户配置
type AdminConfig struct {
	UserName string `mapstructure:"userName"` // 默认用户的用户名
	Password string `mapstructure:"password"` // 默认用户的密码
}

// UserConfig 用户信息配置
type UserConfig struct {
	PwdEncDiff   int  `mapstructure:"pwdEncDiff"`   // 密码加密难度（4~31，默认10）
	CaptchaLogin bool `mapstructure:"captchaLogin"` // 是否开启验证码登录
	IpFrequency  int  `mapstructure:"ipFrequency"`  // IP访问频率或者是次数
	TimeLimit    int  `mapstructure:"timeLimit"`    // 时长（默认限制1分钟只能访问100次）
}

// LanguageCfg 语言相关配置
type LanguageCfg struct {
	LanguageType string `mapstructure:"languageType"` // 语言类型（默认中文）
	Tranfilepath string `mapstructure:"tranfilepath"` // 翻译文件位置
	AllLanguage  string `mapstructure:"allLanguage"`  // 项目会使用到的全部语言
}
