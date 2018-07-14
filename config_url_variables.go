package msf

// ADDRESS
const (
	// URL - api url
	URL = "https://api.mysportsfeeds.com"
)

// VERSION
const (
	// VersionV1_1 - v1.1
	VersionV1_1 = "v1.1"

	// VersionV1_2 - v1.2
	VersionV1_2 = "v1.2"

	// VersionV2_0 - v2.0
	VersionV2_0 = "v2.0"
)

// SPORT
const (
	// SportMLB - baseball
	SportMLB = "mlb"

	// SportNFL - football
	SportNFL = "nfl"

	// SportNBA - Basketball
	SportNBA = "nba"

	// SportNHL - Hockey
	SportNHL = "nhl"
)

// RETURN FORMAT
const (
	// FormatJSON - returns the format in json
	FormatJSON = "json"

	// FormatXML - returns the format in xml
	FormatXML = "xml"

	// FormatCSV - returns the format in csv
	FormatCSV = "csv"
)

// SEASON
const (
	// SeasonCurrent - current season
	SeasonCurrent = "current"

	// SeasonLatest - latest season
	SeasonLatest = "latest"

	// SeasonUpcoming - upcoming season
	SeasonUpcoming = "upcoming"
)

// DATE
const (
	// DateYesterday -
	DateYesterday = "yesterday"

	// DateToday -
	DateToday = "today"

	// DateTomorrow -
	DateTomorrow = "tomorrow"
)

// STATUS
const (
	// StatusUnplayed - returns games scheduled but not yet started
	StatusUnplayed = "unplayed"

	// StatusInProgress - returns games currently underway
	StatusInProgress = "in-progress"

	// StatusPostgameReviewing - returns game is over, but we're reviewing against official sources
	StatusPostgameReviewing = "postgame-reviewing"

	// StatusFinal - returns game is final and reviewed
	StatusFinal = "final"
)
