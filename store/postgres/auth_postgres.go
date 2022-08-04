package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
	"github.com/todo_api/models"
)

type AuthPostgres struct {
	store *PostgresDB
}

func NewAuthPostgres(s *PostgresDB) *AuthPostgres {
	return &AuthPostgres{
		store: s,
	}
}

func (r *AuthPostgres) CreateUser(ctx context.Context, user *models.User) (*models.User, error) {
	query := fmt.Sprintf(
		"INSERT INTO %s (enconding_password, email_address, name) VALUES($1, $2, $3) RETURNING id",
		tableUsers,
	)
	err := r.store.db.QueryRow(query, user.Password, user.EmailAddress, user.Name).Scan(&user.ID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *AuthPostgres) GetUser(ctx context.Context, user *models.AuthUser) (int, error) {
	var id = 0
	query := fmt.Sprintf("SELECT id FROM %s WHERE enconding_password=$1 AND email_address=$2", tableUsers)
	err := r.store.db.QueryRow(query, user.Password, user.EmailAddress).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (r *AuthPostgres) CheckSession(ctx context.Context, session *models.Session) error {
	check := -1
	query := fmt.Sprintf("SELECT id FROM %s WHERE userID = $1", tableSessions)
	err := r.store.db.QueryRow(query, session.UserID).Scan(&check)
	if err != nil {
		if err == sql.ErrNoRows {
			// logrus.Info("HERE")
			return r.SetSession(ctx, session)
		}
		return errors.Wrap(err, "failed check session table")
	}

	return r.UpdateSession(ctx, session)

}

func (r *AuthPostgres) SetSession(ctx context.Context, session *models.Session) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (userID, refreshToken, expiresAt) VALUES($1, $2, $3) RETURNING id", tableSessions)
	err := r.store.db.QueryRow(query, session.UserID, session.RefreshToken, session.ExpiresAt).Scan(&id)
	if err != nil {
		return errors.Wrap(err, "session not inserted")
	}

	return nil
}

func (r *AuthPostgres) UpdateSession(ctx context.Context, session *models.Session) error {
	var id int
	query := fmt.Sprintf("UPDATE %s SET refreshToken = $1, expiresAt = $2 WHERE userId = $3 RETURNING id", tableSessions)

	err := r.store.db.QueryRow(query, session.RefreshToken, session.ExpiresAt, session.UserID).Scan(&id)
	if err != nil {
		return errors.Wrap(err, "session not update")
	}

	return nil
}
