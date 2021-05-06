package main

import (
	"google.golang.org/grpc/reflection"
	"log"
	"net"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_15/users/proto/go_proto"
	"github.com/bukhavtsov/go-training-spring-2021/lesson_15/users/user_server/pkg/api"

	"google.golang.org/grpc"

)

const port = ":8080"

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	reflection.Register(server)
	pb.RegisterUserServiceServer(server, api.NewUserServer())

	if err = server.Serve(listener); err != nil {
		log.Fatal(err)
	}
}
