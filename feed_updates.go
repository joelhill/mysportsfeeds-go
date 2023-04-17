package msf

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	retryablehttp "github.com/hashicorp/go-retryablehttp"
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
func (s *Service) FeedUpdates(options *FeedUpdatesOptions) (FeedUpdatesIO, int, error) {

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
	client := retryablehttp.NewClient()

	req, err := retryablehttp.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		s.Logger.Errorf("client: could not create request: %s", err.Error())
		return mapping, 0, err
	}
	req.Header.Add("Authorization", CompressionHeaderGzip)
	req.Header.Add("Authorization", s.Config.Authorization)

	response, err := client.Do(req)
	if err != nil {
		s.Logger.Errorf("client: error making http request: %s", err.Error())
		return mapping, 0, err
	}

	if response.StatusCode < 200 || response.StatusCode > 300 {
		s.Logger.Errorf("client: something went wrong making the get request for FeedUpdates: %s", err.Error())
		return mapping, response.StatusCode, err
	}

	s.Logger.Infof("FeedUpdates Status Code: %d", response.StatusCode)

	if jErr := json.NewDecoder(response.Body).Decode(&mapping); jErr != nil {
		s.Logger.Errorf("client: error decoding response for FeedUpdates: %s", err.Error())
		return mapping, response.StatusCode, err
	}

	return mapping, response.StatusCode, nil
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
