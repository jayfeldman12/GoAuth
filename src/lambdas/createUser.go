package lambdas

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"GoAuth/src/types"
)

func CreateUser(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request info: %s", request.Body)

	var person types.Person
	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		errorMessage := fmt.Sprintf("Error unmarshalling request %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	if firstName := person.FirstName; firstName == nil {
		errorMessage := "Missing required field: firstName"
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	responseBody, err := json.Marshal(CreateUserResponse{Message: fmt.Sprintf("Hello %s", *person.FirstName)})
	if err != nil {
		errorMessage := fmt.Sprintf("Error creating response body, %s", err)
		log.Print(errorMessage)
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: errorMessage}, nil
	}

	log.Printf("Received person info: %+v", person)
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(responseBody),
	}
	return response, nil
}

type CreateUserResponse struct {
	Message string `json:"message"`
}
