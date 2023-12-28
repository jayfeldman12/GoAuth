package main

import (
	"GoAuth/cmd/common"
)

func main() {
	resources := []string{
		"./cmd/createUser/createUserCmd.go",
		"./cmd/confirmSignUp/confirmSignUpCmd.go",
		"./cmd/login/loginCmd.go",
		"./cmd/refreshToken/refreshTokenCmd.go",
		"./cmd/securedResource/securedResourceCmd.go",
		"./cmd/authorizer/authorizerCmd.go",
	}
	commandBase := "go run ./cmd/lambdaBuild/lambdaBuild.go -- "

	for _, resource := range resources {
		common.Execute(commandBase + resource)
	}
}
