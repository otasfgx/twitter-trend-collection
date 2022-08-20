package usecase

import "context"

type (
	// Model
	Trend struct {
		Value string
		Count int
	}
	// IF
	TimelineDataSource interface {
		ListTrends(context.Context) ([]Trend, error)
	}
)

// usecase
type (
	Usecase interface {
		GetTrends(context.Context) ([]Trend, error)
	}
	usecase struct {
		timelineDsrc TimelineDataSource
	}
)

func NewUseCase(tlDsrc TimelineDataSource) Usecase {
	return &usecase{
		timelineDsrc: tlDsrc,
	}
}

func (u *usecase) GetTrends(ctx context.Context) ([]Trend, error) {
	return u.timelineDsrc.ListTrends(ctx)
}
