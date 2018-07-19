package msf

import (
	"context"
	"errors"
	"fmt"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// SeasonalPlayerGamelogsOptions - Are the options to hit the seasonal games endpoint
type SeasonalPlayerGamelogsOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Format  string

	// Optional URL Params
	Team     string
	Player   string
	Position string
	Game     string
	Date     string
	Stats    string
	Sort     string
	Offset   string
	Limit    string
	Force    string
}

// NewSeasonalPlayerGamelogsOptions - Returns the options with most url parts already set to hit the seasonal games endpoint
func (s *Service) NewSeasonalPlayerGamelogsOptions() *SeasonalPlayerGamelogsOptions {
	return &SeasonalPlayerGamelogsOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// SeasonalPlayerGamelogs - hits the https://api.mysportsfeeds.com/v2.0/pull/mlb/{season}/player_gamelogs.{format} endoint
func (s *Service) SeasonalPlayerGamelogs(c context.Context, options *SeasonalPlayerGamelogsOptions) (GameLogIO, error) {
	errorPayload := make(map[string]interface{})
	mapping := GameLogIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateSeasonalPlayerGamelogsURI(options)
	if err != nil {
		return mapping, err
	}

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/player_gamelogs.%s?1=1", options.URL, options.Version, options.Sport, options.Season, options.Format)

	if len(options.Team) > 0 {
		uri = fmt.Sprintf("%s&team=%s", uri, options.Team)
	}

	if len(options.Player) > 0 {
		uri = fmt.Sprintf("%s&player=%s", uri, options.Player)
	}

	if len(options.Position) > 0 {
		uri = fmt.Sprintf("%s&position=%s", uri, options.Position)
	}

	if len(options.Game) > 0 {
		uri = fmt.Sprintf("%s&game=%s", uri, options.Game)
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
	s.Logger.Debug("SeasonalPlayerGamelogs API Call")

	// make you a client
	client, err := blaster.NewClient(uri)
	if err != nil {
		s.Logger.Errorf("failed to create a http client: %s", err.Error())
		return mapping, err
	}

	client.SetHeader("Accept-Encoding", CompressionHeaderGzip)
	client.SetHeader("Authorization", s.Config.Authorization)
	client.WillSaturateOnError(&errorPayload)
	client.WillSaturate(&mapping)

	ctx := context.Background()
	statusCode, err := client.Get(ctx)
	if err != nil {
		s.Logger.Errorf("something went wrong making the get request for SeasonalPlayerGamelogs: %s", err.Error())
		return mapping, err
	}

	s.Logger.Infof("SeasonalPlayerGamelogs Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("SeasonalPlayerGamelogs retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, nil
}

func validateSeasonalPlayerGamelogsURI(options *SeasonalPlayerGamelogsOptions) error {
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
