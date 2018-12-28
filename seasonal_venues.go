package msf

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// SeasonalVenuesOptions - Are the options to hit the seasonal venues endpoint
type SeasonalVenuesOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Format  string

	// Optional URL Params
	Team  string
	Force string
}

// NewSeasonalVenuesOptions - Returns the options with most url parts already set to hit the seasonal venues endpoint
func (s *Service) NewSeasonalVenuesOptions() *SeasonalVenuesOptions {
	return &SeasonalVenuesOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Format:  s.Config.Format,
		Season:  s.Config.Season,
	}
}

// SeasonalVenues - hits the https://api.mysportsfeeds.com/{version}/pull/{season}/{season}/venues.{format} endoint
func (s *Service) SeasonalVenues(options *SeasonalVenuesOptions) (VenuesIO, int, error) {
	errorPayload := make(map[string]interface{})
	mapping := VenuesIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateSeasonalVenuesURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/%s/venues.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Season, options.Format, cacheBuster)

	if len(options.Team) > 0 {
		uri = fmt.Sprintf("%s&team=%s", uri, options.Team)
	}

	if len(options.Force) > 0 {
		uri = fmt.Sprintf("%s&force=%s", uri, options.Force)
	}

	s.Logger = s.Logger.WithFields(logrus.Fields{
		"URI": uri,
	})
	s.Logger.Debug("SeasonalVenues API Call")

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
		s.Logger.Errorf("something went wrong making the get request for SeasonalVenues: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("SeasonalVenues Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("SeasonalVenues retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
}

func validateSeasonalVenuesURI(options *SeasonalVenuesOptions) error {
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
