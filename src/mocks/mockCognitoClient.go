package mocks

import (
	"GoAuth/src/common"
	"GoAuth/src/types"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
)

var MockCognitoClient = types.CognitoClient{
	SignUp: func(input *cognitoidentityprovider.SignUpInput) (*cognitoidentityprovider.SignUpOutput, error) {
		return &cognitoidentityprovider.SignUpOutput{
			UserSub: common.ToPointer("testSub"),
		}, nil
	},
	ConfirmSignUp: func(input *cognitoidentityprovider.ConfirmSignUpInput) (*cognitoidentityprovider.ConfirmSignUpOutput, error) {
		return &cognitoidentityprovider.ConfirmSignUpOutput{}, nil
	},
	AdminInitiateAuth: func(input *cognitoidentityprovider.AdminInitiateAuthInput) (*cognitoidentityprovider.AdminInitiateAuthOutput, error) {
		return &cognitoidentityprovider.AdminInitiateAuthOutput{
			AuthenticationResult: &cognitoidentityprovider.AuthenticationResultType{
				AccessToken:  common.ToPointer("testAccessToken"),
				RefreshToken: common.ToPointer("testRefreshToken"),
			},
		}, nil
	},
}
