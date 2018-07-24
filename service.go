package msf

import (
	"context"
	"os"

	blaster "github.com/joelhill/go-rest-http-blaster"
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
	CurrentSeason(c context.Context, options *CurrentSeasonOptions) (CurrentSeasonIO, int, error)

	NewDailyDfsOptions() *DailyDfsOptions
	DailyDfs(c context.Context, options *DailyDfsOptions) (DfsIO, int, error)

	NewDailyGamesOptions() *DailyGamesOptions
	DailyGames(c context.Context, options *DailyGamesOptions) (GamesIO, int, error)

	NewDailyPlayerGamelogsOptions() *DailyPlayerGamelogsOptions
	DailyPlayerGamelogs(c context.Context, options *DailyPlayerGamelogsOptions) (GameLogIO, int, error)

	NewDailyStandingsOptions() *DailyStandingsOptions
	DailyStandings(c context.Context, options *DailyStandingsOptions) (StandingsIO, int, error)

	NewDailyTeamGamelogsOptions() *DailyTeamGamelogsOptions
	DailyTeamGamelogs(c context.Context, options *DailyTeamGamelogsOptions) (GameLogIO, int, error)

	NewFeedUpdatesOptions() *FeedUpdatesOptions
	FeedUpdates(c context.Context, options *FeedUpdatesOptions) (FeedUpdatesIO, int, error)

	NewGameBoxscoreOptions() *GameBoxscoreOptions
	GameBoxscore(c context.Context, options *GameBoxscoreOptions) (BoxscoreIO, int, error)

	NewGameLineupOptions() *GameLineupOptions
	GameLineup(c context.Context, options *GameLineupOptions) (BoxscoreIO, int, error)

	NewGamePlayByPlayOptions() *GamePlayByPlayOptions
	GamePlayByPlay(c context.Context, options *GamePlayByPlayOptions) (GamePlayByPlayIO, int, error)

	NewPlayerInjuriesOptions() *PlayerInjuriesOptions
	PlayerInjuries(c context.Context, options *PlayerInjuriesOptions) (PlayerInjuriesIO, int, error)

	NewPlayersOptions() *PlayersOptions
	Players(c context.Context, options *PlayersOptions) (PlayersIO, int, error)

	NewSeasonalDfsOptions() *SeasonalDfsOptions
	SeasonalDfs(c context.Context, options *SeasonalDfsOptions) (DfsIO, int, error)

	NewSeasonalGamesOptions() *SeasonalGamesOptions
	SeasonalGames(c context.Context, options *SeasonalGamesOptions) (GamesIO, int, error)

	NewSeasonalPlayerGamelogsOptions() *SeasonalPlayerGamelogsOptions
	SeasonalPlayerGamelogs(c context.Context, options *SeasonalPlayerGamelogsOptions) (GameLogIO, int, error)

	NewSeasonalPlayerStatsOptions() *SeasonalPlayerStatsOptions
	SeasonalPlayerStats(c context.Context, options *SeasonalPlayerStatsOptions) (PlayerStatsTotalsIO, int, error)

	NewSeasonalTeamGamelogsOptions() *SeasonalTeamGamelogsOptions
	SeasonalTeamGamelogs(c context.Context, options *SeasonalTeamGamelogsOptions) (GameLogIO, int, error)

	NewSeasonalTeamStatsOptions() *SeasonalTeamStatsOptions
	SeasonalTeamStats(c context.Context, options *SeasonalTeamStatsOptions) (TeamStatsTotalsIO, int, error)

	NewSeasonalVenuesOptions() *SeasonalVenuesOptions
	SeasonalVenues(c context.Context, options *SeasonalVenuesOptions) (VenuesIO, int, error)
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

	blaster.SetDefaults(&blaster.Defaults{
		ServiceName:    "mysportsfeeds-go",
		RequireHeaders: false,
		StatsdRate:     0.0,
	})

	l := logrus.New()
	lg := logrus.NewEntry(l)

	return &Service{
		Config: config,
		Logger: lg,
	}
}
