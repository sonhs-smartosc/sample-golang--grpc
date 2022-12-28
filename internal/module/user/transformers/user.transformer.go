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

func (transformer *UserTransformer) CreateUserRequestPbToDto(data *pb.CreateUserRequest) (result *dto.User, err error) {

	result = &dto.User{
		Name: strings.Trim(data.Name, " "),
		Age:  data.Age,
	}
	return result, err
}
