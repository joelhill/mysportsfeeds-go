package msf

// Config - holds your authorization
type Config struct {
	// Authorization - your basic authentication
	Authorization string
}

// NewConfig -
func NewConfig() *Config {
	return &Config{}
}
