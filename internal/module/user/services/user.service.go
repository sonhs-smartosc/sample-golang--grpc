package services

import (
	"context"
	"test-Grpc/internal/module/user/dto"
	"test-Grpc/internal/module/user/entities"
	"test-Grpc/internal/module/user/repository"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repository.NewUserRepository(),
	}
}

//func (userService *UserService) CreateUser(ctx context.Context, request *dto.CreateUserRequest) (*dto.User, error) {
//	user := entities.User{
//		Name: request.Name,
//		Age:  request.Age,
//	}
//	err := userService.UserRepo.Create(&user)
//
//	if err != nil {
//		return nil, err
//	}
//	return &dto.User{
//		Id:   user.Id,
//		Name: user.Name,
//		Age:  user.Age,
//	}, nil
//}

func (userService *UserService) CreateUser(ctx context.Context, request *dto.User) error {
	user := &entities.User{
		Name: request.Name,
		Age:  request.Age,
	}
	err := userService.UserRepo.Create(user)

	if err != "postgres error: email da ton tai!" {
		return "Email da ton tai"
	}

	if err != nil {
		return err
	}

	request.Id = user.Id
	return nil
}
