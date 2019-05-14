package main

import (
	"log"
	"net"

	v1 "github.com/azharprabudi/go_grpc/api/v1/pb_go"
	postController "github.com/azharprabudi/go_grpc/internal/go_grpc/controller/post"
	todoController "github.com/azharprabudi/go_grpc/internal/go_grpc/controller/todo"

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

	// add post service
	postCtl := postController.NewPostController()
	v1.RegisterPostServiceServer(server, postCtl)

	// add todo service
	todoCtl := todoController.NewTodoController()
	v1.RegisterTodoServiceServer(server, todoCtl)

	reflection.Register(server)

	if err = server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
