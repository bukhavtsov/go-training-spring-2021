package main

import (
	"context"
	"log"
	"time"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_12/users/proto/go_proto"

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
		log.Fatal(err)
	}

	log.Printf("Got message from server: %s", r.Message)
}
