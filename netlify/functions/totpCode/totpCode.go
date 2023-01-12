package main

import (
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"
)

func generateErrorResponse(err error) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode:      500,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            err.Error(),
		IsBase64Encoded: false,
	}, nil
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var key *otp.Key
	var err error

	uri, ok := request.QueryStringParameters["uri"]
	if !ok {
		err = fmt.Errorf("uri parameter not defined")
		return generateErrorResponse(err)
	}

	key, err = otp.NewKeyFromURL(uri)
	if err != nil {
		return generateErrorResponse(err)
	}

	code, err := totp.GenerateCodeCustom(key.Secret(), time.Now(), totp.ValidateOpts{
		Period:    uint(key.Period()),
		Digits:    key.Digits(),
		Algorithm: key.Algorithm(),
	})
	if err != nil {
		return generateErrorResponse(err)
	}

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE",
			"Content-Type":                 "text/plain",
			"Cache-Control":                "no-cache",
		},
		Body:            code,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
