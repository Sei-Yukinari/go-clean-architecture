package testutil

import (
	"fmt"
	"go-clean-architecture/src/infrastructure/rdb"
	"gorm.io/gorm"
	"testing"
)

func SetupRDB(t *testing.T) *gorm.DB {
	t.Helper()
	db := rdb.NewRDB()
	tx := db.Begin()
	t.Cleanup(func() {
		fmt.Println("Rollback!!")
		tx.Rollback()
	})
	return tx
}
