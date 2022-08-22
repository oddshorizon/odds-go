package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//
//  LaunchMysqlDB
//  @Description: 实例化 mysql db
//  @param url
//  @return gorm.DB
//  @return error
//
func LaunchMysqlDB(url string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.New(mysql.Config{
		//DSN:                       "root:oddsroot@tcp(120.27.236.164:3306)/space?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DSN:                       url,   // DSN data source name
		DefaultStringSize:         44,    // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	return db, err
}
