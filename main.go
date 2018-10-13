package main

import (
	"context"
	"fmt"

	"github.com/Rick-Palomus-LLC/library/types"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, payment types.PaymentEvent) (string, error) {
	return fmt.Sprintf("%s", payment.Name), nil
}

func main() {
	lambda.Start(HandleRequest)
}
