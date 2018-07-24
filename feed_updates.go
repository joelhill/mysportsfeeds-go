package msf

import (
	"context"
	"errors"
	"fmt"
	"time"

	blaster "github.com/joelhill/go-rest-http-blaster"
	logrus "github.com/sirupsen/logrus"
)

// FeedUpdatesOptions - Are the options to hit the feed updates endpoint
type FeedUpdatesOptions struct {
	// URL Parts
	URL     string
	Version string
	Sport   string
	Season  string
	Format  string

	// Optional URL Params
	Force string
}

// NewFeedUpdatesOptions - Returns the options with most url parts already set to hit the feed updates endpoint
func (s *Service) NewFeedUpdatesOptions() *FeedUpdatesOptions {
	return &FeedUpdatesOptions{
		URL:     s.Config.BaseURL,
		Version: s.Config.Version,
		Sport:   s.Config.Sport,
		Season:  s.Config.Season,
		Format:  s.Config.Format,
	}
}

// FeedUpdates - hits the https://api.mysportsfeeds.com/{version}/pull/{sport}/{season}/latest_updates.{format} endpoint
func (s *Service) FeedUpdates(c context.Context, options *FeedUpdatesOptions) (FeedUpdatesIO, int, error) {
	errorPayload := make(map[string]interface{})
	mapping := FeedUpdatesIO{}

	// make sure we have all the required elements to build the full required url string.
	err := validateFeedUpdatesURI(options)
	if err != nil {
		return mapping, 0, err
	}

	t := time.Now()
	cacheBuster := t.Format("20060102150405")

	uri := fmt.Sprintf("%s/%s/pull/%s/latest_updates.%s?cachebuster=%s", options.URL, options.Version, options.Sport, options.Format, cacheBuster)

	if len(options.Force) > 0 {
		uri = fmt.Sprintf("%s&force=%s", uri, options.Force)
	}

	s.Logger = s.Logger.WithFields(logrus.Fields{
		"URI": uri,
	})
	s.Logger.Debug("FeedUpdates API Call")

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
		s.Logger.Errorf("something went wrong making the get request for FeedUpdates: %s", err.Error())
		return mapping, statusCode, err
	}

	s.Logger.Infof("FeedUpdates Status Code: %d", statusCode)

	if client.StatusCodeIsError() {
		s.Logger.Errorf("FeedUpdates retuned an unsuccessful status code. Error: %+v", errorPayload)
	}

	return mapping, statusCode, nil
}

func validateFeedUpdatesURI(options *FeedUpdatesOptions) error {
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
		return errors.New("missing required option to build the url: Sport")
	}
	if len(options.Format) == 0 {
		return errors.New("missing required option to build the url: Format")
	}
	return nil
}
