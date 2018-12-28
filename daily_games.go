package msf

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// DailyGamesOptions - Are the options to hit the daily games endpoint
type DailyGamesOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Date    string //YYYYMMDD
	Format  string

	// Optional URL Params
	Team   string
	Status string
	Sort   string
	Offset string
	Limit  string
	Force  string
}

// NewDailyGamesOptions - Returns the options with most url parts already set to hit the daily games endpoint
func (s *Service) NewDailyGamesOptions() *DailyGamesOptions {
	return &DailyGamesOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// DailyGames - hits the https://api.mysportsfeeds.com/{version}/pull/{sport}/{season}/date/{date}/games.{format} endoint
func (s *Service) DailyGames(options *DailyGamesOptions) (GamesIO, int, error) {
	errorPayload := make(map[string]interface{})
	mapping := GamesIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateDailyGamesURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/date/%s/games.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Season, options.Date, options.Format, cacheBuster)

	if len(options.Team) > 0 {
		uri = fmt.Sprintf("%s&team=%s", uri, options.Team)
	}

	if len(options.Status) > 0 {
		uri = fmt.Sprintf("%s&status=%s", uri, options.Status)
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
	s.Logger.Debug("DailyGames API Call")

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
		s.Logger.Errorf("something went wrong making the get request for DailyGames: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("DailyGames Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("DailyGames retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
}

func validateDailyGamesURI(options *DailyGamesOptions) error {
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
	if len(options.Date) == 0 {
		return errors.New("missing required option to build the url: Date")
	}
	if len(options.Format) == 0 {
		return errors.New("missing required option to build the url: Format")
	}
	return nil
}
