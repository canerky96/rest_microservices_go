package access_token

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAccessTokenConstants(t *testing.T) {
	assert.EqualValuesf(t, 24, expirationTime, "expiration time should be 24 hours")
}

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	assert.Falsef(t, at.IsExpired(), "Brand new access token should not be expired")
	assert.EqualValuesf(t, "", at.AccessToken, "New Access token should not have defined access token id")
	assert.EqualValuesf(t, 0, at.UserId, "New access token should not have an associated user id")
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := AccessToken{}

	assert.True(t, at.IsExpired(), "empty access token should be expired by default")

	at.Expires = time.Now().UTC().Add(3 * time.Hour).Unix()
	assert.Falsef(t, at.IsExpired(), "Access token created three hours from now should NOT be expired")
}
