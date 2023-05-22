package ads

import . "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"

type FindAllService struct {
	adRepository AdRepository
}

func NewFindAllService(adRepository AdRepository) *FindAllService {
	return &FindAllService{adRepository: adRepository}
}

func (s *FindAllService) Execute() []Ad {
	ads := s.adRepository.FindAllAds()
	return ads[:5]
}
