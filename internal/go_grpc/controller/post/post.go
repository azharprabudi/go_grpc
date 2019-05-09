package controller

import (
	"context"

	v1 "github.com/azharprabudi/go-grpc/api/v1/pb_go"
	service "github.com/azharprabudi/go-grpc/internal/go_grpc/service/post"
)

type PostController struct {
	service service.PostServiceInterface
}

func (pc *PostController) Get(ctx context.Context, in *v1.Empty) (*v1.Posts, error) {
	return pc.service.Get()
}

func NewPostController() v1.PostServiceServer {
	return &PostController{
		service: service.NewPostService(),
	}
}
