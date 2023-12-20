package lambdas

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

func ComputeSecretHash(username string) string {
	clientId := AppContext.CognitoClientID
	clientSecret := AppContext.CognitoClientSecret
	mac := hmac.New(sha256.New, []byte(clientSecret))
	mac.Write([]byte(username + clientId))

	return base64.StdEncoding.EncodeToString(mac.Sum(nil))
}
