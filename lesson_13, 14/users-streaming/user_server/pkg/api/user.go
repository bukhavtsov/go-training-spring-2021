package api

import (
	"io"
	"log"
	"time"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_13/users-streaming/proto/go_proto"
)

type UserServer struct{}

func NewUserServer() *UserServer{
	return &UserServer{}
}

func (u UserServer) HelloUser(stream pb.UserService_HelloUserServer) error {
	go func() {
		for {
			username, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Fatal(err)
			}
			log.Print(username)
		}
	}()
	for {
		if err := stream.Send(&pb.HelloResponse{Message: "test"}); err != nil {
			log.Fatal("got an error form client", err)
		}
		time.Sleep(time.Second)
	}
}
