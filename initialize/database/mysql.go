package database

import (
	"bytes"
	"github.com/tigercandy/prado/global"
	"github.com/tigercandy/prado/global/orm"
	plogger "github.com/tigercandy/prado/internal/pkg/logger"
	"gopkg.in/natefinch/lumberjack.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"io"
	"log"
	"os"
	"strconv"
	"time"
)

type Mysql struct {
}

func InitMysqlGorm() *gorm.DB {
	dbConfig := global.App.Config.Database
	if dbConfig.DBName == "" {
		return nil
	}
	var db Database
	db = new(Mysql)
	orm.MysqlConn = db.GetConnect()
	mysqlConfig := mysql.Config{
		DSN:                       orm.MysqlConn,
		DefaultStringSize:         191,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}
	if db, err := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   getGormLogger(),
	}); err != nil {
		plogger.Errorf("mysql connect failed, err: %s", err)
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(dbConfig.MaxIdleConn)
		sqlDB.SetMaxOpenConns(dbConfig.MaxOpenConn)
		return db
	}
}

func (e *Mysql) GetConnect() string {
	dbConfig := global.App.Config.Database
	var conn bytes.Buffer
	conn.WriteString(dbConfig.UserName)
	conn.WriteString(":")
	conn.WriteString(dbConfig.Password)
	conn.WriteString("@tcp(")
	conn.WriteString(dbConfig.Host)
	conn.WriteString(":")
	conn.WriteString(strconv.Itoa(dbConfig.Port))
	conn.WriteString(")/")
	conn.WriteString(dbConfig.DBName)
	conn.WriteString("?charset=")
	conn.WriteString(dbConfig.Charset)
	conn.WriteString("&parseTime=True&loc=Local")

	return conn.String()
}

func getGormLogger() logger.Interface {
	var logMode logger.LogLevel

	switch global.App.Config.Database.LogMode {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logMode,
		IgnoreRecordNotFoundError: false,
		Colorful:                  !global.App.Config.Database.EnableFileLogWriter,
	})
}

func getGormLogWriter() logger.Writer {
	var writer io.Writer
	if global.App.Config.Database.EnableFileLogWriter {
		writer = &lumberjack.Logger{
			Filename:   global.App.Config.Log.Path + "/" + global.App.Config.Database.LogFilename,
			MaxSize:    global.App.Config.Log.MaxSize,
			MaxBackups: global.App.Config.Log.MaxBackups,
			MaxAge:     global.App.Config.Log.MaxAge,
			Compress:   global.App.Config.Log.Compress,
		}
	} else {
		writer = os.Stdout
	}

	return log.New(writer, "\r\n", log.LstdFlags)
}
