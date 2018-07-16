package msf

import (
	"context"
	"fmt"

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

// DefaultDailyGamesOptions - Returns the default options to hit the daily games endpoint
func DefaultDailyGamesOptions() *DailyGamesOptions {
	return &DailyGamesOptions{
		URL:     URL,
		Version: VersionV2_0,
		Sport:   SportMLB,
		Format:  FormatJSON,
		Season:  SeasonCurrent,
		Date:    DateToday,
	}
}

// DailyGames - hits the https://api.mysportsfeeds.com/v2.0/pull/mlb/{season}/date/{date}/games.{format} endoint
func (s *Service) DailyGames(c context.Context, options *DailyGamesOptions) (GamesIO, error) {
	errorPayload := make(map[string]interface{})
	mapping := GamesIO{}

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/date/%s/games.%s?1=1", options.URL, options.Version, options.Sport, options.Season, options.Date, options.Format)

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
		return mapping, err
	}

	client.SetHeader("Authorization", s.Config.Authorization)
	client.WillSaturateOnError(&errorPayload)
	client.WillSaturate(&mapping)

	ctx := context.Background()
	statusCode, err := client.Get(ctx)
	if err != nil {
		s.Logger.Errorf("something went wrong making the get request for DailyGames: %s", err.Error())
		return mapping, err
	}

	s.Logger.Infof("DailyGames Status Code: %d", statusCode)

	return mapping, nil
}
