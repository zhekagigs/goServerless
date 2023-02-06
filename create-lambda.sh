aws lambda create-function --function-name FindAllMovies \
     --zip-file fileb://./deployment.zip \
     --runtime go1.x --handler main \
     --role arn:aws:iam::776928313458:role/FindAllMoviesRole \
     --region us-east-1