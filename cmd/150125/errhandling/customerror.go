package errhandling

import (
	"errors"
	"fmt"

	"weezel/meetup/internal/logger"
)

type CustomSQLError struct {
	Err    error  // Original error
	Query  string // SQL query that failed
	UserID int64  // User ID who ran the query
}

// Implements error.Error() interface
func (c CustomSQLError) Error() string {
	return fmt.Sprintf("query failed: %s", c.Err)
}

func funcThatReturnsCustomErr() error {
	return CustomSQLError{
		Err:    errors.New("the real error is here"),
		Query:  "SELECT * FROM aaaaaaa;",
		UserID: 100,
	}
}

func customErrorDemo() {
	err := funcThatReturnsCustomErr()
	if err != nil {
		var sqlErr CustomSQLError
		if errors.As(err, &sqlErr) {
			logger.Logger.Error().Err(err).
				Int64("user_id", sqlErr.UserID).
				Str("query", sqlErr.Query).
				Msg("SQL query failed")
		} else {
			logger.Logger.Error().Err(err).Msg("Failed to return custom error")
		}
	}
}
