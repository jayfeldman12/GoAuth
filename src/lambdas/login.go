package lambdas

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func Login(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody LoginRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		errorMessage := fmt.Sprintf("Error unmarshalling request %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	if requestBody.Username == nil || requestBody.Password == nil {
		errorMessage := "Error missing required fields"
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	authRequest := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:   aws.String("ADMIN_USER_PASSWORD_AUTH"),
		ClientId:   aws.String(AppContext.CognitoClientID),
		UserPoolId: aws.String(AppContext.UserPoolID),
		AuthParameters: map[string]*string{
			"USERNAME":    requestBody.Username,
			"PASSWORD":    requestBody.Password,
			"SECRET_HASH": aws.String(ComputeSecretHash(*requestBody.Username)),
		},
	}

	authResp, err := AppContext.CognitoClient.AdminInitiateAuth(authRequest)
	if err != nil {
		errorMessage := fmt.Sprintf("Error logging in %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	accessToken := *authResp.AuthenticationResult.AccessToken
	refreshToken := *authResp.AuthenticationResult.RefreshToken
	idToken := *authResp.AuthenticationResult.IdToken

	responseBody, err := json.Marshal(LoginResponse{AccessToken: accessToken, RefreshToken: refreshToken, IDToken: idToken})
	if err != nil {
		errorMessage := fmt.Sprintf("Error creating response body, %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}
	return response, nil
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	IDToken      string `json:"idToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginRequest struct {
	Username *string `json:"username" required:"true"`
	Password *string `json:"password" required:"true"`
}
