package tables

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/aws"
)

func CreateForumTable(svc *dynamodb.DynamoDB) (*dynamodb.CreateTableOutput, error){
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Forum"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("ForumName"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("Subject"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("ForumName"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("Subject"),
				KeyType: aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	return svc.CreateTable(params)
}