package msf

// Config - holds your authorization
type Config struct {
	// Authorization - your basic authentication
	// Required to hit any mysportsfeeds api endpoints
	Authorization string

	// BaseURL - the base url for my sports feeds api
	// Defaults to "https://api.mysportsfeeds.com"
	BaseURL string

	// Version - what version of the mysportsfeeds api you want to hit.
	// Defaults to "2.0"
	Version string

	// Sport - what sport you want data back from
	// Defaults to "mlb"
	Sport string

	// Format - what format you want to receive back from the api
	// Defaults to "json"
	Format string

	// Season - what season you want to access from the api
	// Defaults to "current"
	Season string
}

// NewConfig -
func NewConfig(authorization string) *Config {
	return &Config{
		Authorization: authorization,
		BaseURL:       MySportsFeedBaseURL,
		Version:       VersionV2_0,
		Sport:         SportMLB,
		Format:        FormatJSON,
		Season:        SeasonCurrent,
	}
}
