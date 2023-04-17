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

// CurrentSeasonOptions - Are the options to hit the current season endpoint
type CurrentSeasonOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Format  string

	// Optional URL Params
	Date  string
	Force string
}

// NewCurrentSeasonOptions - Returns the options with most url parts already set to hit the current season endpoint
func (s *Service) NewCurrentSeasonOptions() *CurrentSeasonOptions {
	return &CurrentSeasonOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
	}
}

// CurrentSeason - hits the https://api.mysportsfeeds.com/{version}/pull/{sport}/current_season.{format} endpoint
func (s *Service) CurrentSeason(options *CurrentSeasonOptions) (CurrentSeasonIO, int, error) {
	mapping := CurrentSeasonIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateCurrentSeasonURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/current_season.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Format, cacheBuster)

	if len(options.Date) > 0 {
		uri = fmt.Sprintf("%s&date=%s", uri, options.Date)
	}

	if len(options.Force) > 0 {
		uri = fmt.Sprintf("%s&force=%s", uri, options.Force)
	}

	s.Logger = s.Logger.WithFields(logrus.Fields{
		"URI": uri,
	})
	s.Logger.Debug("CurrentSeason API Call")

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
		s.Logger.Errorf("client: something went wrong making the get request for CurrentSeason: %s", err.Error())
		return mapping, response.StatusCode, err
	}

	s.Logger.Infof("CurrentSeason Status Code: %d", response.StatusCode)

	if jErr := json.NewDecoder(response.Body).Decode(&mapping); jErr != nil {
		s.Logger.Errorf("client: error decoding response for CurrentSeason: %s", err.Error())
		return mapping, response.StatusCode, err
	}

	return mapping, response.StatusCode, nil
}

func validateCurrentSeasonURI(options *CurrentSeasonOptions) error {
	if len(options.URL) == 0 {
		return errors.New("missing required option to build the url: URL")
	}
	if len(options.Version) == 0 {
		return errors.New("missing required option to build the url: Version")
	}
	if len(options.Sport) == 0 {
		return errors.New("missing required option to build the url: Sport")
	}
	if len(options.Format) == 0 {
		return errors.New("missing required option to build the url: Format")
	}
	return nil
}
