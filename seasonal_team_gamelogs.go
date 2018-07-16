package msf

// SeaonalTeamGamelogsOptions - Are the options to hit the seasonal team gamelogs endpoint
type SeaonalTeamGamelogsOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Format  string

	// Optional URL Params
	Team   string
	Game   string
	Date   string
	Stats  string
	Sort   string
	Offset string
	Limit  string
	Force  string
}

//TODO: fill this in
