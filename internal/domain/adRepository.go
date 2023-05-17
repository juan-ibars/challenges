package domain

import (
	. "github.com/google/uuid"
)

type AdRepository interface {
	Save(ad Ad)
	FindById(id UUID) *Ad
	FindAllAds() []Ad
}
