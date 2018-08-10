package auth

import (
	"bytes"
	"context"
	"encoding/base64"
	"net/http"

	bca "github.com/haritsfahreza/go-bca"
)

//Client is used to invoke BCA OAuth 2.0 API
type Client struct {
	config bca.Config
	client bca.ClientImplementation
}

//GetToken is used to get OAuth 2.0 token
func (c Client) GetToken(ctx context.Context) (bca.AuthToken, error) {
	path := "/api/oauth/token"

	var header http.Header
	header.Add("Authorization", base64.StdEncoding.EncodeToString([]byte(c.config.ClientID+":"+c.config.ClientSecret)))

	var response bca.AuthToken
	if err := c.client.CallRaw("POST", path, "application/x-www-form-urlencoded",
		header, bytes.NewBufferString("grant_type=client_credentials"), response); err != nil {
		return response, err
	}
	return response, nil
}
