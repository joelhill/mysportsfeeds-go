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
	SeasonalGames(c context.Context, options *SeasonalGamesOptions) (GamesIO, error)
}

// Service - dependencies for the my sports feed service
type Service struct {
	Config *Config
	Logger *logrus.Entry
}

// newService -
func newService(config *Config) *Service {

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
