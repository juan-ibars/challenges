package application

import (
	"github.com/google/uuid"
)

type AdErrors struct {
	Msg string
	Id  uuid.UUID
}

func (e *AdErrors) Error() string {
	return e.Msg
}
