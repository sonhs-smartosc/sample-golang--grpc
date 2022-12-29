package dto

import "test-Grpc/internal/module/user/entities"

type User struct {
	Id       uint64
	Name     string
	Age      uint64
	Email    string
	Password string
	Phone    string
	//Email -> validate
	//Phone --> validate
	//Password --> validate
	//Age uint64
}

type CreateUserRequest struct {
	Name     string `validate:"required"`
	Age      uint64 `validate:"required"`
	Email    string `validate:"required,email,checkLocalPartEmail"`
	Password string `validate:"required"`
	Phone    string `validate:"required"`
}

type CreateUserResponse struct {
	User *entities.User
}

type FindUserById struct {
	Id uint64
}

type UpdateUserByIdRequest struct {
	Id    uint64 `validate:"required"`
	Name  string
	Age   uint64
	Phone string
}

type DeleteUserByIdResponse struct {
	Result bool
}
