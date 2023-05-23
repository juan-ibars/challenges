package ad

import (
	"fmt"
	. "github.com/google/uuid"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/application"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

type FindByIdService struct {
	adRepository AdRepository
}

func NewFindByIdService(adRepository AdRepository) *FindByIdService {
	return &FindByIdService{adRepository: adRepository}
}

func (s *FindByIdService) Execute(id UUID) (Ad, error) {
	foundAd := s.adRepository.FindById(id)
	if foundAd == nil {
		return Ad{}, &application.AdErrors{
			Id:  id,
			Msg: fmt.Sprintf(fmt.Sprintf("ad with id %s not found", id.String())),
		}
	}
	return *foundAd, nil
}
