package postgres

import (
	"github.com/google/uuid"
	. "github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	. "gorm.io/gorm"
)

type Repository struct {
	dbConnection *DB
}

func NewPostgresRepository(db *DB) *Repository {
	return &Repository{dbConnection: db}
}

func (p *Repository) Save(ad Ad) error {
	tx := p.dbConnection.Create(&ad)
	err := tx.Error
	return err
}

func (p *Repository) FindById(id uuid.UUID) (*Ad, error) {
	//TODO implement me
	panic("implement me")
}

func (p *Repository) FindAllAds() ([]Ad, error) {
	//TODO implement me
	panic("implement me")
}
