package common

import (
	"GoAuth/src/lambdas"
	"GoAuth/src/types"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

func Init() {
	cognitoRegion := os.Getenv("COGNITO_REGION")
	conf := &aws.Config{Region: aws.String(cognitoRegion)}
	mySession := session.Must(session.NewSession(conf))

	client := cognitoidentityprovider.New(mySession)

	cognitoClient := types.CognitoClient{
		SignUp:            client.SignUp,
		ConfirmSignUp:     client.ConfirmSignUp,
		AdminInitiateAuth: client.AdminInitiateAuth,
	}

	lambdas.AppContext = types.Context{
		CognitoClient:       &cognitoClient,
		UserPoolID:          os.Getenv("COGNITO_USER_POOL_ID"),
		CognitoClientID:     os.Getenv("COGNITO_APP_CLIENT_ID"),
		CognitoClientSecret: os.Getenv("COGNITO_APP_CLIENT_SECRET"),
	}
}
