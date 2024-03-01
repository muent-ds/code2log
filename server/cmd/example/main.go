package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type item struct {
	Name     string    `dynamo:"Name,hash"`
	Link     string    `dynamo:"Link"`
	Datetime time.Time `dynamo:"Datetime"`
}

type Response events.APIGatewayProxyResponse

func Handler(ctx context.Context) (Response, error) {
	/**
	 * DynamoDB初期設定
	 */
	dynamoDbEndpoint := os.Getenv("DYNAMODB_ENDPOINT")
	dynamoDbTable := os.Getenv("DYNAMODB_TABLE")

	disableSsl := false
	if len(dynamoDbEndpoint) != 0 {
		disableSsl = true
	}

	dynamoDbRegion := os.Getenv("AWS_REGION")
	if len(dynamoDbRegion) == 0 {
		dynamoDbRegion = "ap-northeast-1"
	}

	sess := session.Must(session.NewSession())
	db := dynamo.New(sess, &aws.Config{Endpoint: aws.String(dynamoDbEndpoint), DisableSSL: aws.Bool(disableSsl), Region: aws.String(dynamoDbRegion)})
	table := db.Table(dynamoDbTable)

	/**
	 * 保存するデータを生成
	 */
	// JSTで現在時刻を取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
		return Response{StatusCode: 500}, err
	}
	now := time.Now().In(jst) // ex. 2023-02-01T12:30:00+09:00

	item := item{
		Name:     "テスト",
		Link:     "https://zenn.dev",
		Datetime: now,
	}

	/**
	 * DynamoDBにPUT
	 */
	if err := table.Put(item).Run(); err != nil {
		log.Fatal(err)
		return Response{StatusCode: 500}, err
	}

	/**
	 * 結果をレスポンス
	 */
	var buf bytes.Buffer
	body, err := json.Marshal(map[string]interface{}{
		"message": "function executed successfully!",
	})
	if err != nil {
		return Response{StatusCode: 404}, err
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
