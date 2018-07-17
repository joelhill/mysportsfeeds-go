package msf

import (
	"context"
	"os"

	blaster "github.com/joelhill/go-rest-http-blaster"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/ssh/terminal"
)

// IService - all functions that implement the My Sports Feed Service Interface.
type IService interface {
	NewDailyDfsOptions() *DailyDfsOptions
	DailyDfs(c context.Context, options *DailyDfsOptions) (DfsIO, error)

	NewDailyGamesOptions() *DailyGamesOptions
	DailyGames(c context.Context, options *DailyGamesOptions) (GamesIO, error)

	NewDailyPlayerGamelogsOptions() *DailyPlayerGamelogsOptions
	DailyPlayerGamelogs(c context.Context, options *DailyPlayerGamelogsOptions) (GameLogIO, error)

	NewDailyTeamGamelogsOptions() *DailyTeamGamelogsOptions
	DailyTeamGamelogs(c context.Context, options *DailyTeamGamelogsOptions) (GameLogIO, error)

	NewGameBoxscoreOptions() *GameBoxscoreOptions
	GameBoxscore(c context.Context, options *GameBoxscoreOptions) (BoxscoreIO, error)

	NewGameLineupOptions() *GameLineupOptions
	GameLineup(c context.Context, options *GameLineupOptions) (BoxscoreIO, error)

	NewSeasonalDfsOptions() *SeasonalDfsOptions
	SeasonalDfs(c context.Context, options *SeasonalDfsOptions) (DfsIO, error)

	NewSeasonalGamesOptions() *SeasonalGamesOptions
	SeasonalGames(c context.Context, options *SeasonalGamesOptions) (GamesIO, error)

	NewSeasonalPlayerGamelogsOptions() *SeasonalPlayerGamelogsOptions
	SeasonalPlayerGamelogs(c context.Context, options *SeasonalPlayerGamelogsOptions) (GameLogIO, error)

	NewSeasonalTeamGamelogsOptions() *SeasonalTeamGamelogsOptions
	SeasonalTeamGamelogs(c context.Context, options *SeasonalTeamGamelogsOptions) (GameLogIO, error)
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
		ServiceName:    "go-mysportsfeeds",
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
