package repository

import (
	"gorm.io/gorm"
	"test-Grpc/internal/database"
	"test-Grpc/internal/module/user/entities"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{DB: database.Connections}
	//return &UserRepository{DB: db.Connections[constant.DbBeWallet]}
}

func (userRepo *UserRepository) Create(user *entities.User) error {
	result := userRepo.DB.Create(user)

	if result != nil {
		return result.Error
	}

	return nil
}
