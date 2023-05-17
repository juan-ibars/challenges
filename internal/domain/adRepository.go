package domain

import (
	. "github.com/google/uuid"
)

type adRepository interface {
	Save(ad Ad)
	FindById(id UUID) *Ad
	FindAllAds() []Ad
}
