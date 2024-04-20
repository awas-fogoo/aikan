package util

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

func WithTransaction(db *gorm.DB, fn func(*gorm.DB) error) error {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Fatalf(fmt.Sprintf("Database transaction failed: %v", r))
		} else if err := recover(); err != nil {
			tx.Rollback()
			log.Fatalf(fmt.Sprintf("Rollbacking transation should not result in an error: %v", err))
		} else {
			tx.Commit()
		}
	}()
	return fn(tx)
}
