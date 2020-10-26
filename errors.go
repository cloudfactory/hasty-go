package hasty

import "errors"

// RateError will be returned if API requests rate is too high
var RateError = errors.New("too many requests")
