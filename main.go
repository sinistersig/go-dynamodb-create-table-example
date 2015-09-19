package main
import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"log"
	"github.com/sinistersig/aws-practice-dynamodb-create-table/tables"
)

const CREATING = "CREATING..."

func main(){

	config := &aws.Config{
		Region: aws.String("us-east-1"),
	}

	svc := dynamodb.New(config)

	//Product
	_, err := tables.CreateProductTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(CREATING)

	//Forum
	_, err = tables.CreateForumTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(CREATING)

	//Reply
	_, err = tables.CreateReplyTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(CREATING)

	//Thread
	_, err = tables.CreateThreadTable(svc)

	if err != nil {
		log.Println(err.Error())
	}

	log.Println(CREATING)
}