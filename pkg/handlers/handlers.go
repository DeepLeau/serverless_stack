package handlers

import (
	"context"
	"net/http"

	"github.com/DeepLeau/serverless_stack/pkg/user"
	"github.com/aws/aws-lambda-go/events"
	"github.com/DeepLeau/serverless_stack/pkg/dynamodb/dynamoapi"
	"github.com/aws/aws-sdk-go-v2/aws"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamoapi.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(context.TODO(), email, tableName, dynaClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{
				ErrorMsg: aws.String(err.Error()),
			})
		}
		return apiResponse(http.StatusOK, result)
	}

	result, err := user.FetchUsers(context.TODO(), tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamoapi.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.CreateUser(context.TODO(), req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusCreated, result)
}

func UpdateUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamoapi.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	result, err := user.UpdateUser(context.TODO(), req, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}
	return apiResponse(http.StatusOK, result)
}

func DeleteUser(req events.APIGatewayProxyRequest, tableName string, dynaClient dynamoapi.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) == 0 {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String("email is required"),
		})
	}

	err := user.DeleteUser(context.TODO(), email, tableName, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{
			ErrorMsg: aws.String(err.Error()),
		})
	}

	return apiResponse(http.StatusOK, nil)
}

func UnhandledMethod() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
