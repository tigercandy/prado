package database

import (
	"github.com/tigercandy/prado/global"
	"gorm.io/gorm"
)

func Setup() *gorm.DB {
	dbDriver := global.App.Config.Database.Driver
	if dbDriver == "mysql" {
		return InitMysqlGorm()
	} else {
		return InitMysqlGorm()
	}
}