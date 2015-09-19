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
				AttributeName: aws.String("Name"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Name"),
				KeyType: aws.String("HASH"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	return svc.CreateTable(params)
}