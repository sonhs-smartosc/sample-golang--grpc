package transformers

import (
	"strings"
	"test-Grpc/internal/module/user/dto"
	"test-Grpc/proto/pb"
)

type UserTransformer struct {
}

func NewUserTransformer() *UserTransformer {
	return &UserTransformer{}
}

//func (transformer *UserTransformer) CreateUserRequestPbToDto(data *pb.CreateUserRequest) (result *dto.CreateUserRequest, err error) {
//	result = &dto.CreateUserRequest{
//		Name: strings.Trim(data.Name, " "),
//		Age:  data.Age,
//	}
//	return result, err
//}

func (transformer *UserTransformer) CreateUserRequestPbToDto(data *pb.CreateUserRequest) (result *dto.CreateUserRequest, err error) {

	result = &dto.CreateUserRequest{
		Name:     strings.Trim(data.Name, " "),
		Age:      data.Age,
		Email:    data.Email,
		Password: data.Password,
		Phone:    data.Phone,
	}
	return result, err
}

func (transformer *UserTransformer) UserDtoToPb(data *dto.User) (result *pb.User) {
	result = &pb.User{
		Id:    data.Id,
		Name:  data.Name,
		Age:   data.Age,
		Email: data.Email,
		Phone: data.Phone,
	}
	return result
}

func (transformer *UserTransformer) UpdateUserPbToDto(data *pb.UpdateUserByIdRequest) (result *dto.UpdateUserByIdRequest, err error) {
	result = &dto.UpdateUserByIdRequest{
		Id:    data.Id,
		Name:  data.Name,
		Age:   data.Age,
		Phone: data.Phone,
	}
	return result, nil
}
