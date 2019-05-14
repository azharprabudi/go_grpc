package controller

import (
	"fmt"
	"log"

	v1 "github.com/azharprabudi/go_grpc/api/v1/pb_go"
)

type TodoController struct{}

func (tc *TodoController) Post(stream v1.TodoService_PostServer) error {
	i := 0
	for i < 3 {
		_, err := stream.Recv()
		if err != nil {
			log.Println(err)
			stream.SendAndClose(&v1.Empty{})
		}

		i++
		fmt.Printf(`title: %s body: %s`, "kondel", "unch")
	}

	return nil
}

func NewTodoController() v1.TodoServiceServer {
	return &TodoController{}
}
