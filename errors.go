package hasty

import "errors"

// ErrRate will be returned if API requests rate is too high
var ErrRate = errors.New("too many requests")

// ErrNotFound will be returned if item is not found
var ErrNotFound = errors.New("not found")

// ErrAuth will be returned if authentication fails
var ErrAuth = errors.New("not authenticated")

// ErrPerm will be returned if authorization fails
var ErrPerm = errors.New("permission denied")
