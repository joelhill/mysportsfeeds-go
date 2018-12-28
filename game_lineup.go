package msf

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// GameLineupOptions - Are the options to hit the game lineup endpoint
type GameLineupOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Game    string
	Format  string

	// Optional URL Params
	Position   string
	LineupType string
	Force      string
}

// NewGameLineupOptions - Returns the options with most url parts already set to hit the game lineup endpoint
func (s *Service) NewGameLineupOptions() *GameLineupOptions {
	return &GameLineupOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// GameLineup - hits the https://api.mysportsfeeds.com/{version}/pull/{sport}/{season}/games/{game}/lineup.{format} endoint
func (s *Service) GameLineup(options *GameLineupOptions) (GameLineupIO, int, error) {
	errorPayload := make(map[string]interface{})
	mapping := GameLineupIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateGameLineupURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/games/%s/lineup.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Season, options.Game, options.Format, cacheBuster)

	if len(options.Position) > 0 {
		uri = fmt.Sprintf("%s&position=%s", uri, options.Position)
	}

	if len(options.LineupType) > 0 {
		uri = fmt.Sprintf("%s&lineuptype=%s", uri, options.LineupType)
	}

	if len(options.Force) > 0 {
		uri = fmt.Sprintf("%s&force=%s", uri, options.Force)
	}

	s.Logger = s.Logger.WithFields(logrus.Fields{
		"URI": uri,
	})
	s.Logger.Debug("GameLineup API Call")

	// make you a client
	client, err := blaster.NewClient(uri)
	if err != nil {
		s.Logger.Errorf("failed to create a http client: %s", err.Error())
		return mapping, 0, err
	}

	client.SetHeader("Accept-Encoding", CompressionHeaderGzip)
	client.SetHeader("Authorization", s.Config.Authorization)
	client.WillSaturateOnError(&errorPayload)
	client.WillSaturate(&mapping)

	statusCode, err := client.Get(context.Background())
	if err != nil {
		s.Logger.Errorf("something went wrong making the get request for GameLineup: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("GameLineup Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("GameLineup retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
}

func validateGameLineupURI(options *GameLineupOptions) error {
	if len(options.URL) == 0 {
		return errors.New("missing required option to build the url: URL")
	}
	if len(options.Version) == 0 {
		return errors.New("missing required option to build the url: Version")
	}
	if len(options.Sport) == 0 {
		return errors.New("missing required option to build the url: Sport")
	}
	if len(options.Season) == 0 {
		return errors.New("missing required option to build the url: Season")
	}
	if len(options.Game) == 0 {
		return errors.New("missing required option to build the url: Game")
	}
	if len(options.Format) == 0 {
		return errors.New("missing required option to build the url: Format")
	}
	return nil
}
