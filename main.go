package main

import (
	"context"
	"encoding/json"
	"log"

	"fmt"
	"net/http"

	"strconv"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/zhekagigs/go-serverless-example/table"
	"github.com/aws/aws-sdk-go-v2/config"
)

var moviesDB *table.TableBasics

var (
	ERRServerResponse = events.APIGatewayProxyResponse{
		StatusCode: http.StatusInternalServerError,

		Body: "Error processing on a server",
	}
	ERRBadMethod = events.APIGatewayProxyResponse{
		StatusCode: http.StatusMethodNotAllowed,
		Body:       "method not supported or path is wrong",
	}
	ERRBadRequest = events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body:       "ID must be a number",
	}
)



func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatalf("failed to load configuration, %v", err)
	}
	moviesDB = table.NewTable(cfg, "movies")
	lambda.Start(PathCaller)
}

func ERRcheck(err error) (events.APIGatewayProxyResponse, error) {
	return ERRServerResponse, err
}

func NewOKResponse(body []byte) events.APIGatewayProxyResponse {
	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Headers: map[string]string{
			"Content-Type": "application/json",
		},
		Body: string(body),
	}
}

func FindAll(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	size, err := strconv.Atoi(req.QueryStringParameters["limit"])
	if err != nil {
		return events.APIGatewayProxyResponse{
		StatusCode: http.StatusBadRequest,
		Body: "limit should be a number",
		}, nil
	}
	movies, err  := moviesDB.Scan(size)
	if err != nil {
		return ERRcheck(err)
	}
	response, err := json.Marshal(movies)
	if err != nil {
		return ERRcheck(err)
	}
	return NewOKResponse(response), nil
}

func PathCaller(req events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("query", req.QueryStringParameters)
	fmt.Println("path", req.RequestContext.HTTP.Method)
	fmt.Println("**************************************************************************************")
	fmt.Printf("%+v\n", req)
	fmt.Println("**************************************************************************************")
	switch req.RequestContext.HTTP.Method {
	case "GET":
		return FindAll(req)
	default:
		return ERRServerResponse, nil
	}
}