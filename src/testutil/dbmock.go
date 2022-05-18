package testutil

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDBMock() (*gorm.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		return nil, nil, err
	}

	gormDB, err := gorm.Open(
		mysql.Dialector{
			Config: &mysql.Config{
				DriverName:                "mysql",
				Conn:                      db,
				SkipInitializeWithVersion: true},
		},
		&gorm.Config{},
	)
	if err != nil {
		return nil, nil, err
	}
	return gormDB, mock, nil
}
