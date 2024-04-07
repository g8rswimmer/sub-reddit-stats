package errorx

import "errors"

// ErrDatabase is when there is a database error.  This is
// a generic error and will be wrapped with a specific database
// error.
var ErrDatabase = errors.New("database error")
