package application

import . "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"

type SaveAdService struct {
	adRepository AdRepository
}

func NewSaveAdService(adRepository AdRepository) *SaveAdService {
	return &SaveAdService{adRepository: adRepository}
}

func (s *SaveAdService) Execute(title string, description string, price float64) Ad {
	ad := CreateAd(title, description, price)
	s.adRepository.Save(ad)
	return ad
}
