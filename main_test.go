package main

import (
	"reflect"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

func Test_handler(t *testing.T) {
	tests := []struct {
		name    string
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "empty",
			want: events.APIGatewayProxyResponse{
				StatusCode: 200,
				Body:       "Welcome to Serverless World, Evgheny",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := handler()
			if (err != nil) != tt.wantErr {
				t.Errorf("handler() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("handler() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findAll(t *testing.T) {
	tests := []struct {
		name    string
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("findAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findAll() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insert(t *testing.T) {
	type args struct {
		req events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := insert(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("insert() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_findOne(t *testing.T) {
	type args struct {
		req events.APIGatewayProxyRequest
	}
	tests := []struct {
		name    string
		args    args
		want    events.APIGatewayProxyResponse
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findOne(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("findOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findOne() = %v, want %v", got, tt.want)
			}
		})
	}
}
