package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "grpc-example/user"
)

type server struct {
	pb.UnimplementedUserServerServer
}

func (s *server) GetUserList(ctx context.Context, in *pb.Request) (*pb.Users, error) {
	return &pb.Users{
		UserList: []*pb.User{
			{
				Name: "小明",
				Age:  18,
				City: "WuHan",
			},
			{
				Name: "小红",
				Age:  16,
				City: "ShangHai",
			},
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "localhost:3000")

	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()

	pb.RegisterUserServerServer(grpcServer, &server{})

	grpcServer.Serve(lis)
}
