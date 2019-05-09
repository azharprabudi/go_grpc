package main

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/azharprabudi/go-grpc/api/v1/pb_go"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:6969", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := v1.NewPostServiceClient(conn)

	ctx := context.Background()
	resp, err := client.Get(ctx, &v1.Empty{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)
}
