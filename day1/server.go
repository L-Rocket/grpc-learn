package main

import (
	"context"
	"log"
	"net"

	pb "day1/grpc_gen/user"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("Go server 收到请求 id =", req.Id)

	return &pb.GetUserResponse{
		Name: "Alice",
		Age:  18,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{})

	log.Println("Go gRPC server listening on :50051")
	grpcServer.Serve(lis)
}
