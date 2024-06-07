package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"wu-user-backend/utils/pkg/constants"
)

var DB *gorm.DB

func SqlInit() {
	var err error
	DB, err = gorm.Open(mysql.Open(constants.MySQLDefaultDSN), &gorm.Config{
		SkipDefaultTransaction: true,                                //跳过默认事务
		Logger:                 logger.Default.LogMode(logger.Info), //设置日志模式
	})
	if err != nil {
		panic(err)
	}
}
