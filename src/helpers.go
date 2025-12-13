package helpers

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// GetDynamoDBSession returns an AWS DynamoDB session object
func GetDynamoDBSession() (*dynamodb.DynamoDB, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(os.Getenv("DYNAMODB_REGION"))}, nil)
	if err!= nil {
		return nil, err
	}

	return dynamodb.New(sess), nil
}

// ScanDynamoDBTable scans a table in DynamoDB and returns all items
func ScanDynamoDBTable(ctx context.Context, dynamoDB *dynamodb.DynamoDB, tableName string) ([]map[string]*dynamodb.AttributeValue, error) {
	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	result, err := dynamoDB.Scan(params)
	if err!= nil {
		return nil, err
	}

	var items []map[string]*dynamodb.AttributeValue
	for _, item := range result.Items {
		itemMap, err := dynamodbattribute.MarshalMap(item)
		if err!= nil {
			return nil, err
		}
		items = append(items, itemMap)
	}

	return items, nil
}

// PutDynamoDBItem puts an item into DynamoDB
func PutDynamoDBItem(ctx context.Context, dynamoDB *dynamodb.DynamoDB, tableName string, item map[string]*dynamodb.AttributeValue) error {
	params := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}

	_, err := dynamoDB.PutItem(params)
	return err
}

// GetEnvString returns the value of an environment variable
func GetEnvString(key string) (string, error) {
	return os.Getenv(key)
}

// GetEnvInt returns the value of an environment variable as an integer
func GetEnvInt(key string) (int, error) {
	value, err := GetEnvString(key)
	if err!= nil {
		return 0, err
	}
	intValue, err := strconv.Atoi(value)
	if err!= nil {
		return 0, err
	}
	return intValue, nil
}

// GetEnvBool returns the value of an environment variable as a boolean
func GetEnvBool(key string) (bool, error) {
	value, err := GetEnvString(key)
	if err!= nil {
		return false, err
	}
	if value == "true" || value == "1" || value == "yes" || value == "Y" || value == "True" || value == "1" {
		return true, nil
	}
	if value == "false" || value == "0" || value == "no" || value == "N" || value == "False" || value == "0" {
		return false, nil
	}
	return false, fmt.Errorf("invalid boolean value for %s: %s", key, value)
}

// LogError logs an error with the given message
func LogError(msg string) {
	log.Printf("ERROR: %s\n", msg)
}