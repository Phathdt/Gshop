package userstorage

import (
	"gorm.io/gorm"
)

type userSQLStorage struct {
	db *gorm.DB
}

func NewUserSQLStorage(db *gorm.DB) *userSQLStorage {
	return &userSQLStorage{db: db}
}
