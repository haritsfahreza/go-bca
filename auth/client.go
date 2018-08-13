package auth

import (
	"context"
	"encoding/base64"
	"net/http"
	"net/url"
	"strings"

	bca "github.com/haritsfahreza/go-bca"
)

//Client is used to invoke BCA OAuth 2.0 API
type Client struct {
	Client       bca.ClientImplementation
	ClientID     string
	ClientSecret string
}

//NewClient is used to initialize new auth.Client
func NewClient(config bca.Config) Client {
	return Client{
		Client:       bca.NewClient(config),
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
	}
}

//GetToken is used to get OAuth 2.0 token
func (c Client) GetToken(ctx context.Context) (bca.AuthToken, error) {
	path := "/api/oauth/token"

	header := http.Header{}
	header.Add("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(c.ClientID+":"+c.ClientSecret)))

	data := url.Values{}
	data.Add("grant_type", "client_credentials")

	var response bca.AuthToken
	if err := c.Client.CallRaw("POST", path, "application/x-www-form-urlencoded",
		header, strings.NewReader(data.Encode()), &response); err != nil {
		return response, err
	}
	return response, nil
}
