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
	GameLog       *[]GameLog  `json:"gamelogs"`
	References    *References `json:"references"`
}

type GameLineupIO struct {
	LastUpdatedOn string        `json:"lastUpdatedOn"`
	Game          *Schedule     `json:"game"`
	TeamLineups   *[]TeamLineup `json:"teamLineups"`
	References    *References   `json:"references"`
}

type GamePlayByPlayIO struct {
	LastUpdatedOn string      `json:"lastUpdatedOn"`
	Game          *Schedule   `json:"game"`
	AtBats        *[]AtBat    `json:"atBats"`
	References    *References `json:"references"`
}

type BoxscoreIO struct {
	LastUpdatedOn string         `json:"lastUpdatedOn"`
	Game          *Schedule      `json:"game"`
	Scoring       *Score         `json:"scoring"`
	Stats         *StatsHomeAway `json:"stats"`
	References    *References    `json:"references"`
}

type CurrentSeasonIO struct {
	LastUpdatedOn string    `json:"lastUpdatedOn"`
	Seasons       *[]Season `json:"seasons"`
}

type PlayerInjuriesIO struct {
	LastUpdatedOn string             `json:"lastUpdatedOn"`
	Players       *[]PlayerReference `json:"players"`
}

type FeedUpdatesIO struct {
	LastUpdatedOn string        `json:"lastUpdatedOn"`
	FeedUpdates   *[]FeedUpdate `json:"feedUpdates"`
}

type TeamStatsTotalsIO struct {
	LastUpdatedOn   string            `json:"lastUpdatedOn"`
	TeamStatsTotals *[]TeamStatsTotal `json:"teamStatsTotals"`
}

type PlayerStatsTotalsIO struct {
	LastUpdatedOn     string              `json:"lastUpdatedOn"`
	PlayerStatsTotals *[]PlayerStatsTotal `json:"playerStatsTotals"`
	References        *References         `json:"references"`
}

type VenuesIO struct {
	LastUpdatedOn string    `json:"lastUpdatedOn"`
	Venues        *[]Venues `json:"venues"`
}

type PlayersIO struct {
	LastUpdatedOn string        `json:"lastUpdatedOn"`
	Players       *[]Players    `json:"players"`
	References    *[]References `json:"references"`
}

type StandingsIO struct {
	LastUpdatedOn string      `json:"lastUpdatedOn"`
	Teams         *[]Teams    `json:"teams"`
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
	CurrentInning       *int      `json:"currentInning"`
	CurrentInningHalf   *string   `json:"currentInningHalf"`
	CurrentIntermission *string   `json:"currentIntermission"`
	AwayScoreTotal      *int      `json:"awayScoreTotal"`
	AwayHitsTotal       *int      `json:"awayHitsTotal"`
	AwayErrorsTotal     *int      `json:"awayErrorsTotal"`
	HomeScoreTotal      *int      `json:"homeScoreTotal"`
	HomeHitsTotal       *int      `json:"homeHitsTotal"`
	HomeErrorsTotal     *int      `json:"homeErrorsTotal"`
	Innings             []*Inning `json:"innings"`
}

type Inning struct {
	InningNumber *int           `json:"inningNumber"`
	AwayScore    *int           `json:"awayScore"`
	HomeScore    *int           `json:"homeScore"`
	ScoringPlays *[]ScoringPlay `json:"scoringPlays"`
}

type ScoringPlay struct {
	InningHalf      *string `json:"inningHalf"`
	Team            *Team   `json:"team"`
	ScoreChange     *int    `json:"scoreChange"`
	AwayScore       *int    `json:"awayScore"`
	HomeScore       *int    `json:"homeScore"`
	PlayDescription *string `json:"playDescription"`
}

type References struct {
	TeamReferences       *[]TeamReference   `json:"teamReferences"`
	TeamStatReferences   *[]StatReference   `json:"teamStatReferences"`
	VenueReferences      *[]VenueReference  `json:"venueReferences"`
	GameReferences       *[]GameReference   `json:"gameReferences"`
	PlayerReferences     *[]PlayerReference `json:"playerReferences"`
	PlayerStatReferences *[]StatReference   `json:"playerStatReferences"`
}

type TeamReference struct {
	ID           *int       `json:"id"`
	City         *string    `json:"city"`
	Name         *string    `json:"name"`
	Abbreviation *string    `json:"abbreviation"`
	HomeVenue    *HomeVenue `json:"homeVenue"`
}

type HomeVenue struct {
	ID   *int    `json:"id"`
	Name *string `json:"name"`
}

type VenueReference struct {
	ID      *int    `json:"id"`
	Name    *string `json:"name"`
	City    *string `json:"city"`
	Country *string `json:"country"`
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
	Height           *string          `json:"height"`
	Weight           *float64         `json:"weight"`
	BirthDate        *string          `json:"birthDate"`
	Age              *int             `json:"age"`
	BirthCity        *string          `json:"birthCity"`
	BirthCountry     *string          `json:"birthCountry"`
	Rookie           *bool            `json:"rookie"`
	College          *string          `json:"college"`
	Twitter          *string          `json:"twitter"`
	Handedness       *Handedness      `json:"handedness"`
	Drafted          *Drafted         `json:"drafted"`
	OfficialImageSrc *string          `json:"officialImageSrc"`
	ExternalMapping  *ExternalMapping `json:"externalMapping"`
}

type StatReference struct {
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

type ExternalMapping struct {
	Source *string `json:"source"`
	ID     *int    `json:"id"`
}

type GameLog struct {
	Game   *GameLogGame `json:"game"`
	Player *Player      `json:"player"`
	Team   *Team        `json:"team"`
	Stats  *Stats       `json:"stats"`
}

type GameLogGame struct {
	ID                   int    `json:"id"`
	StartTime            string `json:"startTime"`
	AwayTeamAbbreviation string `json:"awayTeamAbbreviation"`
	HomeTeamAbbreviation string `json:"homeTeamAbbreviation"`
}

type Player struct {
	ID                  int            `json:"id"`
	FirstName           string         `json:"firstName"`
	LastName            string         `json:"lastName"`
	Position            string         `json:"position"`
	JerseyNumber        *int           `json:"jerseyNumber"`
	CurrentTeam         *Team          `json:"currentTeam"`
	CurrentRosterStatus *string        `json:"currentRosterStatus"`
	CurrentContractYear *bool          `json:"currentContractYear"`
	CurrentInjury       *CurrentInjury `json:"currentInjury"`
}

type Team struct {
	ID           int    `json:"id"`
	Abbreviation string `json:"abbreviation"`
}

type StatsHomeAway struct {
	Away *BoxscoreStats `json:"away"`
	Home *BoxscoreStats `json:"home"`
}

type BoxscoreStats struct {
	TeamStats *[]Stats   `json:"teamStats"`
	Players   *[]Players `json:"players"`
}

type Players struct {
	Player       *PlayerReference `json:"player"`
	TeamAsOfDate *Team            `json:"teamAsOfDate"`
	PlayerStats  *[]Stats         `json:"playerStats"`
}

type Stats struct {
	GamesPlayed   *int           `json:"gamesPlayed"`
	Batting       *Batting       `json:"batting"`
	Pitching      *Pitching      `json:"pitching"`
	Fielding      *Fielding      `json:"fielding"`
	Standings     *Standings     `json:"standings"`
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
	Wins                          *int     `json:"wins"`
	Losses                        *int     `json:"losses"`
	EarnedRunAvg                  *float64 `json:"earnedRunAvg"`
	Saves                         *int     `json:"saves"`
	SaveOpportunities             *int     `json:"SaveOpportunities"`
	InningsPitched                *float64 `json:"inningsPitched"`
	HitsAllowed                   *int     `json:"hitsAllowed"`
	SecondBaseHitsAllowed         *int     `json:"secondBaseHitsAllowed"`
	ThirdBaseHitsAllowed          *int     `json:"thirdBaseHitsAllowed"`
	RunsAllowed                   *int     `json:"runsAllowed"`
	EarnedRunsAllowed             *int     `json:"earnedRunsAllowed"`
	HomerunsAllowed               *int     `json:"homerunsAllowed"`
	PitcherWalks                  *int     `json:"pitcherWalks"`
	PitcherSwings                 *int     `json:"pitcherSwings"`
	PitcherStrikes                *int     `json:"pitcherStrikes"`
	PitcherStrikesFoul            *int     `json:"pitcherStrikesFoul"`
	PitcherStrikesMiss            *int     `json:"pitcherStrikesMiss"`
	PitcherStrikesLooking         *int     `json:"pitcherStrikesLooking"`
	PitcherGroundBalls            *int     `json:"pitcherGroundBalls"`
	PitcherFlyBalls               *int     `json:"pitcherFlyBalls"`
	PitcherLineDrives             *int     `json:"pitcherLineDrives"`
	Pitcher2SeamFastballs         *int     `json:"pitcher2SeamFastballs"`
	Pitcher4SeamFastballs         *int     `json:"pitcher4SeamFastballs"`
	PitcherCurveballs             *int     `json:"pitcherCurveballs"`
	PitcherChangeups              *int     `json:"pitcherChangeups"`
	PitcherCutters                *int     `json:"pitcherCutters"`
	PitcherSliders                *int     `json:"pitcherSliders"`
	PitcherSinkers                *int     `json:"pitcherSinkers"`
	PitcherSplitters              *int     `json:"pitcherSplitters"`
	PitcherSacrificeBunts         *int     `json:"pitcherSacrificeBunts"`
	PitcherSacrificeFlies         *int     `json:"pitcherSacrificeFlies"`
	PitcherStrikeouts             *int     `json:"pitcherStrikeouts"`
	PitchingAvg                   *float64 `json:"pitchingAvg"`
	WalksAndHitsPerInningPitched  *float64 `json:"walksAndHitsPerInningPitched"`
	CompletedGames                *int     `json:"completedGames"`
	Shutouts                      *int     `json:"shutouts"`
	BattersHit                    *int     `json:"battersHit"`
	PitcherIntentionalWalks       *int     `json:"pitcherIntentionalWalks"`
	GamesFinished                 *int     `json:"gamesFinished"`
	Holds                         *int     `json:"holds"`
	PitcherDoublePlays            *int     `json:"pitcherDoublePlays"`
	PitcherTriplePlays            *int     `json:"pitcherTriplePlays"`
	PitcherGroundOuts             *int     `json:"pitcherGroundOuts"`
	PitcherFlyOuts                *int     `json:"pitcherFlyOuts"`
	PitcherWildPitches            *int     `json:"pitcherWildPitches"`
	Balks                         *int     `json:"balks"`
	PitcherStolenBasesAllowed     *int     `json:"pitcherStolenBasesAllowed"`
	PitcherCaughtStealing         *int     `json:"pitcherCaughtStealing"`
	Pickoffs                      *int     `json:"pickoffs"`
	PickoffAttempts               *int     `json:"pickoffAttempts"`
	TotalBattersFaced             *int     `json:"totalBattersFaced"`
	PitchesThrown                 *int     `json:"pitchesThrown"`
	WinPct                        *float64 `json:"winPct"`
	PitcherGroundOutToFlyOutRatio *float64 `json:"pitcherGroundOutToFlyOutRatio"`
	PitcherOnBasePct              *float64 `json:"pitcherOnBasePct"`
	PitcherSluggingPct            *float64 `json:"pitcherSluggingPct"`
	PitcherOnBasePlusSluggingPct  *float64 `json:"pitcherOnBasePlusSluggingPct"`
	StrikeoutsPer9Innings         *float64 `json:"strikeoutsPer9Innings"`
	WalksAllowedPer9Innings       *float64 `json:"walksAllowedPer9Innings"`
	HitsAllowedPer9Innings        *float64 `json:"hitsAllowedPer9Innings"`
	StrikeoutsToWalksRatio        *float64 `json:"strikeoutsToWalksRatio"`
	PitchesPerInning              *float64 `json:"pitchesPerInning"`
	PitcherAtBats                 *int     `json:"pitcherAtBats"`
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
	DefenceEfficiencyRatio    *float64 `json:"defenceEfficiencyRatio"`
	RangeFactor               *float64 `json:"rangeFactor"`
}

type Standings struct {
	Wins            *int     `json:"wins"`
	Losses          *int     `json:"losses"`
	WinPct          *float64 `json:"winPct"`
	GamesBack       *float64 `json:"gamesBack"`
	RunsFor         *int     `json:"runsFor"`
	RunsAgainst     *int     `json:"runsAgainst"`
	RunDifferential *float64 `json:"runDifferential"`
}

type Miscellaneous struct {
	GamesStarted *int `json:"gamesStarted"`
}

type TeamLineup struct {
	Team     *Team     `json:"team"`
	Expected *Expected `json:"expected"`
	Actual   *Actual   `json:"actual"`
}

type Expected struct {
	LineupPositions *[]LineupPosition `json:"lineupPositions"`
}

type Actual struct {
	LineupPositions *[]LineupPosition `json:"lineupPositions"`
}

type LineupPosition struct {
	Position *string `json:"position"`
	Player   *Player `json:"player"`
}

type AtBat struct {
	Inning      *int         `json:"inning"`
	InningHalf  *string      `json:"inningHalf"`
	BattingTeam *Team        `json:"battingTeam"`
	AtBatPlay   *[]AtBatPlay `json:"atBatPlay"`
}

type AtBatPlay struct {
	BatterUp *BatterUp `json:"batterUp"`
	// TODO: fill these out
	// Pitch          *Pitch          `json:"pitch"`
	// PickoffAttempt *PickoffAttempt `json:"pickoffAttempt"`
	// Hit            *Hit            `json:"hit"`
	// BaseRunAttempt *BaseRunAttempt `json:"baseRunAttempt"`
	// BallThrow      *BallThrow      `json:"ballThrow"`
	Description *string `json:"description"`
}

type BatterUp struct {
	BattingPlayer       *Player `json:"battingPlayer"`
	StandingLeftOrRight *string `json:"standingLeftOrRight"`
	Result              *string `json:"result"`
}

type Season struct {
	Name                 *string          `json:"name"`
	Slug                 *string          `json:"slug"`
	StartDate            *string          `json:"startDate"`
	EndDate              *string          `json:"endDate"`
	SeasonInterval       *string          `json:"seasonInterval"`
	SupportedTeamStats   *[]StatReference `json:"supportedTeamStats"`
	SupportedPlayerStats *[]StatReference `json:"supportedPlayerStats"`
}

type CurrentInjury struct {
	Description        *string `json:"description"`
	PlayingProbability *string `json:"playingProbability"`
}

type FeedUpdate struct {
	Feed          *Feed      `json:"feed"`
	ForDate       *[]ForDate `json:"forDates"`
	ForGames      *[]ForGame `json:"forGames"`
	LastUpdatedOn *string    `json:"lastUpdatedOn"`
}

type Feed struct {
	Name         *string `json:"name"`
	Version      *string `json:"version"`
	Abbreviation *string `json:"abbreviation"`
	Description  *string `json:"description"`
}

type ForDate struct {
	ForDate       *string `json:"forDate"`
	LastUpdatedOn *string `json:"lastUpdatedOn"`
}

type ForGame struct {
	Game          *Schedule `json:"game"`
	LastUpdatedOn *string   `json:"lastUpdatedOn"`
}

type TeamStatsTotal struct {
	Team  *TeamReference `json:"team"`
	Stats *Stats         `json:"stats"`
}

type PlayerStatsTotal struct {
	Player *PlayerReference `json:"player"`
	Team   *Team            `json:"team"`
	Stats  *Stats           `json:"stats"`
}

type Venues struct {
	Venue    *VenueReference `json:"venue"`
	HomeTeam *TeamReference  `json:"homeTeam"`
}

type Drafted struct {
	Year        *int  `json:"year"`
	Team        *Team `json:"team"`
	PickTeam    *Team `json:"pickTeam"`
	Round       *int  `json:"round"`
	RoundPick   *int  `json:"roundPick"`
	OverallPick *int  `json:"overallPick"`
}

type Teams struct {
	Team           *TeamReference `json:"team"`
	Stats          *Stats         `json:"stats"`
	OverallRank    *Rank          `json:"overallRank"`
	ConferenceRank *Rank          `json:"conferenceRank"`
	DivisionRank   *Rank          `json:"divisionRank"`
	PlayoffRank    *Rank          `json:"playoffRank"`
}

type Rank struct {
	ConferenceName *string  `json:"conferenceName"`
	DivisionName   *string  `json:"divisionName"`
	AppliesTo      *string  `json:"appliesTo"`
	Rank           *int     `json:"rank"`
	GamesBack      *float64 `json:"gamesBack"`
}
