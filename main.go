package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	// "github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/session"
	// "github.com/aws/aws-sdk-go/service/dynamodb"
	// "github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)
type Movie struct {
	ID int `json:"id"`
	Name string `json:"name"`
  }

var movies = []Movie {
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
func insert(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	var movie Movie
	err := json.Unmarshal([]byte(req.Body), &movie)
	if err != nil {
	  return events.APIGatewayProxyResponse{
		StatusCode: 400,
		Body: "Invalid payload",
	  }, nil
	}
  
	movies = append(movies, movie)
  
	response, err := json.Marshal(movies)
	if err != nil {
	  return events.APIGatewayProxyResponse{
		StatusCode: 500,
		Body: err.Error(),
	  }, nil
	}
  
	return events.APIGatewayProxyResponse{
	  StatusCode: 200,
	  Headers: map[string]string{
		"Content-Type": "application/json",
	  },
	  Body: string(response),
	}, nil
  }
  

func findOne(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	id, err := strconv.Atoi(req.PathParameters["id"])
	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "ID must be a number",
		}, err
	}
	response, err := json.Marshal(movies[id-1])

	if err != nil {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
			Body:       "ID not found",
		}, err
	}

	return events.APIGatewayProxyResponse{
		StatusCode: http.StatusOK,
		Headers: map[string]string{
			"Content-type": "application/json",
		},
		Body: string(response),
	}, nil
}

func findAll() (events.APIGatewayProxyResponse, error) {
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
	lambda.Start(insert)
}
