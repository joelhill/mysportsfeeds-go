package msf

import (
	"context"
	"os"

	blaster "github.com/joelhill/go-rest-http-blaster"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

const (
	// type of compression header
	CompressionHeaderGzip = "gzip"
)

// IService - all functions that implement the My Sports Feed Service Interface.
type IService interface {
	NewCurrentSeasonOptions() *CurrentSeasonOptions
	CurrentSeason(c context.Context, options *CurrentSeasonOptions) (CurrentSeasonIO, error)

	NewDailyDfsOptions() *DailyDfsOptions
	DailyDfs(c context.Context, options *DailyDfsOptions) (DfsIO, error)

	NewDailyGamesOptions() *DailyGamesOptions
	DailyGames(c context.Context, options *DailyGamesOptions) (GamesIO, error)

	NewDailyPlayerGamelogsOptions() *DailyPlayerGamelogsOptions
	DailyPlayerGamelogs(c context.Context, options *DailyPlayerGamelogsOptions) (GameLogIO, error)

	NewDailyStandingsOptions() *DailyStandingsOptions
	DailyStandings(c context.Context, options *DailyStandingsOptions) (StandingsIO, error)

	NewDailyTeamGamelogsOptions() *DailyTeamGamelogsOptions
	DailyTeamGamelogs(c context.Context, options *DailyTeamGamelogsOptions) (GameLogIO, error)

	NewFeedUpdatesOptions() *FeedUpdatesOptions
	FeedUpdates(c context.Context, options *FeedUpdatesOptions) (FeedUpdatesIO, error)

	NewGameBoxscoreOptions() *GameBoxscoreOptions
	GameBoxscore(c context.Context, options *GameBoxscoreOptions) (BoxscoreIO, error)

	NewGameLineupOptions() *GameLineupOptions
	GameLineup(c context.Context, options *GameLineupOptions) (BoxscoreIO, error)

	NewGamePlayByPlayOptions() *GamePlayByPlayOptions
	GamePlayByPlay(c context.Context, options *GamePlayByPlayOptions) (GamePlayByPlayIO, error)

	NewPlayerInjuriesOptions() *PlayerInjuriesOptions
	PlayerInjuries(c context.Context, options *PlayerInjuriesOptions) (PlayerInjuriesIO, error)

	NewPlayersOptions() *PlayersOptions
	Players(c context.Context, options *PlayersOptions) (PlayersIO, error)

	NewSeasonalDfsOptions() *SeasonalDfsOptions
	SeasonalDfs(c context.Context, options *SeasonalDfsOptions) (DfsIO, error)

	NewSeasonalGamesOptions() *SeasonalGamesOptions
	SeasonalGames(c context.Context, options *SeasonalGamesOptions) (GamesIO, error)

	NewSeasonalPlayerGamelogsOptions() *SeasonalPlayerGamelogsOptions
	SeasonalPlayerGamelogs(c context.Context, options *SeasonalPlayerGamelogsOptions) (GameLogIO, error)

	NewSeasonalPlayerStatsOptions() *SeasonalPlayerStatsOptions
	SeasonalPlayerStats(c context.Context, options *SeasonalPlayerStatsOptions) (PlayerStatsTotalsIO, error)

	NewSeasonalTeamGamelogsOptions() *SeasonalTeamGamelogsOptions
	SeasonalTeamGamelogs(c context.Context, options *SeasonalTeamGamelogsOptions) (GameLogIO, error)

	NewSeasonalTeamStatsOptions() *SeasonalTeamStatsOptions
	SeasonalTeamStats(c context.Context, options *SeasonalTeamStatsOptions) (TeamStatsTotalsIO, error)

	NewSeasonalVenuesOptions() *SeasonalVenuesOptions
	SeasonalVenues(c context.Context, options *SeasonalVenuesOptions) (VenuesIO, error)
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
