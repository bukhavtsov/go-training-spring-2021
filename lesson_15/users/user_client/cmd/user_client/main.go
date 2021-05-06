package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_15/users/proto/go_proto"

	"google.golang.org/grpc"

)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)
	username := "bukhavtsov"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.HelloUser(ctx, &pb.HelloRequest{UserName: username})
	if err != nil {
		if err, ok := status.FromError(err); ok {
			if err.Code() == codes.PermissionDenied {
				errMsg := fmt.Sprintf("code: %s, msg: %s, details :%v ", err.Code(), err.Message(), err.Details())
				log.Fatal(errMsg)
			} else {
				log.Fatal("undefined code, err is:", err)
			}
		}
		log.Fatal(err)
	}

	log.Printf("Got message from server: %s", r.Message)
}
