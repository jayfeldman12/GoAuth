package main

import (
	"GoAuth/cmd/common"
	"GoAuth/src/lambdas"

	"github.com/aws/aws-lambda-go/lambda"
)

func init() {
	common.Init()
}

func main() {
	lambda.Start(lambdas.ConfirmSignUp)
}
