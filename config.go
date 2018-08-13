package bca

//Config represents configuration that needed by BCA API
type Config struct {
	ClientID     string
	ClientSecret string
	APIKey       string
	APISecret    string
	URL          string
	CorporateID  string
	OriginHost   string
	LogLevel     int
}
