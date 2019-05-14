package service

import (
	"encoding/json"

	v1 "github.com/azharprabudi/go_grpc/api/v1/pb_go"
	fazzcloud "github.com/payfazz/go-apt/pkg/fazzcloud"
)

type PostService struct {
	httpClient *fazzcloud.HTTPClient
}

type PostServiceInterface interface {
	Get() (*v1.Posts, error)
}

func (ps *PostService) Get() (*v1.Posts, error) {
	_, res, err := ps.httpClient.Get("/", nil, &map[string]string{
		"Content-Type": "application/json",
		"Accept":       "application/json",
	})
	if err != nil {
		return nil, err
	}

	results := []*v1.Post{}
	if err := json.Unmarshal(res, &results); err != nil {
		return nil, err
	}

	return &v1.Posts{
		Post: results,
	}, nil
}

func NewPostService() PostServiceInterface {
	return &PostService{
		httpClient: fazzcloud.NewHTTPClient("https://jsonplaceholder.typicode.com/posts"),
	}
}
