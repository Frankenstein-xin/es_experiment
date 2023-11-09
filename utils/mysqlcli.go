package utils

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var mysqlCli *gorm.DB

func MustInitMysqlClient(conf *MysqlConfig) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)

	var err error
	mysqlCli, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func GetMysqlCli() *gorm.DB {
	return mysqlCli
}
