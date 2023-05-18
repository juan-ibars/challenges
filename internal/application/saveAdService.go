package application

import (
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

type SaveAdService struct {
	idGenerator   IdGenerator
	timeGenerator TimeGenerator
	adRepository  AdRepository
}

func NewSaveAdService(idGenerator IdGenerator, timeGenerator TimeGenerator, adRepository AdRepository) *SaveAdService {
	return &SaveAdService{idGenerator: idGenerator, timeGenerator: timeGenerator, adRepository: adRepository}
}

func (s *SaveAdService) Execute(title string, description string, price float64) Ad {
	ad := Ad{
		Id:          s.idGenerator.Generate(),
		Title:       title,
		Description: description,
		Price:       price,
		Date:        s.timeGenerator.Generate(),
	}
	s.adRepository.Save(ad)
	return ad
}
