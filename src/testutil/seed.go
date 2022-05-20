package testutil

import "gorm.io/gorm"

func Seed(db *gorm.DB, seeds []interface{}) error {
	if seeds == nil {
		return nil
	}
	for _, s := range seeds {
		if err := db.Create(s).Error; err != nil {
			return err
		}
	}
	return nil
}
