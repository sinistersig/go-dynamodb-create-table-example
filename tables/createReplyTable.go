package tables

import "github.com/aws/aws-sdk-go/service/dynamodb"
import "github.com/aws/aws-sdk-go/aws"

func CreateReplyTable(svc *dynamodb.DynamoDB) (*dynamodb.CreateTableOutput, error){
	params := &dynamodb.CreateTableInput{
		TableName: aws.String("Reply"),
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("ReplyDateTime"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("PostedBy"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType: aws.String("HASH"),
			},
			{
				AttributeName: aws.String("ReplyDateTime"),
				KeyType: aws.String("RANGE"),
			},
		},
		LocalSecondaryIndexes: []*dynamodb.LocalSecondaryIndex{
			{
				IndexName: aws.String("PostedBy-index"),
				KeySchema: []*dynamodb.KeySchemaElement{
					{
						AttributeName: aws.String("Id"),
						KeyType: aws.String("HASH"),
					},
					{
						AttributeName: aws.String("PostedBy"),
						KeyType: aws.String("RANGE"),
					},
				},
				Projection: &dynamodb.Projection{
					ProjectionType: aws.String("KEYS_ONLY"),
				},
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	return svc.CreateTable(params)
}