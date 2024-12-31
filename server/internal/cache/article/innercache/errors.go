package innercache

import "errors"

var (
	ErrNotFoundInCache = errors.New("key not found in cache")
)
