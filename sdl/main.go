package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/go-sql-driver/mysql"
)

const query = `
	SELECT target
	FROM simple_dynamic_link
	WHERE link_id = ?;
`

func handleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	dbStr := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"), os.Getenv("password"),
		os.Getenv("host"), os.Getenv("port"), os.Getenv("database"),
	)
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	var target string
	err = db.QueryRow(query, request.PathParameters["linkId"]).Scan(&target)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: "Couldn't find target", StatusCode: 404}, nil
	}

	headers := map[string]string{
		"Location": target,
	}
	return events.APIGatewayProxyResponse{Headers: headers, StatusCode: 302}, nil
}

func main() {
	lambda.Start(handleRequest)
}
