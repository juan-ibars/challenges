package ad

import (
	. "github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
)

//go:generate mockery --name=FindAdService --filename=findservice.go --output=../../mocks --outpkg=mocks
type FindAdService interface {
	Execute(id UUID) *Ad
}
