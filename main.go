package main

import (
	"context"

	"github.com/Rick-Palomus-LLC/library/types"
	"github.com/Rick-Palomus-LLC/library/util/StripeUtil"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, payment types.PaymentEvent) (types.PaymentResponse, error) {
	charge, err := StripeUtil.ChargeCard(payment.Token)

	response := types.PaymentResponse{}

	if err != nil {
		response.Message = err.Error()
		response.Ok = false
	} else {
		response.Ok = true
		response.Message = charge.Status
	}

	return response, err
}

func main() {
	lambda.Start(HandleRequest)
}
