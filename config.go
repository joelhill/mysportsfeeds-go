package msf

// Config - holds your authorization
type Config struct {
	// Authorization - your basic authentication
	Authorization string
}

// newConfig -
func newConfig() *Config {
	return &Config{}
}
