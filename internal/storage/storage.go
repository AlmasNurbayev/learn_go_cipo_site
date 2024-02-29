package storage

import (
	"errors"
)

var (
	ErrNotFound = errors.New("not found")
	ErrExist    = errors.New("already exist")
	ErrUnknown  = errors.New("unknown error")
)




