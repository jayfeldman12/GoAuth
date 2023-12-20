package lambdas

import (
	"GoAuth/src/types"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func CreateUser(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var user types.User
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		errorMessage := fmt.Sprintf("Error unmarshalling request %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	if user.Username == nil || user.Password == nil || user.Email == nil || user.FirstName == nil || user.LastName == nil || user.PhoneNumber == nil {
		errorMessage := "Error missing required fields"
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	signUpRequest := cognitoidentityprovider.SignUpInput{
		ClientId: &AppContext.CognitoClientID,
		SecretHash: aws.String(
			ComputeSecretHash(
				AppContext.CognitoClientID,
				AppContext.CognitoClientSecret,
				*user.Username,
			),
		),
		Username: user.Username,
		Password: user.Password,
		UserAttributes: []*cognitoidentityprovider.AttributeType{
			{
				Name:  aws.String("email"),
				Value: user.Email,
			},
			{
				Name:  aws.String("given_name"),
				Value: user.FirstName,
			},
			{
				Name:  aws.String("family_name"),
				Value: user.LastName,
			},
			{
				Name:  aws.String("phone_number"),
				Value: user.PhoneNumber,
			},
		},
	}

	if user.Pet != nil {
		signUpRequest.UserAttributes = append(signUpRequest.UserAttributes, &cognitoidentityprovider.AttributeType{
			Name:  aws.String("custom:attr1"),
			Value: user.Pet,
		})
	}

	result, err := AppContext.CognitoClient.SignUp(&signUpRequest)
	if err != nil {
		errorMessage := fmt.Sprintf("Error signing up user %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	responseBody, err := json.Marshal(CreateUserResponse{Message: fmt.Sprintf("Hello %s", *user.Username), Id: *(result.UserSub)})
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

type CreateUserResponse struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}
