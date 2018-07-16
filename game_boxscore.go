package msf

// GameBoxscoreOptions - Are the options to hit the game boxscore endpoint
type GameBoxscoreOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Game    string
	Format  string

	// Optional URL Params
	TeamStats   string
	PlayerStats string
	Sort        string
	Offset      string
	Limit       string
	Force       string
}

//TODO: fill this in
