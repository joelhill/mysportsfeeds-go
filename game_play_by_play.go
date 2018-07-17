package msf

// GamePlayByPlayOptions - Are the options to hit the seasonal DFS endpoint
type GamePlayByPlayOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Game    string
	Format  string

	// Optional URL Params
	Playtype string
	Sort     string
	Offset   string
	Limit    string
	Force    string
}

// TODO: add this api call.
