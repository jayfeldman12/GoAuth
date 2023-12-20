package main

import (
	"GoAuth/src/lambdas"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(lambdas.GetSecuredResource)
}
