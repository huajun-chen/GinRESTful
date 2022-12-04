package initialize

import (
	"GinRESTful/restapi/global"
	"GinRESTful/restapi/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// InitMysqlDB 初始化MySQL
func InitMysqlDB() {
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
}
