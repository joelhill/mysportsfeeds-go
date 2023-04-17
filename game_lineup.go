package msf

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
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
	client := retryablehttp.NewClient()

	req, err := retryablehttp.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		s.Logger.Errorf("client: could not create request: %s", err.Error())
		return mapping, 0, err
	}
	req.Header.Add("Authorization", CompressionHeaderGzip)
	req.Header.Add("Authorization", s.Config.Authorization)

	response, err := client.Do(req)
	if err != nil {
		s.Logger.Errorf("client: error making http request: %s", err.Error())
		return mapping, 0, err
	}

	if response.StatusCode < 200 || response.StatusCode > 300 {
		s.Logger.Errorf("client: something went wrong making the get request for GameLineup: %s", err.Error())
		return mapping, response.StatusCode, err
	}

	s.Logger.Infof("GameLineup Status Code: %d", response.StatusCode)

	if jErr := json.NewDecoder(response.Body).Decode(&mapping); jErr != nil {
		s.Logger.Errorf("client: error decoding response for GameLineup: %s", err.Error())
		return mapping, response.StatusCode, err
	}

	return mapping, response.StatusCode, nil
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
