package msf

// DailyPlayerGameslogsOptions - Are the options to hit the daily player gamelogs endpoint
type DailyPlayerGameslogsOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Date    string // YYYYMMDD
	Week    string // 1-40
	Format  string

	// Optional URL Params
	Team     string
	Player   string
	Position string
	Game     string
	Stats    string
	Sort     string
	Offset   string
	Limit    string
	Force    string
}

//TODO: fill this in
