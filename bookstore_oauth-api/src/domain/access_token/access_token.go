package access_token

import (
	"bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime = 24
)

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (token *AccessToken) Validate() *errors.RestErr {
	token.AccessToken = strings.TrimSpace(token.AccessToken)
	if token.AccessToken == "" {
		return errors.NewBadRequestError("Invalid access token id")
	}
	if token.UserId <= 0 {
		return errors.NewBadRequestError("Invalid user id")
	}
	if token.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid client id")
	}
	if token.Expires <= 0 {
		return errors.NewBadRequestError("Invalid expiration time")
	}
	return nil
}

func GetNewAccessToken() AccessToken {
	return AccessToken{
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (token AccessToken) IsExpired() bool {
	return time.Unix(token.Expires, 0).Before(time.Now().UTC())
}

// Web fronted - Client-Id : 123
// Android APP - Client-Id : 234
