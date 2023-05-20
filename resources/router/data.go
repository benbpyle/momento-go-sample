package main

import (
	"context"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

// NewDynamoDBClient inits a DynamoDB session to be used throughout the services
func NewDynamoDBClient(isLocal bool) dynamodbiface.DynamoDBAPI {
	c := &aws.Config{
		Region: aws.String("us-west-2")}

	sess := session.Must(session.NewSession(c))
	svc := dynamodb.New(sess)
	return dynamodbiface.DynamoDBAPI(svc)
}

// RouterRepository maps the interface for the operations when working with Route data
type RouterRepository interface {
	GetRoute(context.Context, string) (*Route, error)
}

// RouterDynamoRepository implements the RouterRepository interface and operates on top
// of DynamoDB
type RouterDynamoRepository struct {
	db dynamodbiface.DynamoDBAPI
}

func (r *RouterDynamoRepository) GetRoute(ctx context.Context, key string) (*Route, error) {
	input := &dynamodb.QueryInput{
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":pk": {
				S: aws.String(key),
			},
		},
		TableName:              aws.String(tableName),
		KeyConditionExpression: aws.String("pk = :pk"),
	}

	result, err := r.db.QueryWithContext(ctx, input)

	if err != nil {
		return nil, err
	}

	if len(result.Items) != 1 {
		return nil, nil
	}

	route := &Route{}
	err = dynamodbattribute.UnmarshalMap(result.Items[0], route)
	return route, err
}
