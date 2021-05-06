package api

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_15/users/proto/go_proto"
)

type UserServer struct{}

func NewUserServer() *UserServer {
	return &UserServer{}
}

var users = []string{"Alex", "Dmitry D", "Dmitry P", "Daniil", "Nikita", "Eugene", "Anton", "Artsiom"}

func isUserExists(username string) bool {
	for _, user := range users {
		if user == username {
			return true
		}
	}
	return false
}

func (u UserServer) HelloUser(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	if (!isUserExists(request.UserName)) {
		s := status.Newf(codes.PermissionDenied, "unauthorized user %s", request.UserName)
		errWithDetails, err := s.WithDetails(request)
		if err != nil {
			return nil, status.Errorf(codes.Unknown, "can't covert status to status with details %v", s)
		}
		return nil, errWithDetails.Err()
	}
	log.Printf("Received username: %s", request.UserName)
	return &pb.HelloResponse{Message: fmt.Sprintf("Hello %s", request.UserName)}, nil
}
