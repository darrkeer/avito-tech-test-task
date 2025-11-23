package repository

import (
	"github.com/darrkeer/avito-tech-test-task/models"

	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) InsertUser(user models.User) error {
	_, err := r.db.Exec("INSERT INTO users(name) VALUES($1)", user.Name)
	return err
}
