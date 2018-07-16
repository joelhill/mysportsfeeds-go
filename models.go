package msf

// Main Models
type GamesIO struct {
	LastUpdatedOn string      `json:"lastUpdatedOn"`
	Games         *[]Game     `json:"games"`
	References    *References `json:"references"`
}

type DfsIO struct {
	LastUpdatedOn string      `json:"lastUpdatedOn"`
	References    *References `json:"references"`
}

type GameLogIO struct {
	LastUpdatedOn string      `json:"lastUpdatedOn"`
	GameLog       *[]GameLog  `json:"gamelog"`
	References    *References `json:"references"`
}

// Sub Models
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
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
}

type HomeTeam struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
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
	TeamReferences       []TeamReference        `json:"teamReferences"`
	VenueReferences      []VenueReference       `json:"venueReferences"`
	GameReferences       *[]GameReference       `json:"gameReferences"`
	PlayerReferences     *[]PlayerReference     `json:"playerReferences"`
	PlayerStatReferences *[]PlayerStatReference `json:"playerStatReferences"`
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

type GameReference struct {
	ID                       int      `json:"id"`
	StartTime                string   `json:"startTime"`
	AwayTeam                 AwayTeam `json:"awayTeam"`
	HomeTeam                 HomeTeam `json:"homeTeam"`
	Venue                    Venue    `json:"venue"`
	VenueAllegiance          string   `json:"venueAllegiance"`
	ScheduleStatus           string   `json:"scheduleStatus"`
	OriginalStartTime        *string  `json:"originalStartTime"`
	DelayedOrPostponedReason *string  `json:"delayedOrPostponedReason"`
	PlayedStatus             *string  `json:"playedStatus"`
}

type PlayerReference struct {
	Player
	CurrentTeam         Team        `json:"currentTeam"`
	CurrentRosterStatus string      `json:"currentRosterStatus"`
	CurrentInjury       *string     `json:"currentInjury"`
	Height              string      `json:"height"`
	Weight              float64     `json:"weight"`
	BirthDate           string      `json:"birthDate"`
	Age                 int         `json:"age"`
	BirthCity           string      `json:"birthCity"`
	BirthCountry        string      `json:"birthCountry"`
	Rookie              bool        `json:"rookie"`
	College             string      `json:"college"`
	Twitter             *string     `json:"twitter"`
	Handedness          *Handedness `json:"handedness"`
}

type PlayerStatReference struct {
	Category     string `json:"category"`
	FullName     string `json:"fullName"`
	Description  string `json:"description"`
	Abbreviation string `json:"abbreviation"`
	Type         string `json:"type"`
}

type Handedness struct {
	Bats   string `json:"bats"`
	Throws string `json:"throws"`
}

type GameLog struct {
	Game   GameLogGame `json:"game"`
	Player Player      `json:"player"`
	Team   `json:"team"`
	Stats  `json:""`
}

type GameLogGame struct {
	ID                   int    `json:"id"`
	StartTime            string `json:"startTime"`
	AwayTeamAbbreviation string `json:"awayTeamAbbreviation"`
	HomeTeamAbbreviation string `json:"homeTeamAbbreviation"`
}

type Player struct {
	ID           int    `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Position     string `json:"position"`
	JerseyNumber *int   `json:"jerseyNumber"`
}

type Team struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
}

type Stats struct {
	Batting       *Batting       `json:"batting"`
	Pitching      *Pitching      `json:"pitching"`
	Fielding      *Fielding      `json:"fielding"`
	Miscellaneous *Miscellaneous `json:"miscellaneous"`
}

type Batting struct {
	AtBats                       *int     `json:"atBats"`
	Runs                         *int     `json:"runs"`
	Hits                         *int     `json:"hits"`
	SecondBaseHits               *int     `json:"secondBaseHits"`
	ThirdBaseHits                *int     `json:"thirdBaseHits"`
	Homeruns                     *int     `json:"homeruns"`
	EarnedRuns                   *int     `json:"earnedRuns"`
	UnearnedRuns                 *int     `json:"unearnedRuns"`
	RunsBattedIn                 *int     `json:"runsBattedIn"`
	BatterWalks                  *int     `json:"batterWalks"`
	BatterSwings                 *int     `json:"batterSwings"`
	BatterStrikes                *int     `json:"batterStrikes"`
	BatterStrikesFoul            *int     `json:"batterStrikesFoul"`
	BatterStrikesMiss            *int     `json:"batterStrikesMiss"`
	BatterStrikesLooking         *int     `json:"batterStrikesLooking"`
	BatterTagOuts                *int     `json:"batterTagOuts"`
	BatterForceOuts              *int     `json:"batterForceOuts"`
	BatterPutOuts                *int     `json:"batterPutOuts"`
	BatterGroundBalls            *int     `json:"batterGroundBalls"`
	BatterFlyBalls               *int     `json:"batterFlyBalls"`
	BatterLineDrives             *int     `json:"batterLineDrives"`
	Batter2SeamFastballs         *int     `json:"batter2SeamFastballs"`
	Batter4SeamFastballs         *int     `json:"batter4SeamFastballs"`
	BatterCurveballs             *int     `json:"batterCurveballs"`
	BatterChangeups              *int     `json:"batterChangeups"`
	BatterCutters                *int     `json:"batterCutters"`
	BatterSliders                *int     `json:"batterSliders"`
	BatterSinkers                *int     `json:"batterSinkers"`
	BatterSplitters              *int     `json:"batterSplitters"`
	BatterStrikeouts             *int     `json:"batterStrikeouts"`
	StolenBases                  *int     `json:"stolenBases"`
	CaughtBaseSteals             *int     `json:"caughtBaseSteals"`
	BatterStolenBasePct          *float64 `json:"batterStolenBasePct"`
	BattingAvg                   *float64 `json:"battingAvg"`
	BatterOnBasePct              *float64 `json:"batterOnBasePct"`
	BatterSluggingPct            *float64 `json:"batterSluggingPct"`
	BatterOnBasePlusSluggingPct  *float64 `json:"batterOnBasePlusSluggingPct"`
	BatterIntentionalWalks       *int     `json:"batterIntentionalWalks"`
	HitByPitch                   *int     `json:"hitByPitch"`
	BatterSacrificeBunts         *int     `json:"batterSacrificeBunts"`
	BatterSacrificeFlies         *int     `json:"batterSacrificeFlies"`
	TotalBases                   *int     `json:"totalBases"`
	ExtraBaseHits                *int     `json:"extraBaseHits"`
	BatterDoublePlays            *int     `json:"batterDoublePlays"`
	BatterTriplePlays            *int     `json:"batterTriplePlays"`
	BatterGroundOuts             *int     `json:"batterGroundOuts"`
	BatterFlyOuts                *int     `json:"batterFlyOuts"`
	BatterGroundOutToFlyOutRatio *float64 `json:"batterGroundOutToFlyOutRatio"`
	PitchesFaced                 *int     `json:"pitchesFaced"`
	PlateAppearances             *int     `json:"plateAppearances"`
	LeftOnBase                   *int     `json:"leftOnBase"`
}

type Pitching struct {
}

type Fielding struct {
	InningsPlayed             *float64 `json:"inningsPlayed"`
	TotalChances              *int     `json:"totalChances"`
	FielderTagOuts            *int     `json:"fielderTagOuts"`
	FielderForceOuts          *int     `json:"fielderForceOuts"`
	FielderPutOuts            *int     `json:"fielderPutOuts"`
	OutsFaced                 *int     `json:"outsFaced"`
	Assists                   *int     `json:"assists"`
	Errors                    *int     `json:"errors"`
	FielderDoublePlays        *int     `json:"fielderDoublePlays"`
	FielderTriplePlays        *int     `json:"fielderTriplePlays"`
	FielderStolenBasesAllowed *int     `json:"fielderStolenBasesAllowed"`
	FielderCaughtStealing     *int     `json:"fielderCaughtStealing"`
	FielderStolenBasePct      *float64 `json:"fielderStolenBasePct"`
	PassedBalls               *int     `json:"passedBalls"`
	FielderWildPitches        *int     `json:"fielderWildPitches"`
	FieldingPct               *float64 `json:"fieldingPct"`
	RangeFactor               *float64 `json:"rangeFactor"`
}

type Miscellaneous struct {
	GamesStarted *int `json:"gamesStarted"`
}
