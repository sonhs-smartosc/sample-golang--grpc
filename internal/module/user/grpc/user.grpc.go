package grpc

import (
	"context"
	"test-Grpc/internal/module/user/services"
	"test-Grpc/internal/module/user/transformers"
	"test-Grpc/proto/pb"
)

type UserGrpcServer struct {
	UseService      *services.UserService
	UserTransformer *transformers.UserTransformer
}

func NewUserGRPCServer() *UserGrpcServer {
	return &UserGrpcServer{
		UseService:      services.NewUserService(),
		UserTransformer: transformers.NewUserTransformer(),
	}
}

func (userGrpcServer *UserGrpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.User, error) {

	//request, err := userGrpcServer.UserTransformer.CreateUserRequestPbToDto(req)
	//*dto.user
	request, err := userGrpcServer.UserTransformer.CreateUserRequestPbToDto(req)

	if err != nil {
		return nil, err
	}

	err = userGrpcServer.UseService.CreateUser(ctx, request)

	if err != nil {
		return nil, err
	}

	return &pb.User{
		Id:   request.Id,
		Name: request.Name,
		Age:  request.Age,
	}, nil
}
