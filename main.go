package main

import (
	"encoding/json"
	"net/http"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var movies = []struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}{
	{
		ID:   1,
		Name: "Avengers",
	},
	{
		ID:   2,
		Name: "Ant-Man",
	},
	{
		ID:   3,
		Name: "Thor",
	},
	{
		ID:   4,
		Name: "Hulk",
	}, {
		ID:   5,
		Name: "Doctor Strange",
	},
}

func findAll()(events.APIGatewayProxyResponse, error) {
	payload, err := json.Marshal(movies)
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}
	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-type": "application/json",
		},
		Body: string(payload), 
	}, nil
}

func handler() (events.APIGatewayProxyResponse, error) {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Welcome to Serverless World, Evgheny",
	}, nil
}

func main() {
	lambda.Start(findAll)
}
