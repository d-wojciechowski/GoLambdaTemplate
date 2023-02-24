package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
	"net/http"
)

type MyEvent struct {
	Url string `json:"url"`
}

type MyResponse struct {
	Message string `json:"message:"`
}

func HandleRequest(ctx context.Context, event MyEvent) (*MyResponse, error) {
	resp, err := http.Get(event.Url)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return &MyResponse{Message: resp.Status}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
