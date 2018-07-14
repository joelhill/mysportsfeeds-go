package msf

type GamesIO struct {
	Games Games `json:"games"`
}

type Games struct {
	LastUpdatedOn string      `json:"lastUpdatedOn"`
	Game          []Game      `json:"game"`
	References    *References `json:"references"`
}

type Game struct {
	Schedule *Schedule `json:"schedule"`
	Score    *Score    `json:"score"`
}

type Schedule struct {
	ID                       int      `json:"id"`
	StartTime                string   `json:"startTime"`
	AwayTeam                 AwayTeam `json:"awayTeam"`
	HomeTeam                 HomeTeam `json:"homeTeam"`
	Venue                    Venue    `json:"venue"`
	VenueAllegiance          string   `json:"venueAllegiance"`
	ScheduleStatus           string   `json:"scheduleStatus"`
	OriginalStartTime        *string  `json:"originalStartTime"`
	DelayedOrPostponedReason *string  `json:"delayedOrPostponedReason"`
	PlayedStatus             string   `json:"playedStatus"`
}

type AwayTeam struct {
	ID           int `json:"id"`
	Abbreviation int `json:"abbreviation"`
}

type HomeTeam struct {
	ID           int `json:"id"`
	Abbreviation int `json:"abbreviation"`
}

type Venue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Score struct {
	CurrentInning       *string   `json:"currentInning"`
	CurrentInningHalf   *string   `json:"currentInningHalf"`
	CurrentIntermission *string   `json:"currentIntermission"`
	PlayStatus          *string   `json:"playStatus"`
	AwayScoreTotal      *int      `json:"awayScoreTotal"`
	AwayHitsTotal       *int      `json:"awayHitsTotal"`
	AwayErrorsTotal     *int      `json:"awayErrorsTotal"`
	HomeScoreTotal      *int      `json:"homeScoreTotal"`
	HomeHitsTotal       *int      `json:"homeHitsTotal"`
	HomeErrorsTotal     *int      `json:"homeErrorsTotal"`
	Innings             []*Inning `json:"innings"`
}

type Inning struct {
	InningNumber *int `json:"inningNumber"`
	AwayScore    *int `json:"awayScore"`
	HomeScore    *int `json:"homeScore"`
}

type References struct {
	TeamReferences  []TeamReference  `json:"teamReferences"`
	VenueReferences []VenueReference `json:"venueReferences"`
}

type TeamReference struct {
	ID           int       `json:"id"`
	City         string    `json:"city"`
	Name         string    `json:"name"`
	Abbreviation string    `json:"abbreviation"`
	HomeVenue    HomeVenue `json:"homeVenue"`
}

type HomeVenue struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type VenueReference struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	Country string `json:"country"`
}
