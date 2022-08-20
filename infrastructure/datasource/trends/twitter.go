package trends

import (
	"context"
	"rss/domain/errs"
	"rss/usecase"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type (
	twi struct {
		client *twitter.Client
	}
)

const (
	JAPAN_WOEID = 23424856
)

func NewTwitter(apiKey, apiSecret, accessToken, accessSecret string) usecase.TimelineDataSource {
	instance := new(twi)
	config := oauth1.NewConfig(apiKey, apiSecret)
	token := oauth1.NewToken(accessToken, accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	instance.client = twitter.NewClient(httpClient)
	return instance
}

func (t *twi) ListTrends(context.Context) ([]usecase.Trend, error) {
	trendsList, _, err := t.client.Trends.Place(JAPAN_WOEID, nil)
	if err != nil {
		return nil, errs.TimelineError{OriginErr: err}
	}

	result := []usecase.Trend{}
	for _, trends := range trendsList {
		for _, trend := range trends.Trends {
			result = append(result, usecase.Trend{
				Value: trend.Name,
				Count: int(trend.TweetVolume),
			})
		}
	}

	return result, nil
}
