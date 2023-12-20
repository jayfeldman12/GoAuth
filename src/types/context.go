package types

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type Context struct {
	CognitoClient       *cognitoidentityprovider.CognitoIdentityProvider
	UserPoolID          string
	CognitoClientID     string
	CognitoClientSecret string
	Token               string
}
