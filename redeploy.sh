#!/bin/bash
echo "Build the binary"
GOOS=linux GOARCH=amd64 go build -o main main.go
chmod +x main
echo "Create a ZIP file"
zip deployment.zip main

echo "Cleaning up"
rm main
echo "Updating Lambda"
aws lambda update-function-code --function-name AllPathsInOne  --zip-file fileb://deployment.zip --region us-east-1

echo "Running test"
sleep 1.5
go test tests/aws_test.go
echo "Finished"