package application

//go:generate mockery --name=Service --filename=service.go --output=../mocks --outpkg=mocks
type Service interface {
	Execute(value ...interface{}) interface{}
}
