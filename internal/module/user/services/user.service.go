package services

import (
	"context"
	"golang.org/x/crypto/bcrypt"
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

func (userService *UserService) CreateUser(ctx context.Context, request *dto.CreateUserRequest) (*dto.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	request.Password = string(hashedPassword)

	user := entities.User{
		Name:     request.Name,
		Age:      request.Age,
		Email:    request.Email,
		Password: request.Password,
		Phone:    request.Phone,
	}
	err = userService.UserRepo.Create(&user)

	if err != nil {
		return nil, err
	}

	return &dto.User{
		Id:    user.Id,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}

func (userService *UserService) GetUserByID(ctx context.Context, id uint64) (*dto.User, error) {
	user, err := userService.UserRepo.FindOneById(id)

	if err != nil {
		return nil, err
	}

	return &dto.User{
		Id:    user.Id,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
		Phone: user.Phone,
	}, nil
}

func (userService *UserService) GetListUser(ctx context.Context) (users []*dto.User, err error) {
	//var usersEntities []*entities.User
	usersEntities, err := userService.UserRepo.GetListUsers()

	if err != nil {
		return nil, err
	}
	///
	//var userTemp *dto.User
	for _, value := range usersEntities {
		userTemp := &dto.User{
			Id:    value.Id,
			Name:  value.Name,
			Age:   value.Age,
			Email: value.Email,
			Phone: value.Phone,
		}
		users = append(users, userTemp)
	}
	return users, nil
}

func (userService *UserService) UpdateUserByID(ctx context.Context, request *dto.UpdateUserByIdRequest) (user *dto.User, err error) {

	//err = userRepo.DB.Where(&entities.User{Email: email}).
	userUpdate, err := userService.UserRepo.FindOneById(request.Id)

	if err != nil {
		return nil, err
	}

	newUser := entities.User{
		Id:       request.Id,
		Name:     request.Name,
		Age:      request.Age,
		Email:    userUpdate.Email,
		Password: userUpdate.Password,
		Phone:    request.Phone,
	}

	_, err = userService.UserRepo.UpdateUserById(&newUser)

	if err != nil {
		return nil, err
	}

	return &dto.User{
		Id:       newUser.Id,
		Name:     newUser.Name,
		Age:      newUser.Age,
		Email:    newUser.Email,
		Password: newUser.Password,
		Phone:    newUser.Phone,
	}, nil
}

func (userService *UserService) DeleteUserById(ctx context.Context, id uint64) (response *dto.DeleteUserByIdResponse, err error) {
	err = userService.UserRepo.DeleteUserById(id)
	if err != nil {
		return nil, err
	}

	return &dto.DeleteUserByIdResponse{
		Result: true,
	}, nil
}

//func (userService *UserService) CreateUser(ctx context.Context, request *dto.User) error {
//	user := &entities.User{
//		Name: request.Name,
//		Age:  request.Age,
//	}
//	err := userService.UserRepo.Create(user)
//
//	if err != "postgres error: email da ton tai!" {
//		return "Email da ton tai"
//	}
//
//	if err != nil {
//		return err
//	}
//
//	request.Id = user.Id
//	return nil
//}
