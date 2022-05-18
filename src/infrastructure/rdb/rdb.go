package rdb

import (
	"fmt"
	"go-clean-architecture/src/config"
	"go-clean-architecture/src/infrastructure/logger"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var dbInstance *gorm.DB

func NewRDB() *gorm.DB {
	if dbInstance != nil {
		return dbInstance
	}
	dbConf := config.Conf.Db
	conn, err := gorm.Open(mysql.Open(connString(dbConf)), &gorm.Config{
		Logger: newLogger(),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
	})
	if err != nil {
		logger.Fatalf("db connection apperror: %v", err)
	} else {
		logger.Info("success to connect db!")
		dbConfig, _ := conn.DB()
		dbConfig.SetMaxOpenConns(dbConf.MaxOpenConnection)
		dbConfig.SetMaxIdleConns(dbConf.MaxIdleConnection)
		dbConfig.SetConnMaxLifetime(time.Hour * time.Duration(dbConf.ConnectionMaxLifetimeHour))
	}

	dbInstance = conn
	return dbInstance
}

func connString(dbConf config.Db) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)
}

func newLogger() gormlogger.Interface {
	loggerConfig := gormlogger.Config{
		SlowThreshold: time.Second,
		Colorful:      false,
	}
	if config.IsDev() {
		loggerConfig.LogLevel = gormlogger.Info
	} else {
		loggerConfig.LogLevel = gormlogger.Silent
	}

	return gormlogger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		loggerConfig,
	)
}
