package msf

import (
	"context"
	"errors"
	"fmt"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// DailyPlayerGamelogsOptions - Are the options to hit the daily player gamelogs endpoint
type DailyPlayerGamelogsOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Date    string // YYYYMMDD
	Week    string // 1-40
	Format  string

	// Optional URL Params
	Team     string
	Player   string
	Position string
	Game     string
	Stats    string
	Sort     string
	Offset   string
	Limit    string
	Force    string
}

// NewDailyPlayerGamelogsOptions - Returns the options with most url parts already set to hit the seasonal games endpoint
func (s *Service) NewDailyPlayerGamelogsOptions() *DailyPlayerGamelogsOptions {
	return &DailyPlayerGamelogsOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// DailyPlayerGamelogs - hits the https://api.mysportsfeeds.com/{version}/pull/{sport}/{season}/date/{date}/week/{week}/player_gamelogs.{format} endoint
func (s *Service) DailyPlayerGamelogs(c context.Context, options *DailyPlayerGamelogsOptions) (GameLogIO, error) {
	errorPayload := make(map[string]interface{})
	mapping := GameLogIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateDailyPlayerGamelogsURI(options)
	if err != nil {
		return mapping, err
	}

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/date/%s/week/%s/player_gamelogs.%s?1=1", options.URL, options.Version, options.Sport, options.Season, options.Date, options.Week, options.Format)

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
	s.Logger.Debug("DailyPlayerGamelogs API Call")

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
		s.Logger.Errorf("something went wrong making the get request for DailyPlayerGamelogs: %s", err.Error())
		return mapping, err
	}

	s.Logger.Infof("DailyPlayerGamelogs Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("DailyPlayerGamelogs retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, nil
}

func validateDailyPlayerGamelogsURI(options *DailyPlayerGamelogsOptions) error {
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
	if len(options.Week) == 0 {
		return errors.New("missing required option to build the url: Week")
	}
	if len(options.Format) == 0 {
		return errors.New("missing required option to build the url: Format")
	}
	return nil
}
