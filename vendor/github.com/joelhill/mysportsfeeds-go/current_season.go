package msf

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
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
	errorPayload := make(map[string]interface{})
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
		s.Logger.Errorf("something went wrong making the get request for CurrentSeason: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("CurrentSeason Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("CurrentSeason retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
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
