package errs

import (
	"fmt"
)

type TimelineError struct {
	OriginErr error
}

func (t TimelineError) Error() string {
	if t.OriginErr == nil {
		panic("no error")
	}
	return fmt.Errorf("timeline error: %w", t.OriginErr).Error()
}

func (s TimelineError) UnWrap() error {
	return s.OriginErr
}
