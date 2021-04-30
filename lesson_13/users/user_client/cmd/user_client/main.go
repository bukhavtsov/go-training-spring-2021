package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/bukhavtsov/go-training-spring-2021/lesson_13/users/proto/go_proto"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8282", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	grpcMux := runtime.NewServeMux()
	err = pb.RegisterUserServiceHandler(context.Background(), grpcMux, conn)
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe("localhost:8383", grpcMux))
}
