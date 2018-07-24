package msf

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// SeasonalPlayerStatsOptions - Are the options to hit the seasonal player stats endpoint
type SeasonalPlayerStatsOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Format  string

	// Optional URL Params
	Player   string
	Position string
	Country  string
	Team     string
	Date     string
	Stats    string
	Sort     string
	Offset   string
	Limit    string
	Force    string
}

// NewSeasonalPlayerStatsOptions - Returns the options with most url parts already set to hit the seasonal player stats endpoint
func (s *Service) NewSeasonalPlayerStatsOptions() *SeasonalPlayerStatsOptions {
	return &SeasonalPlayerStatsOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// SeasonalPlayerStats - hits the https://api.mysportsfeeds.com/{version}/pull/{sport}/{season}/player_stats_totals.{format} endoint
func (s *Service) SeasonalPlayerStats(c context.Context, options *SeasonalPlayerStatsOptions) (PlayerStatsTotalsIO, int, error) {
	errorPayload := make(map[string]interface{})
	mapping := PlayerStatsTotalsIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateSeasonalPlayerStatsURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/player_stats_totals.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Season, options.Format, cacheBuster)

	if len(options.Player) > 0 {
		uri = fmt.Sprintf("%s&player=%s", uri, options.Player)
	}

	if len(options.Position) > 0 {
		uri = fmt.Sprintf("%s&position=%s", uri, options.Position)
	}

	if len(options.Country) > 0 {
		uri = fmt.Sprintf("%s&country=%s", uri, options.Country)
	}

	if len(options.Team) > 0 {
		uri = fmt.Sprintf("%s&team=%s", uri, options.Team)
	}

	if len(options.Date) > 0 {
		uri = fmt.Sprintf("%s&date=%s", uri, options.Date)
	}

	if len(options.Stats) > 0 {
		uri = fmt.Sprintf("%s&stats=%s", uri, options.Stats)
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
	s.Logger.Debug("SeasonalPlayerStats API Call")

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

	ctx := context.Background()
	statusCode, err := client.Get(ctx)
	if err != nil {
		s.Logger.Errorf("something went wrong making the get request for SeasonalPlayerStats: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("SeasonalPlayerStats Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("SeasonalPlayerStats retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
}

func validateSeasonalPlayerStatsURI(options *SeasonalPlayerStatsOptions) error {
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
	if len(options.Format) == 0 {
		return errors.New("missing required option to build the url: Format")
	}
	return nil
}
