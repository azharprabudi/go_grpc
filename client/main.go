package main

import (
	"context"
	"log"

	v1 "github.com/azharprabudi/go_grpc/api/v1/pb_go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:6969", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	client := v1.NewTodoServiceClient(conn)

	call, err := client.Post(ctx)
	if err != nil {
		log.Fatal(err)
	}

	call.Send(&v1.Todo{
		Body:  "kondel",
		Title: "kondel123",
	})

	call.Send(&v1.Todo{
		Body:  "kondel",
		Title: "kondel123",
	})

	call.CloseSend()
}
