package services

import "errors"

type Services struct {
}

var (
	ErrSendingEmail = errors.New("cannot send email")
)
