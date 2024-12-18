package models

import "errors"

// Custom error constants
var (
	ErrMemberNotFound      = errors.New("member not found")
	ErrFailedToUpdate      = errors.New("failed to update member")
	ErrFailedToDelete      = errors.New("failed to delete member")
	ErrInvalidInput        = errors.New("invalid input")
	ErrMemberAlreadyExists = errors.New("member already exists")
	ErrEmptyNameORID       = errors.New("not allowed empty name or id")
)
