package msf

import (
	"context"
	"errors"
	"fmt"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// GamePlayByPlayOptions - Are the options to hit the seasonal DFS endpoint
type GamePlayByPlayOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Game    string
	Format  string

	// Optional URL Params
	Playtype string
	Sort     string
	Offset   string
	Limit    string
	Force    string
}

// NewGamePlayByPlayOptions - Returns the options with most url parts already set to hit the game lineup endpoint
func (s *Service) NewGamePlayByPlayOptions() *GamePlayByPlayOptions {
	return &GamePlayByPlayOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// GamePlayByPlay - hits the https://api.mysportsfeeds.com/{version}/pull/{sport}/{season}/games/{game}/playbyplay.{format} endoint
func (s *Service) GamePlayByPlay(c context.Context, options *GamePlayByPlayOptions) (GamePlayByPlayIO, error) {
	errorPayload := make(map[string]interface{})
	mapping := GamePlayByPlayIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateGamePlayByPlayURI(options)
	if err != nil {
		return mapping, err
	}

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/games/%s/playbyplay.%s?1=1", options.URL, options.Version, options.Sport, options.Season, options.Game, options.Format)

	if len(options.Playtype) > 0 {
		uri = fmt.Sprintf("%s&playtype=%s", uri, options.Playtype)
	}

	if len(options.Sort) > 0 {
		uri = fmt.Sprintf("%s&sort=%s", uri, options.Sort)
	}

	if len(options.Offset) > 0 {
		uri = fmt.Sprintf("%s&offset=%s", uri, options.Offset)
	}

	if len(options.Limit) > 0 {
		uri = fmt.Sprintf("%s&limit=%s", uri, options.Limit)
	}

	if len(options.Force) > 0 {
		uri = fmt.Sprintf("%s&force=%s", uri, options.Force)
	}

	s.Logger = s.Logger.WithFields(logrus.Fields{
		"URI": uri,
	})
	s.Logger.Debug("GamePlayByPlay API Call")

	// make you a client
	client, err := blaster.NewClient(uri)
	if err != nil {
		s.Logger.Errorf("failed to create a http client: %s", err.Error())
		return mapping, err
	}

	client.SetHeader("Authorization", s.Config.Authorization)
	client.WillSaturateOnError(&errorPayload)
	client.WillSaturate(&mapping)

	ctx := context.Background()
	statusCode, err := client.Get(ctx)
	if err != nil {
		s.Logger.Errorf("something went wrong making the get request for GamePlayByPlay: %s", err.Error())
		return mapping, err
	}

	s.Logger.Infof("GamePlayByPlay Status Code: %d", statusCode)

	return mapping, nil
}

func validateGamePlayByPlayURI(options *GamePlayByPlayOptions) error {
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
