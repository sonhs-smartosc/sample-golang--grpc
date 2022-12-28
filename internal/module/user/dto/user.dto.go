package dto

import "test-Grpc/internal/module/user/entities"

type User struct {
	Id   uint64
	Name string
	//Email -> validate
	//Phone --> validate
	//Password --> validate

	Age uint64
}

type CreateUserRequest struct {
	Name string `validate:"required"`
	Age  uint64 `validate:"required"`
}

type CreateUserResponse struct {
	User *entities.User
}
