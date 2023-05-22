package ads

import . "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"

//go:generate mockery --name=FindAdsService --filename=findallservice.go --output=../../mocks --outpkg=mocks
type FindAdsService interface {
	Execute() []Ad
}
