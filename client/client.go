package main

import (
	"context"
	"fmt"
	pb "grpc-example/user"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure(), grpc.WithBlock()) //grpc.WithInsecure(), grpc.WithBlock()

	if err != nil {
		log.Fatalln(err)
	}

	defer conn.Close()

	c := pb.NewUserServerClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second) //返回一个client，并设置超时时间
	defer cancel()

	res, err := c.GetUserList(ctx, &pb.Request{})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
