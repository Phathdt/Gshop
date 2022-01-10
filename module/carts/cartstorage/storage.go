package cartstorage

import (
	"gorm.io/gorm"
)

type cartSQLStorage struct {
	db *gorm.DB
}

func NewCartSQLStorage(db *gorm.DB) *cartSQLStorage {
	return &cartSQLStorage{db: db}
}
