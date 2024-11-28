package user

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/DeepLeau/serverless-stack/pkg/validators"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-lambda-go/events"
)

var (
	ErrorFailedToFetchRecord    = "failed to fetch record"
	ErrorFailedToUnmarshalRecord = "failed to unmarshal record"
	ErrorInvalidUserData         = "invalid user data"
	ErrorInvalidEmail            = "invalid email"
	ErrorCouldNotMarshalItem     = "could not marshal item"
	ErrorCouldNotDeleteItem      = "could not delete item"
	ErrorCouldNotDynamoPutItem   = "could not dynamo put item"
	ErrorUserAlreadyExists       = "user.User already exists"
	ErrorUserDoesNotExist        = "user.User does not exist"
)

type User struct {
	Email     string `json:"email"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func FetchUser(ctx context.Context, email, tableName string, dynaClient *dynamodb.Client) (*User, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: email},
		},
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.GetItem(ctx, input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	if result.Item == nil {
		return nil, errors.New(ErrorUserDoesNotExist)
	}

	item := new(User)
	err = attributevalue.UnmarshalMap(result.Item, item)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return item, nil
}

func FetchUsers(ctx context.Context, tableName string, dynaClient *dynamodb.Client) (*[]User, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynaClient.Scan(ctx, input)
	if err != nil {
		return nil, errors.New(ErrorFailedToFetchRecord)
	}

	var items []User
	err = attributevalue.UnmarshalListOfMaps(result.Items, &items)
	if err != nil {
		return nil, errors.New(ErrorFailedToUnmarshalRecord)
	}
	return &items, nil
}

func CreateUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*User, error) {
	var u User

	if err := json.Unmarshal([]byte(req.Body), &u); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}
	if !validators.IsEmailValid(u.Email) {
		return nil, errors.New(ErrorInvalidEmail)
	}

	currentUser, _ := FetchUser(ctx, u.Email, tableName, dynaClient)
	if currentUser != nil && len(currentUser.Email) != 0 {
		return nil, errors.New(ErrorUserAlreadyExists)
	}

	av, err := attributevalue.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(ctx, input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &u, nil
}

func UpdateUser(ctx context.Context, req events.APIGatewayProxyRequest, tableName string, dynaClient *dynamodb.Client) (*User, error) {
	var u User

	if err := json.Unmarshal([]byte(req.Body), &u); err != nil {
		return nil, errors.New(ErrorInvalidUserData)
	}

	currentUser, _ := FetchUser(ctx, u.Email, tableName, dynaClient)
	if currentUser == nil || len(currentUser.Email) == 0 {
		return nil, errors.New(ErrorUserDoesNotExist)
	}

	av, err := attributevalue.MarshalMap(u)
	if err != nil {
		return nil, errors.New(ErrorCouldNotMarshalItem)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(tableName),
	}

	_, err = dynaClient.PutItem(ctx, input)
	if err != nil {
		return nil, errors.New(ErrorCouldNotDynamoPutItem)
	}
	return &u, nil
}

func DeleteUser(ctx context.Context, email, tableName string, dynaClient *dynamodb.Client) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"email": &types.AttributeValueMemberS{Value: email},
		},
		TableName: aws.String(tableName),
	}

	_, err := dynaClient.DeleteItem(ctx, input)
	if err != nil {
		return errors.New(ErrorCouldNotDeleteItem)
	}

	return nil
}
