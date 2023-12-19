package lambdas

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"

	"GoAuth/src/types"
)

func CreateUser(request events.APIGatewayWebsocketProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("Received request info: %q", request.Body)

	var person types.Person
	err := json.Unmarshal([]byte(request.Body), &person)
	if err != nil {
		log.Fatalf("Error: %q", err)
	}

	log.Printf("Received person info: %q", person)
	response := events.APIGatewayProxyResponse{StatusCode: 200}
	return response, nil
}
