package msf

// Config - holds your authorization
type Config struct {
	// Authorization - your basic authentication
	Authorization string
}

// NewConfig -
func NewConfig(authorization string) *Config {
	return &Config{
		Authorization: authorization,
	}
}
