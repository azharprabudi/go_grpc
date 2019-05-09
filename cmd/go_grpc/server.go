package main

import (
	"log"
	"net"

	v1 "github.com/azharprabudi/go-grpc/api/v1/pb_go"
	controller "github.com/azharprabudi/go-grpc/internal/go_grpc/controller/post"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type serviceServer struct{}

func main() {
	lis, err := net.Listen("tcp", ":6969")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()

	// register service
	postCtl := controller.NewPostController()
	v1.RegisterPostServiceServer(server, postCtl)

	reflection.Register(server)

	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
