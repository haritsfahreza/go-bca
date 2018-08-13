package bca

//Config represents configuration that needed by BCA API
type Config struct {
	ClientID     string
	ClientSecret string
	APIKey       string
	APISecret    string
	CorporateID  string
	OriginHost   string
}
