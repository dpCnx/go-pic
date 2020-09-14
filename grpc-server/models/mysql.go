package models

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"grpc-server/conf"
)

var (
	_   *gorm.DB
	err error
)

func init() {

	_, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		conf.C.MySQL.User,
		conf.C.MySQL.Password,
		conf.C.MySQL.IP+":"+conf.C.MySQL.Port,
		conf.C.MySQL.Database)), &gorm.Config{NamingStrategy: schema.NamingStrategy{
		SingularTable: true,
	}})
	if err != nil {
		panic(err)
	}

	zap.L().Info("pic-server mysql start...")
}

/*func CloseDb() {
	if mysqlDb != nil {
		_ = mysqlDb.Close()
	}
}*/
