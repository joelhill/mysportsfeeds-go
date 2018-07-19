package msf

import (
	"context"
	"errors"
	"fmt"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// PlayerInjuriesOptions - Are the options to hit the player injuries endpoint
type PlayerInjuriesOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Format  string

	// Optional URL Params
	Player   string
	Team     string
	Position string
	Sort     string
	Offset   string
	Limit    string
	Force    string
}

// NewPlayerInjuriesOptions - Returns the options with most url parts already set to hit the player injuries endpoint
func (s *Service) NewPlayerInjuriesOptions() *PlayerInjuriesOptions {
	return &PlayerInjuriesOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
	}
}

// PlayerInjuries - hits the https://api.mysportsfeeds.com/{version/pull/{sport}/{season}/date/{date}/dfs.{format} endpoint
func (s *Service) PlayerInjuries(c context.Context, options *PlayerInjuriesOptions) (PlayerInjuriesIO, error) {
	errorPayload := make(map[string]interface{})
	mapping := PlayerInjuriesIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validatePlayerInjuriesURI(options)
	if err != nil {
		return mapping, err
	}

	uri := fmt.Sprintf("%s/%s/pull/%s/current_season.%s?1=1", options.URL, options.Version, options.Sport, options.Format)

	if len(options.Player) > 0 {
		uri = fmt.Sprintf("%s&player=%s", uri, options.Player)
	}

	if len(options.Team) > 0 {
		uri = fmt.Sprintf("%s&team=%s", uri, options.Team)
	}

	if len(options.Position) > 0 {
		uri = fmt.Sprintf("%s&position=%s", uri, options.Position)
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
	s.Logger.Debug("PlayerInjuries API Call")

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
		s.Logger.Errorf("something went wrong making the get request for PlayerInjuries: %s", err.Error())
		return mapping, err
	}

	s.Logger.Infof("PlayerInjuries Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("PlayerInjuries retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, nil
}

func validatePlayerInjuriesURI(options *PlayerInjuriesOptions) error {
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
