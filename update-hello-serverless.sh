#!bin/bash
aws lambda update-function-code --function-name FindAllMovies --zip-file fileb://deployment.zip --region us-east-1