package productstorage

import (
	"gorm.io/gorm"
)

type productSQLStorage struct {
	db *gorm.DB
}

func NewProductSQLStorage(db *gorm.DB) *productSQLStorage {
	return &productSQLStorage{db: db}
}
