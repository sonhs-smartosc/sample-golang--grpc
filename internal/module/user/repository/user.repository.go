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

func (userRepo *UserRepository) FindOneById(id uint64) (response *entities.User, err error) {
	err = userRepo.DB.Where(&entities.User{Id: id}).Take(&response).Error
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (userRepo *UserRepository) GetListUsers() (users []*entities.User, err error) {

	err = userRepo.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userRepo *UserRepository) UpdateUserById(user *entities.User) (response *entities.User, err error) {

	//err = userRepo.DB.Where(&entities.User{Id: user.Id}).Take(&response).Error
	result := userRepo.DB.Save(user)

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}

func (userRepo *UserRepository) DeleteUserById(id uint64) error {
	err := userRepo.DB.Delete(&entities.User{Id: id})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
