package model

import "github.com/pkg/errors"

type Direction int32

const (
	First Direction = iota
	Next
	Prev
)

var ErrInvalidDirection = errors.New("invalid direction")
