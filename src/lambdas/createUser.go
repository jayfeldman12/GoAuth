package lambdas

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"GoAuth/src/types"
)

func CreateUser(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	var user types.User
	err := json.Unmarshal([]byte(request.Body), &user)
	if err != nil {
		errorMessage := fmt.Sprintf("Error unmarshalling request %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	responseBody, err := json.Marshal(CreateUserResponse{Message: fmt.Sprintf("Hello %s", user.Username)})
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
}
