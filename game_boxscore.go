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

// GameBoxscoreOptions - Are the options to hit the game boxscore endpoint
type GameBoxscoreOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Game    string
	Format  string

	// Optional URL Params
	TeamStats   string
	PlayerStats string
	Sort        string
	Offset      string
	Limit       string
	Force       string
}

// NewGameBoxscoreOptions - Returns the options with most url parts already set to hit the daily games endpoint
func (s *Service) NewGameBoxscoreOptions() *GameBoxscoreOptions {
	return &GameBoxscoreOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// GameBoxscore - hits the https://api.mysportsfeeds.com/{version}/pull/{sport}/{season}/games/{game}/boxscore.{format} endoint
func (s *Service) GameBoxscore(options *GameBoxscoreOptions) (BoxscoreIO, int, error) {

	mapping := BoxscoreIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateGameBoxscoreURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/games/%s/boxscore.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Season, options.Game, options.Format, cacheBuster)

	if len(options.TeamStats) > 0 {
		uri = fmt.Sprintf("%s&teamstats=%s", uri, options.TeamStats)
	}

	if len(options.PlayerStats) > 0 {
		uri = fmt.Sprintf("%s&playerstats=%s", uri, options.PlayerStats)
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
	s.Logger.Debug("GameBoxscore API Call")

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
		s.Logger.Errorf("client: something went wrong making the get request for GameBoxscore: %s", err.Error())
		return mapping, response.StatusCode, err
	}

	s.Logger.Infof("GameBoxscore Status Code: %d", response.StatusCode)

	if jErr := json.NewDecoder(response.Body).Decode(&mapping); jErr != nil {
		s.Logger.Errorf("client: error decoding response for GameBoxscore: %s", err.Error())
		return mapping, response.StatusCode, err
	}

	return mapping, response.StatusCode, nil
}

func validateGameBoxscoreURI(options *GameBoxscoreOptions) error {
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
