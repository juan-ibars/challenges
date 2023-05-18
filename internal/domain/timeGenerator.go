package domain

import . "time"

//go:generate mockery --name=TimeGenerator --filename=timeGenerator.go --output=../mocks --outpkg=mocks
type TimeGenerator interface {
	Generate() Time
}
