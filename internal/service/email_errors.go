package service

import "errors"

var (
	ErrSendMailTimeout = errors.New("email sending timeout")
)
