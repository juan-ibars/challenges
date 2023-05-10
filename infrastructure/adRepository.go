package infrastructure

import (
	"context"
	"github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/domain"
)

type Ads []Ad

var AdRepository Ads

func Save(ad Ad) {
	AdRepository = append(AdRepository, ad)
}

func findById(id uuid.UUID) {
	context.TODO()
}
