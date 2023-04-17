package msf

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// PlayersOptions - Are the options to hit the players endpoint
type PlayersOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Format  string

	// Optional URL Params
	Date         string
	Team         string
	RosterStatus string
	Player       string
	Position     string
	Country      string
	Sort         string
	Offset       string
	Limit        string
	Force        string
}

// NewPlayersOptions - Returns the options with most url parts already set to hit the players endpoint
func (s *Service) NewPlayersOptions() *PlayersOptions {
	return &PlayersOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
	}
}

// Players - hits the https://api.mysportsfeeds.com/v2.0/pull/mlb/players.{format} endoint
func (s *Service) Players(options *PlayersOptions) (PlayersIO, int, error) {
	errorPayload := make(map[string]interface{})
	mapping := PlayersIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validatePlayersURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/players.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Format, cacheBuster)

	if len(options.Date) > 0 {
		uri = fmt.Sprintf("%s&date=%s", uri, options.Date)
	}

	if len(options.Team) > 0 {
		uri = fmt.Sprintf("%s&team=%s", uri, options.Team)
	}

	if len(options.RosterStatus) > 0 {
		uri = fmt.Sprintf("%s&rosterStatus=%s", uri, options.RosterStatus)
	}

	if len(options.Player) > 0 {
		uri = fmt.Sprintf("%s&player=%s", uri, options.Player)
	}

	if len(options.Position) > 0 {
		uri = fmt.Sprintf("%s&position=%s", uri, options.Position)
	}

	if len(options.Country) > 0 {
		uri = fmt.Sprintf("%s&country=%s", uri, options.Country)
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
	s.Logger.Debug("Players API Call")

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
		s.Logger.Errorf("something went wrong making the get request for Players: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("Players Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("Players retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
}

func validatePlayersURI(options *PlayersOptions) error {
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
