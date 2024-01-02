package types

import (
	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

type CognitoClient struct {
	SignUp            func(input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error)
	ConfirmSignUp     func(input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error)
	AdminInitiateAuth func(input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error)
}

type Context struct {
	CognitoClient       *CognitoClient
	UserPoolID          string
	CognitoClientID     string
	CognitoClientSecret string
	Token               string
}
