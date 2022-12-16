package test

import (
	"GinRESTful/restapi/config"
	"GinRESTful/restapi/dao"
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/models"
	"GinRESTful/restapi/utils"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitTestBase 初始化基础的测试环境
// 参数：
//		无
// 返回值：
//		无
func InitTestBase() {
	// 1.初始化YAML配置
	// 实例化viper
	v := viper.New()
	// 文件的路径设置
	v.SetConfigFile("../../setting-dev.yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(err)
	}
	serverConfig := config.ServerConfig{}
	// 给serverConfig初始化值
	err = v.Unmarshal(&serverConfig)
	if err != nil {
		panic(err)
	}
	// 传递全局变量
	global.Settings = serverConfig

	// 2.初始化MySQL
	mysqlInfo := global.Settings.MysqlInfo
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlInfo.Name,
		mysqlInfo.Password,
		mysqlInfo.Host,
		mysqlInfo.Port,
		mysqlInfo.DBName,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	// 设置连接池
	sqlDB.SetMaxIdleConns(mysqlInfo.MaxIdleConns) // 空闲
	sqlDB.SetMaxOpenConns(mysqlInfo.MaxOpenConns) // 打开
	global.DB = db
	// MYSQL数据迁移
	utils.Migration()

	// 3.初始化Redis
	redisInfo := global.Settings.RedisInfo
	redisAddr := fmt.Sprintf("%s:%d",
		redisInfo.Host,
		redisInfo.Port,
	)
	// 生成Redis客户端
	global.Redis = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisInfo.Password,
		DB:       0,
	})
	// 连接Redis
	_, err = global.Redis.Ping().Result()
	if err != nil {
		panic(err)
	}

	// 4.初始化admin账户
	// 默认配置的管理员用户名
	adminInfo := global.Settings.AdminInfo
	// 查询admin是否存在
	_, ok := dao.DaoFindUserInfoToUserName(adminInfo.UserName)
	// 不存在，创建
	if !ok {
		// 加密密码
		pwdStr, err := utils.SetPassword(adminInfo.Password)
		if err != nil {
			panic(err)
		}
		// 创建管理员账户
		user := models.User{UserName: adminInfo.UserName, Password: pwdStr, Role: 1}
		if err := global.DB.Create(&user).Error; err != nil {
			panic(err)
		}
	}

	// 初始化i18n
	// 读取zh-CN.json
	translations, err := utils.ReadJSON("../../static/i18n/zh-CN.json")
	if err != nil {
		panic(err)
	}
	// 传递全局变量
	global.I18nMap = translations
}
