package domain

import (
	. "github.com/google/uuid"
)

//go:generate mockery --name=AdRepository --filename=adRepository.go --output=../mocks --outpkg=mocks
type AdRepository interface {
	Save(ad Ad) error
	FindById(id UUID) (*Ad, error)
	FindAllAds() ([]Ad, error)
}
