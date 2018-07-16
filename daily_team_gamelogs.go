package msf

// DailyTeamGameslogsOptions - Are the options to hit the daily team gamelogs endpoint
type DailyTeamGameslogsOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Date    string // YYYYMMDD
	Format  string

	// Optional URL Params
	Team   string
	Game   string
	Stats  string
	Sort   string
	Offset string
	Limit  string
	Force  string
}

//TODO: fill this in
