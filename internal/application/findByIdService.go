package application

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

func (s *FindByIdService) Execute(id UUID) Ad {
	foundAd := s.adRepository.FindById(id)
	if foundAd != nil {
		return *foundAd
	}
	return Ad{}
}
