package main
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"github.com/sinistersig/aws-practice-dynamodb/tables"
)

func main(){

	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}

	svc := dynamodb.New(config)

	//create product table
	resp, err := tables.CreateProductTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(resp)

	//Create forum table
	resp, err = tables.CreateForumTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(resp)

	//Reply
	resp, err = tables.CreateReplyTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(resp)
}