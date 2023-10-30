package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

/*
	Description: init mysql
*/

// gorm初始化
//var DB *gorm.DB

func InitGorm(dsn string) *gorm.DB {
	// 将日志写进kafka
	// logx.SetWriter(*LogxKafka())
	var err error
	DB, err := gorm.Open(mysql.Open(dsn),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
			},
			Logger: logger.Default.LogMode(logger.Silent),
		})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		fmt.Println("连接mysql数据库成功")
	}
	db, _ := DB.DB()
	db.SetMaxIdleConns(50)
	db.SetMaxOpenConns(500)
	db.SetConnMaxLifetime(time.Minute)
	return DB
}
