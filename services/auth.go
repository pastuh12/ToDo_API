package services

import (
	"context"
	"crypto/sha1"
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/todo_api/models"
	"github.com/todo_api/store"
)

const (
	salt string = "438j0984jf29f22d"
)

type AuthService struct {
	ctx   context.Context
	store *store.Store
}

func NewAuthService(ctx context.Context, store *store.Store) *AuthService {
	return &AuthService{
		ctx:   ctx,
		store: store,
	}
}

func (s *AuthService) CreateUser(ctx context.Context, user *models.User) (*Token, error) {
	user.Password = s.EncryptPassword(user.Password)
	user, err := s.store.Authtorization.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := s.CreateSession(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) LoginUser(ctx context.Context, user *models.AuthUser) (*Token, error) {
	user.Password = s.EncryptPassword(user.Password)
	id, err := s.store.Authtorization.GetUser(ctx, user)
	if err != nil {
		return nil, err
	}

	token, err := s.CreateSession(ctx, id)
	if err != nil {
		return nil, err
	}

	logrus.Info(token)

	return token, nil
}

func (s *AuthService) CreateSession(ctx context.Context, id int) (*Token, error) {
	var session models.Session

	token, err := NewToken(id)
	if err != nil {
		return nil, errors.Wrap(err, "token not created")
	}

	logrus.Info(token)

	session.UserID = id
	session.RefreshToken = token.RefreshToken
	session.ExpiresAt = token.ExpiresAt

	err = s.store.Authtorization.CheckSession(ctx, &session)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) UpdateToken(ctx context.Context, id int) (*Token, error) {
	var session models.Session

	token, err := NewToken(id)
	if err != nil {
		return nil, errors.Wrap(err, "token not created")
	}

	session.UserID = id
	session.RefreshToken = token.RefreshToken
	session.ExpiresAt = token.ExpiresAt

	err = s.store.Authtorization.UpdateSession(ctx, &session)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (s *AuthService) EncryptPassword(str string) string {
	hash := sha1.New()
	hash.Write([]byte(str))
	hashBytes := hash.Sum([]byte(salt))

	return fmt.Sprintf("%x", hashBytes)
}
