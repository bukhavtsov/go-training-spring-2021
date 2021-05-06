package main

import (
	"log"
	"net"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_13, 14/users-streaming/proto/go_proto"
	"github.com/bukhavtsov/go-training-spring-2021/lesson_13, 14/users-streaming/user_server/pkg/api"

	"google.golang.org/grpc"

)

const port = ":8282"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	pb.RegisterUserServiceServer(server, api.NewUserServer())

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
