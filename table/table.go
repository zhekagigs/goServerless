package table
import (
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	// "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)


var TABLE_NAME = "movies"

type TableBasics struct {
	DynamoDbClient *dynamodb.Client
	TableName      string
}

type Movie struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func NewTable(cfg aws.Config, tablename string) (*TableBasics){
	client := dynamodb.NewFromConfig(cfg)
	return &TableBasics{
		client,
		tablename,
	}
}

func (basics TableBasics) Scan(limit int) ([]Movie, error) {
	var movies []Movie
	var err error
	var response *dynamodb.ScanOutput
	
	response, err = basics.DynamoDbClient.Scan(context.TODO(), &dynamodb.ScanInput{
		TableName:                 aws.String(basics.TableName),
		Limit: aws.Int32(int32(limit)),
	})
	if err != nil {
		return nil, err
	} else {
		err = attributevalue.UnmarshalListOfMaps(response.Items, &movies)
		if err != nil {
			log.Printf("Couldn't unmarshal query response. Here's why: %v\n", err)
		}
	}
	
	return movies, err
}