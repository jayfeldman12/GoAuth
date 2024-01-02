package lambdas

import (
	"GoAuth/src/common"
	"GoAuth/src/mocks"
	"GoAuth/src/types"
	"encoding/json"
	"testing"

	"github.com/aws/aws-lambda-go/events"
)

var validUsers = []types.User{
	{
		Username:    common.ToPointer("test"),
		Password:    common.ToPointer("test123!"),
		Email:       common.ToPointer("email@email.email"),
		FirstName:   common.ToPointer("Bobby"),
		LastName:    common.ToPointer("Tables"),
		PhoneNumber: common.ToPointer("2222222222"),
	},
	{
		Username:    common.ToPointer("test"),
		Password:    common.ToPointer("test123!"),
		Email:       common.ToPointer("email@email.email"),
		FirstName:   common.ToPointer("Bobby"),
		LastName:    common.ToPointer("Tables"),
		PhoneNumber: common.ToPointer("2222222222"),
		Pet:         common.ToPointer("dog"),
	},
}

var invalidUsers = []types.User{
	{
		Password:    common.ToPointer("test123!"),
		Email:       common.ToPointer("email@email.email"),
		FirstName:   common.ToPointer("Bobby"),
		LastName:    common.ToPointer("Tables"),
		PhoneNumber: common.ToPointer("2222222222"),
	},
	{
		Username:    common.ToPointer("test"),
		Email:       common.ToPointer("email@email.email"),
		FirstName:   common.ToPointer("Bobby"),
		LastName:    common.ToPointer("Tables"),
		PhoneNumber: common.ToPointer("2222222222"),
		Pet:         common.ToPointer("dog"),
	},
	{
		Username:    common.ToPointer("test"),
		Password:    common.ToPointer("test123!"),
		FirstName:   common.ToPointer("Bobby"),
		LastName:    common.ToPointer("Tables"),
		PhoneNumber: common.ToPointer("2222222222"),
	},
	{
		Username:    common.ToPointer("test"),
		Password:    common.ToPointer("test123!"),
		Email:       common.ToPointer("email@email.email"),
		LastName:    common.ToPointer("Tables"),
		PhoneNumber: common.ToPointer("2222222222"),
	},
	{
		Username:    common.ToPointer("test"),
		Password:    common.ToPointer("test123!"),
		Email:       common.ToPointer("email@email.email"),
		FirstName:   common.ToPointer("Bobby"),
		PhoneNumber: common.ToPointer("2222222222"),
	},
	{
		Username:  common.ToPointer("test"),
		Password:  common.ToPointer("test123!"),
		Email:     common.ToPointer("email@email.email"),
		FirstName: common.ToPointer("Bobby"),
		LastName:  common.ToPointer("Tables"),
	},
}

func marshalBody(input types.User, t *testing.T) string {
	body, err := json.Marshal(input)
	if err != nil {
		t.Fatalf("Error converting user to string %q", err)
	}
	return string(body)
}

func TestCreateUser(t *testing.T) {
	for _, user := range validUsers {
		body := marshalBody(user, t)
		event := events.APIGatewayV2HTTPRequest{Body: string(body)}
		response, err := createUser(event, mocks.MockContext)
		if err != nil {
			t.Fatalf("Error creating valid user, %+v %s", user, err)
		}

		if response.StatusCode != 200 {
			t.Fatalf("Non-200 status code in response %d %+v", response.StatusCode, response.Body)
		}
	}

	for _, user := range invalidUsers {
		body := marshalBody(user, t)
		event := events.APIGatewayV2HTTPRequest{Body: string(body)}
		response, _ := createUser(event, mocks.MockContext)

		if response.StatusCode == 200 {
			t.Fatalf("200 status code for invalid user %d %+v", response.StatusCode, response.Body)
		}
	}
}
