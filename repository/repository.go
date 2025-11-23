package repository

import (
	"errors"

	"github.com/darrkeer/avito-tech-test-task/models"

	"database/sql"
)

type Repository struct {
	db *sql.DB
}

func New(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) createTeam(tx *sql.Tx, title string) (*models.Team, error) {
	var id int
	err := tx.QueryRow(
		"INSERT INTO teams(title) VALUES($1) RETURNING id",
		title,
	).Scan(&id)

	if err != nil {
		return nil, err
	}
	return models.NewTeam(id, title), nil
}

func (r *Repository) updateUser(tx *sql.Tx, name string, isActive bool) (*models.User, error) {
	var id int
	err := tx.QueryRow(
		`INSERT INTO users(name, is_active)
		VALUES($1, $2)
		ON CONFLICT (name) DO UPDATE SET is_active = EXCLUDED.is_active
		RETURNING id`,
		name, isActive,
	).Scan(&id)

	if err != nil {
		return nil, err
	}
	return models.NewUser(id, name, isActive), nil
}

func (r *Repository) addUserToTeam(tx *sql.Tx, teamId int, userId int) error {
	_, err := tx.Exec(
		"INSERT INTO team_members(team_id, user_id) VALUES($1, $2) ON CONFLICT DO NOTHING",
		teamId, userId,
	)

	return err
}

func (r *Repository) AddTeam(r_team *models.ReadTeam) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	team, err := r.createTeam(tx, r_team.TeamName)
	if err != nil {
		return errors.New("TEAM_EXISTS")
	}

	for _, r_member := range r_team.Members {
		member, err := r.updateUser(tx, r_member.Username, r_member.IsActive)
		if err != nil {
			return err
		}

		if err := r.addUserToTeam(tx, team.Id, member.Id); err != nil {
			return err
		}
	}

	return tx.Commit()
}
