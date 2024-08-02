package store

import "errors"

var (
	ErrRecordExists = errors.New("ERROR: record exists")
	ErrNotFound     = errors.New("ERROR: record not found")
	ErrBadRequest   = errors.New("ERROR: bad request")
)
