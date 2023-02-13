package main

// import (
// 	"os"
// 	"reflect"
// 	"testing"

// 	"github.com/aws/aws-lambda-go/events"
// 	"github.com/aws/aws-sdk-go/aws/session"
// 	"github.com/aws/aws-sdk-go/service/dynamodb"
// )

// var test_svc *dynamodb.DynamoDB

// func TestMain(m *testing.M) {
// 	sess := session.Must(session.NewSessionWithOptions(session.Options{
// 		SharedConfigState: session.SharedConfigEnable,
// 	}))

// 	// Create DynamoDB client
// 	test_svc = dynamodb.New(sess)
// 	svc = test_svc
// 	os.Exit(m.Run())
// }

// func Test_insert(t *testing.T) {
// 	type args struct {
// 		req events.APIGatewayProxyRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    events.APIGatewayProxyResponse
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := Insert(tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("insert() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("insert() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_findOne(t *testing.T) {
// 	type args struct {
// 		req events.APIGatewayProxyRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    events.APIGatewayProxyResponse
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := FindOne(tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("findOne() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("findOne() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func Test_splitPath(t *testing.T) {
// 	type args struct {
// 		req *events.APIGatewayProxyRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    *URLPath
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "test GET /movies",
// 			args: args{
// 				req: &events.APIGatewayProxyRequest{
// 					HTTPMethod: "GET",
// 					Path:       "/movies",
// 				},
// 			},
// 			want: &URLPath{
// 				root:   "/movies",
// 				id:     "",
// 				method: "GET",
// 			},
// 			wantErr: false,
// 		}, {
// 			name: "test GET /movies/1",
// 			args: args{
// 				req: &events.APIGatewayProxyRequest{
// 					HTTPMethod: "GET",
// 					Path:       "/movies/1",
// 				},
// 			},
// 			want: &URLPath{
// 				root:   "/movies",
// 				id:     "/1",
// 				method: "GET",
// 			},
// 			wantErr: false,
// 		}, {
// 			name: "test GET /movies/1255165",
// 			args: args{
// 				req: &events.APIGatewayProxyRequest{
// 					HTTPMethod: "GET",
// 					Path:       "/movies/1255165",
// 				},
// 			},
// 			want: &URLPath{
// 				root:   "/movies",
// 				id:     "/1255165",
// 				method: "GET",
// 			},
// 			wantErr: false,
// 		}, {
// 			name: "test GET /movies/",
// 			args: args{
// 				req: &events.APIGatewayProxyRequest{
// 					HTTPMethod: "GET",
// 					Path:       "/movies/",
// 				},
// 			},
// 			want: &URLPath{
// 				root:   "/movies",
// 				id:     "",
// 				method: "GET",
// 			},
// 			wantErr: false,
// 		}, {
// 			name: "test GET movies/",
// 			args: args{
// 				req: &events.APIGatewayProxyRequest{
// 					HTTPMethod: "GET",
// 					Path:       "movies/",
// 				},
// 			},
// 			want: &URLPath{
// 				root:   "movies",
// 				id:     "",
// 				method: "GET",
// 			},
// 			wantErr: false,
// 		}, {
// 			name: "test GET /movies/",
// 			args: args{
// 				req: &events.APIGatewayProxyRequest{
// 					HTTPMethod: "GET",
// 					Path:       "/movies/",
// 				},
// 			},
// 			want: &URLPath{
// 				root:   "/movies",
// 				id:     "",
// 				method: "GET",
// 			},
// 			wantErr: false,
// 		}, {
// 			name: "test GET empty",
// 			args: args{
// 				req: &events.APIGatewayProxyRequest{
// 					HTTPMethod: "GET",
// 					Path:       "",
// 				},
// 			},
// 			want:    nil,
// 			wantErr: true,
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := ParsePath(tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("splitPath() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("splitPath() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestPathCaller(t *testing.T) {

// 	type args struct {
// 		req events.APIGatewayProxyRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    events.APIGatewayProxyResponse
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "test errors GET /movies",
// 			args: args{
// 				req: events.APIGatewayProxyRequest{
// 					HTTPMethod:            "GET",
// 					QueryStringParameters: map[string]string{"limit": "2", "pages": "2"},
// 				},
// 			},
// 			want: events.APIGatewayProxyResponse{
// 				StatusCode: 200,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			out, err := PathCaller(tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("PathCaller() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if out.StatusCode != 200 {
// 				t.Errorf("Error %v", out.StatusCode)
// 			}
// 		})
// 	}
// }

// func TestFindAll(t *testing.T) {
// 	type args struct {
// 		req events.APIGatewayProxyRequest
// 	}
// 	tests := []struct {
// 		name    string
// 		args    args
// 		want    events.APIGatewayProxyResponse
// 		wantErr bool
// 	}{
// 		// TODO: Add test cases.
// 		{
// 			name: "limit 2 pages 2",
// 			args: args{
// 				req: events.APIGatewayProxyRequest{
// 					HTTPMethod:            "GET",
// 					QueryStringParameters: map[string]string{"limit": "2", "pages": "2"},
// 				},
// 			},
// 			want: events.APIGatewayProxyResponse{
// 				StatusCode: 200,
// 			},
// 			wantErr: false,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			got, err := FindAll(tt.args.req)
// 			if (err != nil) != tt.wantErr {
// 				t.Errorf("FindAll() error = %v, wantErr %v", err, tt.wantErr)
// 				return
// 			}
// 			if len(got.Body) < 2 {
// 				t.Error("empty body")
// 				return
// 			}
// 		})
// 	}
// }
