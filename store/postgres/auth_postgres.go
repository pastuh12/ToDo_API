package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
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
	sessionID := -1
	refreshExt := 0
	query := fmt.Sprintf("SELECT id, expiresAt FROM %s WHERE userID = $1", tableSessions)
	err := r.store.db.QueryRow(query, session.UserID).Scan(&sessionID, &refreshExt)
	if err != nil {
		if err == sql.ErrNoRows {
			return r.SetSession(ctx, session)
		}
		return errors.Wrap(err, "failed checking session table")
	}

	return r.UpdateSession(ctx, session)

}

func (r *AuthPostgres) SetSession(ctx context.Context, session *models.Session) error {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (userID, refreshToken, expiresAt) VALUES($1, $2, $3) RETURNING id", tableSessions)
	err := r.store.db.QueryRow(query, session.UserID, session.RefreshToken, session.ExpiresAt).Scan(&id)
	if err != nil {
		return errors.Wrap(err, "failed session inserted")
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

func (r *AuthPostgres) GetSessionByToken(ctx context.Context, oldRefreshToken string) (int, error) {
	var refreshExt int64
	var id int
	var userID int
	query := fmt.Sprintf("SELECT id, userID, expiresAt FROM %s WHERE refreshToken = $1", tableSessions)

	err := r.store.db.QueryRow(query, oldRefreshToken).Scan(&id, &userID, &refreshExt)
	if err != nil {
		logrus.Info(err)
		if err == sql.ErrNoRows {
			return 0, errors.New("refreshToken not valid")
		}
		return 0, errors.Wrap(err, "session not update by refreshToken")

	}

	if time.Now().Unix()-refreshExt > 0 {
		deleteQuery := fmt.Sprintf("DELETE FROM %s WHERE id = $1", tableSessions)
		err = r.store.db.QueryRow(deleteQuery, id).Scan()
		if err != nil {
			return 0, errors.Wrap(err, "deletion from session table failed")
		}
		return 0, errors.New("refreshToken expired")
	}

	return userID, nil
}
