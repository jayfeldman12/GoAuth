package lambdas

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

func GetSecuredResource(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	jBytes, err := json.Marshal(request)
	if err != nil {
		return events.APIGatewayProxyResponse{StatusCode: 400, Body: "Error marshalling request"}, nil
	}
	log.Println(string(jBytes))

	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Super secret string",
	}
	return response, nil
}
