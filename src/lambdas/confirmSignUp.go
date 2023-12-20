package lambdas

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func ConfirmSignUp(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody ConfirmSignUpRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		errorMessage := fmt.Sprintf("Error unmarshalling request %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	if requestBody.Username == nil || requestBody.Code == nil {
		errorMessage := "Error missing required fields"
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	confirmSignUpRequest := cognitoidentityprovider.ConfirmSignUpInput{
		ClientId: &AppContext.CognitoClientID,
		SecretHash: aws.String(
			ComputeSecretHash(
				AppContext.CognitoClientID,
				AppContext.CognitoClientSecret,
				*requestBody.Username,
			),
		),
		Username:         requestBody.Username,
		ConfirmationCode: requestBody.Code,
	}

	result, err := AppContext.CognitoClient.ConfirmSignUp(&confirmSignUpRequest)
	if err != nil {
		errorMessage := fmt.Sprintf("Error signing up user %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}
	log.Printf("User signed up %s", result.GoString)

	if err != nil {
		errorMessage := fmt.Sprintf("Error creating response body, %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
	}
	return response, nil
}

type ConfirmSignUpRequest struct {
	Username *string `json:"username" required:"true"`
	Code     *string `json:"code" required:"true"`
}
