package main

import (
	"context"
	"fmt"

	"github.com/Rick-Palomus-LLC/library/types"
	"github.com/Rick-Palomus-LLC/library/util/StripeUtil"
	"github.com/Rick-Palomus-LLC/library/util/TwilioUtil"
	"github.com/aws/aws-lambda-go/lambda"
)

func HandleRequest(ctx context.Context, payment types.PaymentEvent) (types.PaymentResponse, error) {
	charge, err := StripeUtil.ChargeCard(payment.Token)

	response := types.PaymentResponse{}

	if err != nil {
		return errorResponse(err.Error())
	} else {
		// Database insert

		smsBody := fmt.Sprintf("%s - %s", payment.Name, payment.Number)

		message, err := TwilioUtil.SendMessage(smsBody)

		if err != nil {
			// Refund charge
			return errorResponse(err.Error())
		}

		response.Ok = true
		response.Message = fmt.Sprintf("%s %s %s %s", charge.Status, message.ErrorCode, message.Status, message.Body)
	}

	return response, nil
}

func errorResponse(message string) (types.PaymentResponse, error) {
	return types.PaymentResponse{
		Message: message,
		Ok:      false,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
