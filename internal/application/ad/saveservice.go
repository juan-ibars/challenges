package ad

import . "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"

//go:generate mockery --name=SaveService --filename=saveservice.go --output=../../mocks --outpkg=mocks
type SaveService interface {
	Execute(title string, description string, price float64) (Ad, error)
}
