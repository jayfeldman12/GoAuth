package lambdas

import (
	"context"
	"log"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/dgrijalva/jwt-go"
)

func Authorizer(ctx context.Context, event events.APIGatewayV2CustomAuthorizerV2Request) (events.APIGatewayV2CustomAuthorizerIAMPolicyResponse, error) {
	var rawToken string
	for _, identity := range event.IdentitySource {
		if strings.HasPrefix(identity, "Bearer ") {
			rawToken = strings.TrimPrefix(identity, "Bearer ")
			break
		}
	}
	token := strings.TrimPrefix(rawToken, "Bearer ")

	if token == "" {
		log.Println("No token found in request")
		return generatePolicy("user", "Deny", nil, event.RouteArn), nil
	}

	// Validate the token
	// This is where you would decode and check the token
	jwtToken, err := ValidateToken(token)

	if err == nil {
		// If valid, return an allow policy
		return generatePolicy("user", "Allow", jwtToken, event.RouteArn), nil
	} else {
		// If invalid, return a deny policy
		log.Println(err)
		return generatePolicy("user", "Deny", jwtToken, event.RouteArn), nil
	}
}

func generatePolicy(principalID string, effect string, token *jwt.Token, resource string) events.APIGatewayV2CustomAuthorizerIAMPolicyResponse {
	authResponse := events.APIGatewayV2CustomAuthorizerIAMPolicyResponse{PrincipalID: principalID}
	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
		if token != nil {
			authResponse.Context = map[string]interface{}{"Sub": token.Claims.(jwt.MapClaims)["sub"]}
		}
	}
	return authResponse
}
