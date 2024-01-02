package mocks

import "GoAuth/src/types"

var MockContext = types.Context{
	CognitoClient:       &MockCognitoClient,
	UserPoolID:          "testId",
	CognitoClientID:     "testClientId",
	CognitoClientSecret: "testSecret",
	Token:               "testToken",
}
