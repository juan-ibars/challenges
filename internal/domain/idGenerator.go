package domain

import . "github.com/google/uuid"

type IdGenerator interface {
	Generate() UUID
}
