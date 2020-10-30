package main

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nupamore/lambda/services"
)

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	guildID := request.PathParameters["guildId"]
	target, err := services.GetRandomImage(guildID)

	if err != nil {
		return events.APIGatewayProxyResponse{Body: "Couldn't find target", StatusCode: 404}, nil
	}

	headers := map[string]string{
		"Location": *target,
	}
	return events.APIGatewayProxyResponse{Headers: headers, StatusCode: 302}, nil
}

func main() {
	lambda.Start(handleRequest)
}
