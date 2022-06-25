package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	CreateUser(user User) (User, error)
	VerifyUser(digest string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func UserRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r repository) CreateUser(user User) (User, error) {
	err := r.db.Create(&user).Error

	return user, err
}

func (r repository) VerifyUser(digest string) (User, error) {
	var user User

	err := r.db.Where("digest_user_auth = ?", digest).First(&user).Error

	return user, err
}
