package testutil

import (
	"github.com/DATA-DOG/go-txdb"
	"go-clean-architecture/src/config"
	"go-clean-architecture/src/infrastructure/rdb"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"testing"
)

func SetupRDB(name string, t *testing.T) (*gorm.DB, error) {
	dsn := rdb.ConnString(config.Conf.Db)
	txdb.Register(name, "mysql", dsn)
	d := mysql.New(mysql.Config{
		DriverName: name,
		DSN:        dsn,
	})
	tx, err := gorm.Open(d, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		t.Errorf("test rdb open error:%v\n", err)
	}
	t.Cleanup(func() {
		db, _ := tx.DB()
		db.Close()
	})
	return tx, nil
}
