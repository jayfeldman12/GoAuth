package lambdas

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func RefreshToken(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var requestBody RefreshTokenRequest
	err := json.Unmarshal([]byte(request.Body), &requestBody)
	if err != nil {
		errorMessage := fmt.Sprintf("Error unmarshalling request %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	if requestBody.RefreshToken == nil {
		errorMessage := "Error missing required fields"
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	authRequest := &cognitoidentityprovider.AdminInitiateAuthInput{
		AuthFlow:   aws.String("REFRESH_TOKEN_AUTH"),
		ClientId:   aws.String(AppContext.CognitoClientID),
		UserPoolId: aws.String(AppContext.UserPoolID),
		AuthParameters: map[string]*string{
			"REFRESH_TOKEN": aws.String(*requestBody.RefreshToken),
			"SECRET_HASH":   aws.String(ComputeSecretHash(*requestBody.Username)),
		},
	}

	authResp, err := AppContext.CognitoClient.AdminInitiateAuth(authRequest)
	if err != nil {
		errorMessage := fmt.Sprintf("Error logging in %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	accessToken := *authResp.AuthenticationResult.AccessToken
	idToken := *authResp.AuthenticationResult.IdToken

	responseBody, err := json.Marshal(RefreshTokenResponse{AccessToken: accessToken, IdToken: idToken})
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

type RefreshTokenResponse struct {
	AccessToken string `json:"accessToken"`
	IdToken     string `json:"idToken"`
}

type RefreshTokenRequest struct {
	RefreshToken *string `json:"refreshToken" required:"true"`
	Username     *string `json:"username" required:"true"`
}
