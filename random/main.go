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
	SELECT channel_id, file_id, file_name
	FROM discord_images
	WHERE guild_id = ?
	ORDER BY rand() limit 1;
`

type imageRow struct {
	ChannelID string
	FileID    string
	FileName  string
}

func imgURL(row imageRow) string {
	return fmt.Sprintf(
		"https://cdn.discordapp.com/attachments/%s/%s/%s",
		row.ChannelID, row.FileID, row.FileName,
	)
}

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

	var image imageRow
	err = db.QueryRow(query, request.PathParameters["guildId"]).Scan(
		&image.ChannelID,
		&image.FileID,
		&image.FileName,
	)
	if err != nil {
		fmt.Println(err)
		return events.APIGatewayProxyResponse{Body: "Couldn't find any image", StatusCode: 404}, nil
	}

	headers := map[string]string{
		"Location": imgURL(image),
	}
	return events.APIGatewayProxyResponse{Headers: headers, StatusCode: 302}, nil
}

func main() {
	lambda.Start(handleRequest)
}
