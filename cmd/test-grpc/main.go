package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
	"test-Grpc/internal/database"
	userModule "test-Grpc/internal/module/user/grpc"
	"test-Grpc/proto/pb"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading dotenv file")
	}
	databaseInfo := &database.DatabaseInfo{
		Name:     os.Getenv("DB_NAME"),
		Driver:   os.Getenv("DB_DRIVER"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		SSLMode:  os.Getenv("DB_SSL_MODE"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
	}
	err = database.ConnectToDB(databaseInfo)

	if err != nil {
		log.Fatalf("connect database fail with error: %v", err.Error())
	}
}

func runGRPC(userServer pb.UserServiceServer, listener net.Listener) error {
	serverOptions := []grpc.ServerOption{
		//grpc.UnaryInterceptor(interceptor.Unary()),
	}
	grpcServer := grpc.NewServer(serverOptions...)
	//------START register services Server------
	pb.RegisterUserServiceServer(grpcServer, userServer)
	//------END register services Server------
	reflection.Register(grpcServer)
	log.Printf("Start GRPC server at %s", listener.Addr().String())
	return grpcServer.Serve(listener)
}

func main() {
	fmt.Println("Starting GRPC Backend test_grpc ....")
	host := os.Getenv("HOST")

	if runGrpc, _ := strconv.ParseBool(os.Getenv("RUN_GRPC_ENDPOINT")); runGrpc {
		port := os.Getenv("LISTEN_GRPC_PORT")

		address := fmt.Sprintf("%s:%s", host, port)
		listener, err := net.Listen("tcp", address)

		if err != nil {
			log.Fatalf("[Backend SmartContract Wallet] Cannot start server at address: %s", address)
		}

		userServer := userModule.NewUserGRPCServer()

		err = runGRPC(userServer, listener)

		if err != nil {
			log.Fatalf("%v", err)
		}
	}
}
