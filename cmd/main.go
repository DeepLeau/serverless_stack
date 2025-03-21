package main

import (
	"context"
	"fmt"
	"os"
	
	"github.com/DeepLeau/serverless_stack/pkg/handlers"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type DynamoDBAPI interface {
	GetItem(ctx context.Context, params *dynamodb.GetItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.GetItemOutput, error)
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
	DeleteItem(ctx context.Context, params *dynamodb.DeleteItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.DeleteItemOutput, error)
	Scan(ctx context.Context, params *dynamodb.ScanInput, optFns ...func(*dynamodb.Options)) (*dynamodb.ScanOutput, error)
}

var dynaClient DynamoDBAPI

func main() {
	region := os.Getenv("AWS_REGION")
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		fmt.Printf("Error loading AWS configuration: %v\n", err)
		return
	}
	fmt.Printf("Configuration loaded for the region: %s\n", cfg.Region)

	dynaClient = dynamodb.NewFromConfig(cfg)

	lambda.Start(handler)
}

const tableName = "LambdaInGoUser"

func handler(req events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	fmt.Printf("Received request: %v\n", req)
	switch req.HTTPMethod {
	case "GET":
		return handlers.GetUser(req, tableName, dynaClient)
	case "POST":
		return handlers.CreateUser(req, tableName, dynaClient)
	case "PUT":
		return handlers.UpdateUser(req, tableName, dynaClient)
	case "DELETE":
		return handlers.DeleteUser(req, tableName, dynaClient)
	default:
		fmt.Printf("Unhandled method: %s\n", req.HTTPMethod)
		return handlers.UnhandledMethod()
	}
}
