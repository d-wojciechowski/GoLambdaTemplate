package main

import (
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Url string `json:"url"`
}

type MyResponse struct {
	StatusCode int               `json:"statusCode"`
	Headers    map[string]string `json:"headers"`
	Body       string            `json:"body"`
	Base64     bool              `json:"isBase64Encoded"`
}

func HandleRequest(ctx context.Context, event MyEvent) (MyResponse, error) {
	return MyResponse{
		StatusCode: 200,
		Headers:    map[string]string{"Access-Control-Allow-Origin": "*"},
		Body:       "Keep being awesome Cloud Gurus!",
		Base64:     false,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
