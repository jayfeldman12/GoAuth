package lambdas

import (
	"github.com/aws/aws-lambda-go/events"
)

func GetSecuredResource(request events.APIGatewayV2HTTPRequest) (events.APIGatewayProxyResponse, error) {
	response := events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       "Super secret string",
	}
	return response, nil
}
