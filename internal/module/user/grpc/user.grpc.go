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

func (userGrpcServer *UserGrpcServer) DeleteUserById(ctx context.Context, request *pb.DeleteUserByIdRequest) (*pb.DeleteUserByIdResponse, error) {
	//TODO implement me
	//panic("implement me")
	response, err := userGrpcServer.UseService.DeleteUserById(ctx, request.Id)

	if err != nil {
		return nil, err
	}

	return &pb.DeleteUserByIdResponse{
		Result: response.Result,
	}, nil
}

func (userGrpcServer *UserGrpcServer) mustEmbedUnimplementedUserServiceServer() {
	//TODO implement me
	panic("implement me")
}

func (userGrpcServer *UserGrpcServer) UpdateUserById(ctx context.Context, request *pb.UpdateUserByIdRequest) (*pb.User, error) {
	//TODO implement me
	newUser, err := userGrpcServer.UserTransformer.UpdateUserPbToDto(request)
	if err != nil {
		return nil, err
	}
	user, err := userGrpcServer.UseService.UpdateUserByID(ctx, newUser)

	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:    user.Id,
		Name:  user.Name,
		Age:   user.Age,
		Email: user.Email,
		Phone: user.Phone,
	}, nil

	//dto -> pb.
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
	user, err := userGrpcServer.UseService.CreateUser(ctx, request)

	if err != nil {
		return nil, err
	}

	return userGrpcServer.UserTransformer.UserDtoToPb(user), nil
}

func (userGrpcServer *UserGrpcServer) GetUserById(ctx context.Context, request *pb.GetUserByIdRequest) (*pb.User, error) {
	id := request.Id

	user, err := userGrpcServer.UseService.GetUserByID(ctx, id)

	if err != nil {
		return nil, err
	}
	return userGrpcServer.UserTransformer.UserDtoToPb(user), nil
}

func (userGrpcServer *UserGrpcServer) GetUsers(ctx context.Context, request *pb.GetUsersRequest) (users *pb.ListUser, err error) {
	usersDto, err := userGrpcServer.UseService.GetListUser(ctx)

	if err != nil {
		return nil, err
	}

	var data []*pb.User
	for _, value := range usersDto {
		data = append(data, &pb.User{
			Id:    value.Id,
			Name:  value.Name,
			Age:   value.Age,
			Email: value.Email,
			Phone: value.Phone,
		})
	}

	return &pb.ListUser{Users: data}, nil
}

//func (userGrpcServer *UserGrpcServer)UpdateUserById(context.Context, *pb.UpdateUserByIdRequest) (*pb.User, error) {
//
//	UserGrpcServer
//	return nil, nil
//}
