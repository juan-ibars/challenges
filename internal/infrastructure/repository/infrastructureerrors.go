package repository

type InfrastructureErrors struct {
	Msg string
}

func (e *InfrastructureErrors) Error() string {
	return e.Msg
}
