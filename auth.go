package bca

//AuthToken represents response of BCA OAuth 2.0 response message
type AuthToken struct {
	Error
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}
