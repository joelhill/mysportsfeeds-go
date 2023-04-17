package msf

import (
	"os"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	// CompressionHeaderGzip - type of compression header
	CompressionHeaderGzip = "gzip"
)

// IService - all functions that implement the My Sports Feed Service Interface.
type IService interface {
	NewCurrentSeasonOptions() *CurrentSeasonOptions
	CurrentSeason(options *CurrentSeasonOptions) (CurrentSeasonIO, int, error)

	NewDailyDfsOptions() *DailyDfsOptions
	DailyDfs(options *DailyDfsOptions) (DfsIO, int, error)

	NewDailyGamesOptions() *DailyGamesOptions
	DailyGames(options *DailyGamesOptions) (GamesIO, int, error)

	NewDailyPlayerGamelogsOptions() *DailyPlayerGamelogsOptions
	DailyPlayerGamelogs(options *DailyPlayerGamelogsOptions) (GameLogIO, int, error)

	NewDailyStandingsOptions() *DailyStandingsOptions
	DailyStandings(options *DailyStandingsOptions) (StandingsIO, int, error)

	NewDailyTeamGamelogsOptions() *DailyTeamGamelogsOptions
	DailyTeamGamelogs(options *DailyTeamGamelogsOptions) (GameLogIO, int, error)

	NewFeedUpdatesOptions() *FeedUpdatesOptions
	FeedUpdates(options *FeedUpdatesOptions) (FeedUpdatesIO, int, error)

	NewGameBoxscoreOptions() *GameBoxscoreOptions
	GameBoxscore(options *GameBoxscoreOptions) (BoxscoreIO, int, error)

	NewGameLineupOptions() *GameLineupOptions
	GameLineup(options *GameLineupOptions) (GameLineupIO, int, error)

	NewGamePlayByPlayOptions() *GamePlayByPlayOptions
	GamePlayByPlay(options *GamePlayByPlayOptions) (GamePlayByPlayIO, int, error)

	NewPlayerInjuriesOptions() *PlayerInjuriesOptions
	PlayerInjuries(options *PlayerInjuriesOptions) (PlayerInjuriesIO, int, error)

	NewPlayersOptions() *PlayersOptions
	Players(options *PlayersOptions) (PlayersIO, int, error)

	NewSeasonalDfsOptions() *SeasonalDfsOptions
	SeasonalDfs(options *SeasonalDfsOptions) (DfsIO, int, error)

	NewSeasonalGamesOptions() *SeasonalGamesOptions
	SeasonalGames(options *SeasonalGamesOptions) (GamesIO, int, error)

	NewSeasonalPlayerGamelogsOptions() *SeasonalPlayerGamelogsOptions
	SeasonalPlayerGamelogs(options *SeasonalPlayerGamelogsOptions) (GameLogIO, int, error)

	NewSeasonalPlayerStatsOptions() *SeasonalPlayerStatsOptions
	SeasonalPlayerStats(options *SeasonalPlayerStatsOptions) (PlayerStatsTotalsIO, int, error)

	NewSeasonalTeamGamelogsOptions() *SeasonalTeamGamelogsOptions
	SeasonalTeamGamelogs(options *SeasonalTeamGamelogsOptions) (GameLogIO, int, error)

	NewSeasonalTeamStatsOptions() *SeasonalTeamStatsOptions
	SeasonalTeamStats(options *SeasonalTeamStatsOptions) (TeamStatsTotalsIO, int, error)

	NewSeasonalVenuesOptions() *SeasonalVenuesOptions
	SeasonalVenues(options *SeasonalVenuesOptions) (VenuesIO, int, error)
}

// Service - dependencies for the my sports feed service
type Service struct {
	Config *Config
	Logger *logrus.Entry
}

// NewService -
func NewService(config *Config) *Service {

	// JSON formatter for log output if not running in a TTY
	// because Loggly likes JSON but humans like colors
	if !terminal.IsTerminal(int(os.Stderr.Fd())) {
		logrus.SetFormatter(&logrus.JSONFormatter{})
		logrus.SetOutput(os.Stderr)
	}

	logLevel, err := logrus.ParseLevel("debug")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error_message": err.Error(),
		}).Error("log level error")
	}
	logrus.SetLevel(logLevel)

	l := logrus.New()
	lg := logrus.NewEntry(l)

	return &Service{
		Config: config,
		Logger: lg,
	}
}
