package api

import (
	"context"
	"errors"
	"fmt"
	"log"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_13,14/users/proto/go_proto"
)

type UserServer struct{}

func NewUserServer() *UserServer{
	return &UserServer{}
}

func (u UserServer) HelloUser(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received username: %s", request.UserName)
	if request.UserName == "" {
		return nil, errors.New("empty username")
	}
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello %s", request.UserName)}, nil
}
