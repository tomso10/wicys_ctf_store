package data

import "errors"

var (
	ErrUserExists        = errors.New("user already exists")
	ErrInvalidTeam       = errors.New("invalid team")
	ErrInvalidPassword   = errors.New("invalid password")
	ErrInsufficientFunds = errors.New("insufficient funds")
)
