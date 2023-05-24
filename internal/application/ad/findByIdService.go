package ad

import (
	. "github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

type FindByIdService struct {
	adRepository AdRepository
}

func NewFindByIdService(adRepository AdRepository) *FindByIdService {
	return &FindByIdService{adRepository: adRepository}
}

func (s *FindByIdService) Execute(id UUID) (Ad, error) {
	foundAd, err := s.adRepository.FindById(id)
	if err != nil {
		return Ad{}, err
	}
	return *foundAd, err
}
