package main

import (
	"context"

	"github.com/Rick-Palomus-LLC/library/types"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, payment types.PaymentEvent) (types.PaymentResponse, error) {

	return types.PaymentResponse{
		Message: "ok",
		Ok:      true,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
