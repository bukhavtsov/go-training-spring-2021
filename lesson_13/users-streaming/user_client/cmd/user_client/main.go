package main

import (
	"context"
	"log"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_13/users-streaming/proto/go_proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8282", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	stream, err := pb.NewUserServiceClient(conn).HelloUser(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	for {
		message, err := stream.Recv()
		if err != nil {
			log.Fatal("got an error:", err)
		}

		log.Println("got a message:", message)

		err = stream.Send(&pb.HelloRequest{UserName: "Artsiom"})
		if err != nil {
			log.Fatal(err)
		}
	}
}
