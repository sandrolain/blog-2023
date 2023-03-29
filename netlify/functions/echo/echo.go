package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func generateErrorResponse(err error) (*events.APIGatewayProxyResponse, error) {
	return &events.APIGatewayProxyResponse{
		StatusCode:      500,
		Headers:         map[string]string{"Content-Type": "text/plain"},
		Body:            err.Error(),
		IsBase64Encoded: false,
	}, nil
}

type Echo struct {
	Protocol string
	Hostname string
	Method   string
	Path     string
	Query    map[string]string
	Headers  map[string]string
	Body     string
}

func handler(request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var err error

	e, err := json.Marshal(Echo{
		Protocol: request.RequestContext.Protocol,
		Hostname: request.RequestContext.DomainName,
		Method:   request.HTTPMethod,
		Path:     request.Path,
		Query:    request.QueryStringParameters,
		Headers:  request.Headers,
		Body:     request.Body,
	})
	if err != nil {
		return generateErrorResponse(err)
	}
	body := string(e)

	return &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Headers": "Content-Type",
			"Access-Control-Allow-Methods": "GET, POST, PUT, DELETE",
			"Content-Type":                 "text/plain",
			"Cache-Control":                "no-cache",
		},
		Body:            body,
		IsBase64Encoded: false,
	}, nil
}

func main() {
	// Make the handler available for Remote Procedure Call
	lambda.Start(handler)
}
