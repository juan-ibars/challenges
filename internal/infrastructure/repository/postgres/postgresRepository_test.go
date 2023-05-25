package postgres

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.mpi-internal.com/juan-ibars/learning-go/internal/domain"
	"gorm.io/driver/postgres"
	. "gorm.io/gorm"
	"testing"
	"time"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "challenges"
)

var dbCon *DB
var dbErr error

func TestMain(m *testing.M) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	dbCon, dbErr = Open(postgres.Open(dsn), &Config{})

	if dbErr != nil {
		panic("failed to connect database")
	}

	table := &domain.Ad{}
	intDbErr := createTable(dbCon, table)
	defer func() { _ = removeTable(dbCon, table) }()

	if intDbErr != nil {
		panic("failed to create table")
	}

	m.Run()
}

func TestSaveInPostgres(t *testing.T) {
	repository := NewPostgresRepository(dbCon)

	id, _ := uuid.NewRandom()
	date, _ := time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	ad := domain.Ad{
		Id:          id,
		Title:       "Some title",
		Description: "Some description",
		Price:       10.5,
		Date:        date,
	}
	err := repository.Save(ad)

	var actualAd domain.Ad
	dbCon.Raw("SELECT * FROM ads WHERE id = ?", id).Scan(&actualAd)
	assert.Nil(t, err)
	assert.Equal(t, ad, actualAd)
}

func createTable(db *DB, table interface{}) error {
	return db.Migrator().CreateTable(table)
}

func removeTable(db *DB, table interface{}) error {
	return db.Migrator().DropTable(table)
}
