package main

import (
	"fmt"
	"context"
	"github.com/aws/aws-lambda-go/lambda"
)

type PaymentEvent struct {
	Name string `json:"name"`
	Number string `json:"number"`
	Token string `json:"token"`
}

func HandleRequest(ctx context.Context, payment PaymentEvent) (string, error) {
	return fmt.Sprintf("%s", payment.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}