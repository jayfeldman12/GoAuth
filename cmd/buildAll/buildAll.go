package main

import (
	"GoAuth/scripts"
)

func main() {
	resources := []string{
		"./cmd/createUser/createUserCmd.go",
		"./cmd/confirmSignUp/confirmSignUpCmd.go",
		"./cmd/login/loginCmd.go",
		"./cmd/refreshToken/refreshTokenCmd.go",
		"./cmd/securedResource/securedResourceCmd.go",
	}
	commandBase := "go run ./cmd/lambdaBuild/lambdaBuild.go -- "

	for _, resource := range resources {
		scripts.Execute(commandBase + resource)
	}
}
