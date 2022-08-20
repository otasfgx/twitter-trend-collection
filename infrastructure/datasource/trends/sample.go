package trends

import (
	"context"
	"rss/usecase"

	"github.com/dghubble/go-twitter/twitter"
)

type (
	twiTest struct {
		client *twitter.Client
	}
)

func NewTwitterSample() usecase.TimelineDataSource {
	return new(twiTest)
}

func (*twiTest) ListTrends(context.Context) ([]usecase.Trend, error) {
	return []usecase.Trend{
		{
			Value: "test",
			Count: 10,
		},
	}, nil
}
